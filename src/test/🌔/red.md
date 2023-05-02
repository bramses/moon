In order to change `interpolateCommand` to get the content from the `selectedFile`, you need to update the `promptPhaseFilePicker` function to return the file path along with the file name, and then read the content of the file in the `interpolateCommand` function. Here's the updated code:

```go
func interpolateCommand(command string, parentFolder string) (string, string) {
	userPromptRegex := regexp.MustCompile(`\{user_prompt\}`)
	phasePickerRegex := regexp.MustCompile(`\{phase_(\d+)__file_picker\}`)

	for {
		userPromptMatch := userPromptRegex.FindStringIndex(command)
		phasePickerMatch := phasePickerRegex.FindStringSubmatchIndex(command)

		if userPromptMatch == nil && phasePickerMatch == nil {
			break
		}

		if userPromptMatch != nil && (phasePickerMatch == nil || userPromptMatch[0] < phasePickerMatch[0]) {
			userInput := promptUserInput()
			command = command[:userPromptMatch[0]] + userInput + command[userPromptMatch[1]:]
		} else if phasePickerMatch != nil {
			phaseNumber := command[phasePickerMatch[2]:phasePickerMatch[3]]
			selectedFile, selectedFilePath := promptPhaseFilePicker(phaseNumber, parentFolder)
			command = command[:phasePickerMatch[0]] + selectedFile + command[phasePickerMatch[1]:]

			content, err := ioutil.ReadFile(selectedFilePath)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return "", ""
			}

			return command, string(content)
		}
	}

	return command, ""
}

// Update the promptPhaseFilePicker function to return the file path
func promptPhaseFilePicker(phaseNumber string, parentFolder string) (string, string) {
	// Replace this list with actual file names from the specified phase
	fileList := []string{"file1.md", "file2.md", "file3.md"}

	prompt := promptui.Select{
		Label: "Select a file from phase " + phaseNumber,
		Items: fileList,
	}

	_, selectedFile, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", ""
	}

	folderName := phaseFolderName(phaseNumber)
	selectedFilePath := filepath.Join(parentFolder, folderName, selectedFile)

	return selectedFile, selectedFilePath
}
```

Now, the `interpolateCommand` function returns the command with the selected file name and the content of the selected file. The `promptPhaseFilePicker` function has also been updated to return the file path along with the file name, so you can read the content of the file in the `interpolateCommand` function.