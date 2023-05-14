# Function to Send Prompt to OpenAI GPT-3 API and Extract Delta Content via Server-Sent Events (SSE).md
# Function to Send Prompt to OpenAI GPT-3 API and Extract Delta Content via Server-Sent Events (SSE)

This code defines a function `ssereq` that sends a prompt to the OpenAI GPT-3 API via an HTTP POST request and receives a response using Server-Sent Events (SSE). The prompt is provided as an argument to the function and is sent as a message in the request body. 

The response received via SSE is a sequence of JSON messages that are parsed to extract the delta.content field, which is then written to a file and returned by the function as a string. 

The function also sets some headers in the HTTP request and handles errors that may occur during the process.
# Parsing JSON Data to Custom Structs.md
# Parsing JSON Data to Custom Structs

The code in this file defines the `Command` and `Config` structs and their respective fields using the `json` package to define the field names in JSON format. 

Moreover, the file includes a function named `ReadConfig()` that accepts a file path, extracts the JSON data from the file and maps it to the defined `Config` struct. It returns an error in case of any reading or parsing issues.
# README.md
# Helpful Assistant README

This file contains a brief instruction on how to summarize a document in README format. The content emphasizes the importance of creating a clear and concise summary with an accurate title that reflects the main topics covered in the document. The goal is to provide a quick and easy-to-understand overview of the content for readers who may not have the time or interest to read the whole document. A good README should be easy to scan, highlight the key points, and provide a clear structure that enables readers to quickly find the information they need. By following these guidelines, you can create an effective README that helps readers quickly understand the purpose and content of a document.
# file name: How to gracefully close an SSE connection with response.close() or body.close() in Python.md

