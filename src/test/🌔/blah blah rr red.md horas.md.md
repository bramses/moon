---
command: "blah blah rr red.md horas.md"
time: "2021-10-20T15:00:00-07:00"
---

blah blah rr In order to change `interpolateCommand` to get the content from the `selectedFile`, you need to update the `promptPhaseFilePicker` function to return the file path along with the file name, and then read the content of the file in the `interpolateCommand` function. Here's the updated code:

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

Now, the `interpolateCommand` function returns the command with the selected file name and the content of the selected file. The `promptPhaseFilePicker` function has also been updated to return the file path along with the file name, so you can read the content of the file in the `interpolateCommand` function. Here's the merged version of the `promptPhaseFilePicker` function:

```go
func promptPhaseFilePicker(phaseNumber, parentFolder string) (string, string) {
	folderName := phaseFolderName(phaseNumber)
	if folderName == "" {
		return "", ""
	}

	phaseFolderPath := filepath.Join(parentFolder, folderName)
	files, err := ioutil.ReadDir(phaseFolderPath)
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
		Label: "Select a file from phase " + phaseNumber,
		Items: fileList,
	}

	_, selectedFile, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", ""
	}

	selectedFilePath := filepath.Join(phaseFolderPath, selectedFile)
	return selectedFile, selectedFilePath
}
```

Now, the `promptPhaseFilePicker` function returns both the selected file name and the file path. This change should be compatible with the updated `interpolateCommand` function provided in the previous response.