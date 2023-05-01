To build a CLI app as described, follow these steps:

1. Create the main function for the CLI app
2. Initialize Cobra and Promptui packages
3. Define commands and their behavior
4. Implement the folder structure
5. Define and call the llm-engine() function
6. Implement git commit functionality
7. Create a configuration file for prompts and settings

Here's a sample algorithm to guide you through the process:

```
Algorithm Build_CLI_App

1. Initialize Cobra and Promptui packages
2. Define commands and their behavior:
   a. create - create a new folder with four phases
   b. --1, --2, --3, --4 - go to specific phase and commit
   c. config - edit prompts for phase shifts and edit git settings
   d. commit-all - commit all files

3. Implement the folder structure
   a. Create the required folders (phase0, phase1, p2, p3, p4) and markdown files
   b. For phase 3 and 4, create code files with the specified target language
   c. Create a config.local.json file

4. Define and call the llm-engine() function
   a. Accepts from and to phases as arguments
   b. Processes the command and returns the result

5. Implement git commit functionality
   a. After each command, make a git commit with the command as the commit message

6. Create a configuration file for prompts and settings
   a. Specify target language (js, ts, py, go)
   b. Define commands with their properties (from, to, command, prompt, name, description)

7. Run the CLI app from any directory
```

After defining the algorithm, use Go to implement the CLI app. Start by initializing the necessary packages, define the required commands and their behavior, create the folder structure, implement the llm-engine() function, and integrate git commit functionality. Finally, create a configuration file to manage prompts and settings.