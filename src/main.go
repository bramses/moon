package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "moon",
		Short: "Moon is a CLI tool for using LLMs to phase ideas to programs",
	}

	rootCmd.AddCommand(newCmd, chatCmd, phaseCmd, orbitCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
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
	newCmd.AddCommand(newChatCmd)
}

func newProject(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide a project name.")
		return
	}

	projectName := args[0]
	if err := os.Mkdir(projectName, 0755); err != nil {
		fmt.Printf("Error creating project directory: %s, %v\n", projectName, err)
		return
	}

	folderNames := []string{"🌑", "🌒", "🌓", "🌔", "🌕"}

	for _, folderName := range folderNames {
		fullPath := filepath.Join(projectName, folderName)
		if err := os.Mkdir(fullPath, 0755); err != nil {
			fmt.Printf("Error creating directory: %s, %v\n", fullPath, err)
			return
		}
	}

	configContent := `{
        commands: [
            {
                from: 1,
                to: 3,
                command: "blah blah {user_prompt} {phase3.md} {phase1.md}",
                prompt: true,
                name: "blah",
                description: "this is a long description for blah"
            },
            {
                from: 1,
                to: 1,
                command: "write a fix for {user_prompt}",
                name: "name",
                "description": "fixes x,y,z"
            }
        ]
    }`

	configPath := filepath.Join(projectName, "moon.config.json")
	if err := ioutil.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		fmt.Printf("Error creating moon.config.js: %v\n", err)
		return
	}

	fmt.Println("New project structure created successfully.")
}

// newChatCmd
var newChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Create a new chat",
	Run:   newChat,
}

func newChat(cmd *cobra.Command, args []string) {
	// Implement the new chat logic here
}

// chatCmd
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Interact with a chat",
	Run:   chat,
}

func chat(cmd *cobra.Command, args []string) {
	// Implement the chat logic here
}

// phaseCmd
var phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "Manage phases",
	Run:   phase,
}

func phase(cmd *cobra.Command, args []string) {
	from, _ := cmd.Flags().GetInt("from")
	to, _ := cmd.Flags().GetInt("to")
	parentFolder, _ := cmd.Flags().GetString("parentFolder")
	if parentFolder == "" {
		parentFolder = "."
	}

	if from == 0 || to == 0 {
		fmt.Println("Please provide --from and --to flags with valid phase numbers (1, 2, 3, 4)")
		return
	}

	config, err := ReadConfig("moon.config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Filter commands based on --from and --to flags
	filteredCommands := filterCommands(config.Commands, from, to)

	// Display the filtered commands and execute the selected one
	selectedCommand := displayCommands(filteredCommands)
	if selectedCommand == nil {
		fmt.Println("No command selected")
		return
	}

	strTo := strconv.Itoa(to)

	executeCommand(selectedCommand, parentFolder, strTo)
}

// orbitCmd
var orbitCmd = &cobra.Command{
	Use:   "orbit",
	Short: "Orbit options",
	Run:   orbit,
}

func orbit(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a valid orbit number (1, 2, 3, 4)")
		return
	}

	// flag --parentFolder optional current folder if not provided
	parentFolder, _ := cmd.Flags().GetString("parentFolder")
	if parentFolder == "" {
		parentFolder = "."
	}

	number, err := strconv.Atoi(args[0])
	if err != nil || (number < 1 || number > 4) {
		fmt.Println("Please provide a valid orbit number (1, 2, 3, 4)")
		return
	}

	config, err := ReadConfig("moon.config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Filter commands based on the orbit number
	filteredCommands := filterOrbitCommands(config.Commands, number)

	// Display the filtered commands and execute the selected one
	selectedCommand := displayCommands(filteredCommands)
	if selectedCommand == nil {
		fmt.Println("No command selected")
		return
	}

	executeCommand(selectedCommand, parentFolder, args[0])
}

func init() {
	newChatCmd.Flags().Int("start", 0, "Start phase (1, 2, 3, 4)")
	newChatCmd.Flags().Bool("insert", false, "Open a select to choose files to put into chat")
	chatCmd.Flags().String("file", "", "File to chat with")
	phaseCmd.Flags().Int("from", 0, "From phase (1, 2, 3, 4)")
	phaseCmd.Flags().Int("to", 0, "To phase (1, 2, 3, 4)")
	phaseCmd.Flags().String("parentFolder", "", "Parent folder")
	orbitCmd.Flags().Int("number", 0, "Orbit number (1, 2, 3, 4)")
	orbitCmd.Flags().String("parentFolder", "", "Parent folder")
}

func filterCommands(commands []Command, from, to int) []Command {
	var filtered []Command
	for _, cmd := range commands {
		if cmd.From == from && cmd.To == to {
			filtered = append(filtered, cmd)
		}
	}
	return filtered
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
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return &commands[index]
}

func executeCommand(command *Command, parentFolder string, phase string) {
	// Get user inputs and put them into the command template
	interpolatedCommand, interpolatedTitle := interpolateCommand(command.Command, parentFolder)

	// println(interpolatedCommand)
	// Call the LLM API (or any other external function)
	// This is a placeholder function and should be replaced with the actual API call
	// response := callLLM(interpolatedCommand)
	var fullres = ssereq(interpolatedCommand)

	println(fullres)
	// Generate the inferred title
	inferredTitle := time.Now().Format("20060102150405") + ".md"

	// Get the title from the LLM API
	summary := ssereq("Summarize the following into a file name: " + fullres)

	yaml := "---\n" +
		"title: " + summary + "\n" +
		"phase: " + phase + "\n" +
		"command: " + interpolatedTitle + "\n" +
		"time: " + time.Now().Format("2006-01-02 15:04:05") + "\n" +
		"---\n\n"

	if summary != "" {
		inferredTitle = summary + ".md"
	}

	// Save the content to a file with the inferred title
	saveToFile(inferredTitle, yaml+fullres, parentFolder, phase)
}

func interpolateCommand(command string, parentFolder string) (string, string) {
	userPromptRegex := regexp.MustCompile(`\{user_prompt\}`)
	phasePickerRegex := regexp.MustCompile(`\{phase_(\d+)__file_picker\}`)

	titleString := command

	for {
		userPromptMatch := userPromptRegex.FindStringIndex(command)
		phasePickerMatch := phasePickerRegex.FindStringSubmatchIndex(command)

		titleUserPromptMatch := userPromptRegex.FindStringIndex(titleString)
		titlePhasePickerMatch := phasePickerRegex.FindStringSubmatchIndex(titleString)

		if userPromptMatch == nil && phasePickerMatch == nil {
			break
		}

		if userPromptMatch != nil && (phasePickerMatch == nil || userPromptMatch[0] < phasePickerMatch[0]) {
			userInput := promptUserInput()
			command = command[:userPromptMatch[0]] + userInput + command[userPromptMatch[1]:]
			titleString = titleString[:titleUserPromptMatch[0]] + userInput + titleString[titleUserPromptMatch[1]:]
		} else if phasePickerMatch != nil {
			phaseNumber := command[phasePickerMatch[2]:phasePickerMatch[3]]
			selectedFile, selectedFilePath := promptPhaseFilePicker(phaseNumber, parentFolder)
			// command = command[:phasePickerMatch[0]] + selectedFile + command[phasePickerMatch[1]:]

			content, err := ioutil.ReadFile(selectedFilePath)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return "", ""
			}

			command = command[:phasePickerMatch[0]] + string(content) + command[phasePickerMatch[1]:]

			// replace {phase_NUMBER__file_picker} with the selected file name
			titleString = titleString[:titlePhasePickerMatch[0]] + selectedFile + titleString[titlePhasePickerMatch[1]:]
		}
	}

	return command, titleString
}

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

	return selectedFile, filepath.Join(phaseFolderPath, selectedFile)
}

func phaseFolderName(phaseNumber string) string {
	switch phaseNumber {
	case "0":
		return "🌑"
	case "1":
		return "🌒"
	case "2":
		return "🌓"
	case "3":
		return "🌔"
	case "4":
		return "🌕"
	default:
		return ""
	}
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
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func saveToFile(title string, content string, parentFolder string, phaseNumber string) {
	folderName := phaseFolderName(phaseNumber)
	if folderName == "" {
		fmt.Printf("Invalid phase number: %s\n", phaseNumber)
		return
	}

	filePath := filepath.Join(parentFolder, folderName, title)
	println(filePath)

	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Printf("File saved as %s\n", filePath)
}
