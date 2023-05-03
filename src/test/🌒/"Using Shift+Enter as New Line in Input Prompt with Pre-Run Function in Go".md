---
title: "Using Shift+Enter as New Line in Input Prompt with Pre-Run Function in Go"
phase: 1
command: so if i want multiline i have to add a flag? what if i want shift+enter to mean new line? ---
title: "Cobra Flag Package for Multiline Input"
phase: 1
command: but it only happens when i type a bunch of chars or hit return on accident etc, how can i do multiline input w cobra ---
title: 
phase: 1
command: why is this code causing multipe user inputs to show up sometimes func promptUserInput() string {
	prompt := promptui.Prompt{
		Label: "Enter input",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
time: 2023-05-03 12:08:25
---

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
time: 2023-05-03 12:15:39
---

If you want to use Shift+Enter to mean a new line in the input prompt, you can add a pre-run function to the command that replaces all occurrences of Shift+Enter with a newline character before the flag is parsed. Here is an example of how you can achieve this:

```
cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
    // Replace all Shift+Enter occurrences with a newline character
    userInput = strings.Replace(userInput, "\r\n", "\n", -1)
}

var userInput string

cmd.Flags().StringVarP(&userInput, "input", "i", "", "Multiline input")

// Print the input prompt message
fmt.Printf("Enter input (use Shift+Enter for new line):\n")

// Read user input from standard input
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    // Replace all Shift+Enter occurrences with a newline character
    userInput = strings.Replace(userInput, "\r\n", "\n", -1)

    // Stop reading input when an empty line is entered
    if scanner.Text() == "" {
        break
    }

    // Append the user input to the existing input
    userInput += scanner.Text() + "\n"
}
```

In the pre-run function, we replace all occurrences of Shift+Enter (`\r\n`) with a newline character (`\n`). In the input prompt loop, we do the same thing before appending the user input to the existing input. This allows the user to use Shift+Enter for new lines in the input prompt.