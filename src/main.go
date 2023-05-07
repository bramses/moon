package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	earlyReturnFlag = false
	openAPIKey      string
	envVar          string
)

func main() {
	catchCTRLC()
	var rootCmd = &cobra.Command{
		Use:   "moon",
		Short: "A CLI from the moon, beamed directly to your terminal",
	}

	rootCmd.PersistentFlags().StringVar(&openAPIKey, "openAPIKey", "", "OPENAI API KEY [Required]")
	if env := os.Getenv("OPENAI_API_KEY"); env != "" {
		envVar = env
		rootCmd.PersistentFlags().Set("openAPIKey", envVar)
	}

	rootCmd.AddCommand(newCmd, orbitCmd, readMeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// readMeCmd
var readMeCmd = &cobra.Command{
	Use:   "readMe",
	Short: "Read all files in the directory",
	Run:   readMe,
}

func readMe(cmd *cobra.Command, args []string) {

	parentFolder, _ := cmd.Flags().GetString("parentFolder")
	if parentFolder == "" { // set to os.Getwd() if empty
		parentFolder, _ = os.Getwd()
	}

	readmeStr := ""

	err := filepath.Walk(parentFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			fmt.Printf("running on %s \n", content)
			// fmt.Printf("%s: %s \n", path, content)
			res := ssereq("Summarize the file in README format: {h1 - title of file}\\n {summary of file content} for this content:\n" + string(content))
			readmeStr += fmt.Sprintf("# %s\n%s\n", info.Name(), res)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		return
	}

	// write to file parentFolder/README.md
	err = ioutil.WriteFile(parentFolder+"/README.md", []byte(readmeStr), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}

// newCmd
// Create a new command called "new" that runs the newProject function
var newCmd = &cobra.Command{
	Use:   "new [project_name]",
	Short: "Create a new project",
	Long:  `Create a new project with the specified structure and configuration file.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   newProject,
}

func init() {
}

func newProject(cmd *cobra.Command, args []string) {
	// get current working directory
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}

	folderNames := []string{"moon"}

	for _, folderName := range folderNames {
		fullPath := filepath.Join(cwd, folderName)
		if err := os.Mkdir(fullPath, 0755); err != nil {
			fmt.Printf("Error creating directory: %s, %v\n", fullPath, err)
			return
		}
	}

	configContent := `{
        commands: [
            {
                command: "{user_prompt} {file_picker}",
                name: "custom + file picker",
                description: "custom command with file picker"
            }
        ]
    }`

	configPath := filepath.Join(cwd, "moon.config.json")
	if err := ioutil.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		fmt.Printf("Error creating moon.config.js: %v\n", err)
		return
	}

	// create a history file
	historyPath := filepath.Join(cwd, "moon.history.json")
	if err := ioutil.WriteFile(historyPath, []byte("{}"), 0644); err != nil {
		fmt.Printf("Error creating moon.history.json: %v\n", err)
		return
	}

	fmt.Println("New project structure created successfully.")
}

func catchCTRLC() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		earlyReturnFlag = true
		fmt.Println("\nCTRL-C detected. Returning early...")
	}()
}

// orbitCmd
var orbitCmd = &cobra.Command{
	Use:   "orbit",
	Short: "Choose a command from moon.config.json to run against",
	Run:   orbit,
}

func orbit(cmd *cobra.Command, args []string) {

	config, err := ReadConfig("moon.config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the filtered commands and execute the selected one
	selectedCommand := displayCommands(config.Commands)
	if selectedCommand == nil {
		fmt.Println("No command selected")
		return
	}

	executeCommand(selectedCommand, parentFolder, args[0])
}

func init() {
	orbitCmd.Flags().Int("number", 0, "Orbit number (1, 2, 3, 4)")
	orbitCmd.Flags().String("parentFolder", "", "Parent folder")
}

func filterOrbitCommands(commands []Command, orbit int) []Command {
	var filtered []Command
	for _, cmd := range commands {
		if cmd.Orbit == orbit {
			filtered = append(filtered, cmd)
		}
	}
	return filtered
}

func displayCommands(commands []Command) *Command {
	prompt := promptui.Select{
		Label: "Select a command",
		Items: commands,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U0001F315 {{ .Name | cyan }} ({{ .Description | red }})",
			Inactive: "  {{ .Name | cyan }} ({{ .Description | red }})",
			Selected: "\U0001F315 {{ .Name | red | cyan }}",
		},
	}

	index, _, err := prompt.Run()

	if err != nil {
		if err == promptui.ErrInterrupt {
			os.Exit(-1)
		}
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return &commands[index]
}

func executeCommand(command *Command, parentFolder string, phase string) {

	// Get user inputs and put them into the command template
	interpolatedCommand, interpolatedTitle := interpolateCommand(command.Command, parentFolder)

	if earlyReturnFlag {
		os.Exit(1)
	}

	// println(interpolatedCommand)
	// Call the LLM API (or any other external function)
	// This is a placeholder function and should be replaced with the actual API call
	// response := callLLM(interpolatedCommand)
	var fullres = ssereq(interpolatedCommand)

	// Generate the inferred title
	inferredTitle := time.Now().Format("20060102150405") + ".md"

	// Get the title from the LLM API
	summary := ssereq("Summarize the following into a file name: " + fullres)

	yaml := "---\n" +
		"title: " + summary + "\n" +
		"phase: " + phase + "\n" +
		"command: " + "\"" + interpolatedTitle + "\"" + "\n" +
		"time: " + time.Now().Format("2006-01-02 15:04:05") + "\n" +
		"---\n\n"

	if summary != "" {
		inferredTitle = summary + ".md"
	}

	// Save the content to a file with the inferred title
	saveToFile(inferredTitle, yaml+fullres, parentFolder, phase)

	// Save command to history
	saveToHistory(inferredTitle, parentFolder, interpolatedTitle, phase)

	// Open the file in the default editor
	// openFile(inferredTitle, parentFolder, phase)
}

type History struct {
	Title   string `json:"title"`
	Command string `json:"command"`
	Phase   string `json:"phase"`
	Time    string `json:"time"`
}

func saveToHistory(title string, parentFolder string, command string, phase string) {
	// Save the content to a file with the inferred title

	// save args to json object
	history := History{
		Title:   title,
		Command: command,
		Phase:   phase,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}

	filePath := filepath.Join(parentFolder, "history.json")

	// Read current content in history.json if it exists
	var historyList []History
	if _, err := os.Stat(filePath); err == nil {
		// Read the file
		historyJson, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		// Convert json to history list
		err = json.Unmarshal(historyJson, &historyList)
		if err != nil {
			fmt.Printf("Error converting json to history list: %v\n", err)
			return
		}
	}

	// Append new history to the list
	historyList = append(historyList, history)

	// Convert history list to json
	// historyJson, err := json.Marshal(historyList)
	// if err != nil {
	// 	fmt.Printf("Error converting history to json: %v\n", err)
	// 	return
	// }

	// pretty print json
	prettyHistoryJson, err := json.MarshalIndent(historyList, "", "    ")
	if err != nil {
		fmt.Printf("Error pretty printing history json: %v\n", err)
	}

	// Save the json to history.json
	err = ioutil.WriteFile(filePath, []byte(prettyHistoryJson), 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

}

func openFile(title string, parentFolder string, phase string) {
	// Open the file in the default editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	// Open the file in the default editor
	cmd := exec.Command(editor, parentFolder+"/"+title)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func interpolateCommand(command string, parentFolder string) (string, string) {

	titleString := command

	commandHandlers := []struct {
		regex   *regexp.Regexp
		handler func(string) (string, error)
	}{
		{regexp.MustCompile(`\{user_prompt\}`), handleUserPrompt},
		{regexp.MustCompile(`\{file_picker\}`), func(match string) (string, error) {
			return handlePhasePicker()
		}},
		{regexp.MustCompile(`\{clipboard\}`), handleClipboard},
	}

	for _, ch := range commandHandlers {
		for {
			match := ch.regex.FindStringIndex(command)

			if match == nil {
				break
			}

			replacement, err := ch.handler(command[match[0]:match[1]])

			if err != nil {
				fmt.Printf("Error handling command: %v\n", err)
				return "", ""
			}

			if command[match[0]:match[1]] == "{file_picker}" {

				escapeQuotes := strings.ReplaceAll(replacement, "\"", "\\\"")

				titleString = titleString[:match[0]] + escapeQuotes + titleString[match[1]:]

				fileContent, err := readContentFromFile(replacement)

				if err != nil {
					fmt.Printf("Error reading file %v\n", err)
				}

				command = command[:match[0]] + fileContent + command[match[1]:]
			} else {
				escapeQuotes := strings.ReplaceAll(replacement, "\"", "\\\"")
				titleString = titleString[:match[0]] + escapeQuotes + titleString[match[1]:]
				command = command[:match[0]] + replacement + command[match[1]:]
			}

		}
	}

	return command, titleString
}

func handleUserPrompt(_ string) (string, error) {
	userInput := promptUserInput()
	return userInput, nil
}

func readContentFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// remove yml from content using regex
	content = regexp.MustCompile(`(?m)^---\n(.|\n)*---\n`).ReplaceAll(content, []byte(""))

	return string(content), nil
}

func handlePhasePicker() (string, error) {

	_, selectedFilePath := promptPhaseFilePicker()

	return selectedFilePath, nil
}

func handleClipboard(_ string) (string, error) {

	clipboardContent, err := clipboard.ReadAll()
	if err != nil {
		return "", fmt.Errorf("error reading from clipboard: %w", err)
	}
	return clipboardContent, nil
}

func moonFolder() string {
	// get working_dir/moon
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return ""
	}

	moonFolder := filepath.Join(workingDir, "moon")

	// create moon folder if it doesn't exist
	if _, err := os.Stat(moonFolder); os.IsNotExist(err) {
		err = os.Mkdir(moonFolder, 0755)
		if err != nil {
			fmt.Printf("Error creating moon folder: %v\n", err)
			return ""
		}
	}

	return moonFolder
}

func promptPhaseFilePicker() (string, string) {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return "", ""
	}

	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		fmt.Printf("Error reading phase folder: %v\n", err)
		return "", ""
	}

	fileList := []string{}
	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	prompt := promptui.Select{
		Label: "Select a file",
		Items: fileList,
	}

	_, selectedFile, err := prompt.Run()
	if err != nil {
		if err == promptui.ErrInterrupt {
			os.Exit(-1)
		}
		fmt.Printf("Prompt failed %v\n", err)
		return "", ""
	}

	return selectedFile, filepath.Join(workingDir, selectedFile)
}

func callLLM(input string) string {
	// Replace this function with the actual LLM API call
	return input
}

func promptUserInput() string {
	prompt := promptui.Prompt{
		Label: "Enter input",
	}

	result, err := prompt.Run()
	if err != nil {
		if err == promptui.ErrInterrupt {
			os.Exit(-1)
		}
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func saveToFile(title string, content string, parentFolder string, phaseNumber string) {
	folderName := moonFolder()
	if folderName == "" {
		fmt.Printf("Invalid phase number: %s\n", phaseNumber)
		return
	}

	filePath := filepath.Join(parentFolder, folderName, title)

	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Printf("File saved as %s\n", filePath)
}
