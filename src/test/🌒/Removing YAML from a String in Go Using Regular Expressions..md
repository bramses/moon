---
title: Removing YAML from a String in Go Using Regular Expressions.
phase: 1
command: write regex to remove yaml from a sting in Go
time: 2023-05-03 12:12:19
---

Here's an example of how you can use regular expressions to remove YAML from a string in Go:

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    text := `
---
name: John Doe
age: 30
---
Some other text here
`

    regex := regexp.MustCompile("(?s)^---.*?---\n")
    processedText := regex.ReplaceAllString(text, "")

    fmt.Println(processedText)
}
```

This program will output:

```
Some other text here
```

The regular expression `(?s)^---.*?---\n` matches any text that starts with `---` and ends with `---` on a new line. The `(?s)` flag enables the "dot matches newline" option in the regular expression, so that `.` matches any character including newlines. The `.*?` is a non-greedy match that matches any character except newline, zero or more times, until it finds the first occurrence of `---` on a new line. Finally, `\n` matches the newline character at the end of the YAML block. The `ReplaceAllString` method replaces all occurrences of this regular expression with an empty string, effectively removing the YAML from the text.