package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Config struct {
	TargetLanguage string `json:"targetLanguage"`
}

func main() {
	var configFile string

	app := &cobra.Command{
		Use:   "llm-cli",
		Short: "CLI app for managing the learning lifecycle",
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Run: func(cmd *cobra.Command, args []string) {
			initProject()
		},
	}

	commitCmd := &cobra.Command{
		Use:   "commit [phase]",
		Short: "Commit files of the specified phase",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			commit(args[0])
		},
	}

	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Edit configuration settings",
		Run: func(cmd *cobra.Command, args []string) {
			editConfig(&configFile)
		},
	}

	commitAllCmd := &cobra.Command{
		Use:   "commit-all",
		Short: "Commit all files",
		Run: func(cmd *cobra.Command, args []string) {
			commitAll()
		},
	}

	readfileCmd := &cobra.Command{
		Use:   "readfile [file]",
		Short: "Read the current file and execute user-selected commands",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			readFileAndExecuteCommand(filename)
		},
	}

	listFilesCmd := &cobra.Command{
		Use:   "listfiles [phase]",
		Short: "List files in the current directory or the specified phase folder",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var folder string
			if len(args) > 0 {
				folder = "phase" + args[0]
			} else {
				folder = "."
			}
			listFilesAndFolders(folder)
		},
	}

	app.PersistentFlags().StringVarP(&configFile, "config", "c", "config.local.json", "configuration file")
	app.AddCommand(initCmd, commitCmd, configCmd, commitAllCmd, readfileCmd, listFilesCmd)

	app.Execute()
}

func initProject() {
	// create folders and files
	os.Mkdir("phase0", os.ModePerm)
	os.Mkdir("phase1", os.ModePerm)
	os.Mkdir("phase2", os.ModePerm)
	os.Mkdir("phase3", os.ModePerm)
	os.Mkdir("phase4", os.ModePerm)

	// create config file
	config := Config{TargetLanguage: "go"}
	file, _ := json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile("config.local.json", file, 0644)
}

func commit(phase string) {
	cmd := exec.Command("git", "add", ".")
	cmd.Run()

	commitMsg := fmt.Sprintf("Commit phase %s", phase)
	cmd = exec.Command("git", "commit", "-m", commitMsg)
	cmd.Run()
}

func editConfig(configFile *string) {
	config := &Config{}
	file, _ := ioutil.ReadFile(*configFile)
	_ = json.Unmarshal(file, config)

	prompt := promptui.Select{
		Label: "Select Target Language",
		Items: []string{"go", "js", "ts", "py"},
	}

	_, language, _ := prompt.Run()
	config.TargetLanguage = language

	newFile, _ := json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile(*configFile, newFile, 0644)
}

func commitAll() {
	cmd := exec.Command("git", "add", ".")
	cmd.Run()

	cmd = exec.Command("git", "commit", "-m", "Commit all phases")
	cmd.Run()
}

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

type FileNode struct {
	Name     string
	Children []*FileNode
}

func listFilesAndFolders(folder string) {
	rootNode, err := buildFileTree(folder, "")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	selectedNode := selectFileNode(rootNode)
	if selectedNode != nil {
		fmt.Println("Selected file:", selectedNode.Name)
	}
}

func buildFileTree(path, prefix string) (*FileNode, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	node := &FileNode{Name: prefix + fileInfo.Name()}
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			childNode, err := buildFileTree(filepath.Join(path, file.Name()), "  ")
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, childNode)
		}
	}

	return node, nil
}

func selectFileNode(node *FileNode) *FileNode {
	if node == nil || len(node.Children) == 0 {
		return node
	}

	prompt := promptui.Select{
		Label: "Select a file",
		Items: node.Children,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U0001F336 {{ .Name | cyan }}",
			Inactive: "  {{ .Name | cyan }}",
			Selected: "{{ .Name | green | bold }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return selectFileNode(node.Children[index])
}

func listFilesInFolder(folder string) {
	// Read the folder content
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Collect file names
	fileNames := []string{}
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	// Show dropdown with file names
	prompt := promptui.Select{
		Label: "Select a file",
		Items: fileNames,
	}

	_, fileName, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the selected file name
	fmt.Println("Selected file:", fileName)
}
