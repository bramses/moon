make a command that:
1. reads the current file
2. has a dropdown showing 3 user selected commands: "test1, "test2", "test3"
3. if "test3" selected, ask for user typed input
4. println selection and optional prompt

---

To implement a command with the specified requirements, add the following code to your existing CLI app:

1. Add a new command called `readfile`:

```go
readfileCmd := &cobra.Command{
	Use:   "readfile [file]",
	Short: "Read the current file and execute user-selected commands",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		readFileAndExecuteCommand(filename)
	},
}
```

2. Add the `readFileAndExecuteCommand` function:

```go
func readFileAndExecuteCommand(filename string) {
	// Read the file content
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the file content
	fmt.Printf("File content:\n%s\n\n", string(data))

	// Show dropdown with user-selected commands
	prompt := promptui.Select{
		Label: "Select a command",
		Items: []string{"test1", "test2", "test3"},
	}

	index, command, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// If "test3" is selected, ask for user input
	var userInput string
	if index == 2 {
		inputPrompt := promptui.Prompt{
			Label: "Type your input",
		}
		userInput, err = inputPrompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Print the selected command and optional user input
	fmt.Println("Selected command:", command)
	if userInput != "" {
		fmt.Println("User input:", userInput)
	}
}
```

3. Register the new `readfile` command:

```go
app.AddCommand(initCmd, commitCmd, configCmd, commitAllCmd, readfileCmd)
```

Now you should have a new command called `readfile` that reads the content of a file, displays a dropdown with the commands "test1", "test2", and "test3", asks for user input if "test3" is selected, and prints the selection and optional input.

To test the new command, rebuild and run the CLI app:

```bash
go build
./llm-cli readfile sample.txt
```

Replace `sample.txt` with the name of the file you want to read.