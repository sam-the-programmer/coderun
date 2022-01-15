package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`Language Runner CLI
Based off of the VS Code extensions language runner,
this app allows you to run any code from one executable.

Run (on windows and linux):
coderun <filename>

Run (on macos):
coderun-macos <filename>

The app supports...

	- C#
	- C++
	- Go
	- Javascript
	- Python
	- Typescript
	`)
		os.Exit(0)
	}
	args := os.Args[1:]
	filename := os.Args[1]

	if strings.HasSuffix(filename, ".py") {
		exec.Command("py", args...)
	} else if strings.HasSuffix(filename, ".go") {
		exec.Command("go run", args...)
	} else if strings.HasSuffix(filename, ".js") {
		exec.Command("node", args...)

	} else if strings.HasSuffix(filename, ".ts") {
		fmt.Print("\033[33m \bCompiling Typescript File...\033[32m")
		fmt.Print("\rTypescript File Compiled!   \033[0m\n\n")

		run([]string{"tsc", filename})

		compilerArgs := getCompilerArgs("node", filename, "ts", "js", args)
		run(compilerArgs)

	} else if strings.HasSuffix(filename, ".cpp") {
		fmt.Print("\033[33m \bCompiling C++ File...\033[32m")
		fmt.Print("\rC++ File Compiled!   \033[0m\n\n")

		run([]string{"tsc", filename})

		compilerArgs := getCompilerArgs("", filename, "cpp", "exe", args)
		run(compilerArgs)

	} else if strings.HasSuffix(filename, ".cs") {
		fmt.Print("\033[33m \bCompiling C# File...\033[32m")
		fmt.Print("\rC# File Compiled!   \033[0m\n\n")

		run([]string{"tsc", filename})

		compilerArgs := getCompilerArgs("", filename, "cs", "exe", args)
		run(compilerArgs)

	} else {
		fmt.Println("\033[31m \bThe file", filename, "could not be read.", "\033[0m")
	}
}

func getCompilerArgs(exe string, filename string, oldext string, newext string, args []string) []string {
	var compilerArgs []string

	if exe != "" {
		compilerArgs = []string{exe, strings.TrimSuffix(filename, oldext) + newext}
		if len(args) > 1 {
			for _, v := range args {
				compilerArgs = append(compilerArgs, v)
			}
		}
	} else {
		compilerArgs = []string{strings.TrimSuffix(filename, oldext) + newext}
		if len(args) > 1 {
			for _, v := range args {
				compilerArgs = append(compilerArgs, v)
			}
		}
	}
	return compilerArgs
}

func run(args []string) {
	cmd := exec.Command(args[0], args[1:]...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	go runAsync(stdout)
	go runAsync(stderr)

	cmd.Wait()
}

func runAsync(pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		fmt.Print(m, "\n")
	}
}
