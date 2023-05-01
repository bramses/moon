I'm building a CLI app using promptui and cobra (Go) that will:

- have five phase folders
    1. phase 0: pre concrete idea phase, many competing ideas
    2. phase 1: concrete idea phase, one idea
    3. phase 2: psuedocode/ algorithm phase
    4. phase 3: concrete code phase -- programming. language; draft state
    5. phase 4: final state, ready to deploy, supplemental environment features
- each phase folder will have a markdown file that has data about the phase. phases 3 and 4 will also have code files
- the CLI will look at the open file and use the command input by the user to transform it to another phase
- there are "intra" commands and "extra" commands -- interfacing within the phase and outside of the phase respectively
- each phase shift will have a dropdown of available commands conditional on the current phase
- the command will call `llm-engine()` with the from and to phases as arguments
- after each command fires a git commit will be made with the command as the commit message
- the CLI will be able to be run from any directory


## folder structure example

```
phase0
	phase0.md
phase1
	phase1.md
p2
	phase2.md
p3
	phase3.md
	phase3.{target language (inferred or explicit)}
p4
	phase4.md
	code
	tests
	docs
config.local.json
```


## cli example

```
moon -- options
	create - create a new folder with four phases
		-- target-language (js, ts, py, go
		)
	--1
		go to phase 1 (context specific options based on file selected) + commit
	--2
		go to phase 2 (context specific options based on file selected) + commit
	--3
		...
	--4
		...
	config
		edit prompts for phase shifts
		edit git settings
	commit-all
		commit all files

ex:

moon --2
	(current file is phase1.md)
	[ ] for app idea, write a list of features
	[ ] write the above as a psuedocode algorithm
	...

moon --2
	(current file is phase4.md)
	[ ] fix this structural issue (prompt)
		change the algorithm to not use a double for loop but a map instead
```

## llm prompt call example

```
llm_map(engine, prompt) => engine.call(prompt)
.then(write_file)
.then(commit => message)
```

### a prompt example

```
prompt example: for {phase2.md} fix the structural issue: {user_prompt}
```

### config example

```
config:
{
	target-language: "ts"
	commands: [
		{
			from: phase1,
			to: phase3,
			command: "blah blah {user_prompt} {phase3.md} {phase1.md}",
			prompt: true,
            name: "blah do a thing",
            description: "blah blah blah"
		}
	]
}
```