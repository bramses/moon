---
title: Function to Send Prompt to OpenAI GPT-3 API and Extract Delta Content via Server-Sent Events (SSE)
command: "what does this code do /Users/bram/Dropbox/PARA/Projects/moon/src/stream.go"
time: 2023-05-07 20:07:39
---

This code defines a function `ssereq` that sends a prompt to the OpenAI GPT-3 API via an HTTP POST request and receives a response using Server-Sent Events (SSE). The prompt is provided as an argument to the function and is sent as a message in the request body. The response received via SSE is a sequence of JSON messages that are parsed to extract the delta.content field which is then written to a file and returned by the function as a string. The function also sets some headers in the HTTP request and handles errors that may occur during the process.