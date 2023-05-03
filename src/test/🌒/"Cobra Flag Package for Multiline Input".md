---
title: "Cobra Flag Package for Multiline Input"
phase: 1
command: but it only happens when i type a bunch of chars or hit return on accident etc, how can i do multiline input w cobra

The code is causing multiple user inputs to show up because the function `promptUserInput()` is being called multiple times. Each time the function is called, it creates a new prompt for the user to enter their input. 

To avoid this issue, you can try to store the user input once it is prompted and use it for subsequent processing. Another solution would be to ensure that the function is only called when necessary or limit the number of times it can be called.
time: 2023-05-03 12:10:37
---

In terms of multiline input with Cobra, you can use the `Cobra` flag package to define a flag for the multiline input. Here is an example of how you can define a flag for multiline input:

```
var userInput string

cmd.Flags().StringVarP(&userInput, "input", "i", "", "Multiline input")
```

With this flag, you can use the `Prompt` package to prompt the user for multiline input and store the input in the `userInput` variable :

```
fmt.Printf("Enter input:\n")
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    if scanner.Text() == "" {
        break
    }
    userInput += scanner.Text() + "\n"
}
```

I hope this helps!