---
title: file name: How to gracefully close an SSE connection with response.close() or body.close() in Python
command: "getting err in code no body.close /Users/bram/Dropbox/PARA/Projects/moon/src/moon-test/Function to Send Prompt to OpenAI GPT-3 API and Extract Delta Content via Server-Sent Events (SSE).md"
time: 2023-05-07 20:19:24
---

It seems like you are missing the `response.close()` or `body.close()` method, which is used to close the SSE connection to the API after receiving the response. This may cause the connection to remain open and consume resources unnecessarily.

You can add `response.close()` or `body.close()` at the end of the `for` loop where the SSE response messages are being processed. This will ensure that the connection is closed gracefully after processing the response.

Here is an example:

```python
for event in response.raw.stream():
    # process SSE events
    ...

response.close()  # or body.close()
```

Make sure to call the correct method depending on how you are accessing the response. If you are using the `requests` library, you can access the response body using `response.raw` and use `response.close()`. If you are using the built-in `http.client.HTTPSConnection` object, you can access the response body using `body` and use `body.close()`.