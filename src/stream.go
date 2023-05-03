package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream"`
	Temperature float64   `json:"temperature,omitempty"`
	// Add other fields as needed
}

type SSEData struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

func ssereq(prompt string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Define the request body
	reqBody := RequestBody{
		Model: "gpt-3.5-turbo", // Replace with the desired model ID
		Messages: []Message{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: prompt},
		},
		Stream:      true,
		Temperature: 0.0,
	}

	// Convert the request body to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY")) // Replace with your actual API key

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Create a file to write the received tokens
	file, err := os.Create("tokens.txt")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Get the current working directory and construct the file path
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	filePath := filepath.Join(wd, "tokens.txt")

	// Print the file path as a clickable link
	fmt.Printf("Tokens will be written to the following file: file://%s\n", filePath)

	// full text
	var fulltext string

	// Listen for Server-Sent Events (SSE)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()

		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				// The stream has ended
				fmt.Println("The stream has ended.")
				break
			}
			log.Fatalf("Error reading SSE: %v", err)
		}

		// Check for the "data:" prefix to extract the event data
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))

			// Check for the "[DONE]" message indicating the end of the stream
			if data == "[DONE]" {
				break
			}

			// Parse the JSON data to extract the delta.content field
			var sseData SSEData
			err := json.Unmarshal([]byte(data), &sseData)
			if err != nil {
				log.Fatalf("Error parsing JSON data: %v", err)
			}

			// Write the delta.content to the file
			_, err = file.WriteString(sseData.Choices[0].Delta.Content)
			if err != nil {
				log.Fatalf("Error writing to file: %v", err)
			}

			fulltext += sseData.Choices[0].Delta.Content
			// print token to console
			fmt.Print(sseData.Choices[0].Delta.Content)

			// Flush the file buffer to disk to ensure immediate update
			err = file.Sync()
			if err != nil {
				log.Fatalf("Error syncing file: %v", err)
			}
		}
	}

	fmt.Println("Tokens written to file successfully.")

	return fulltext
}
