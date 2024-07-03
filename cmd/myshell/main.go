package main

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"

	"fmt"
	"os"
)

func main() {
	path, foundPath := os.LookupEnv("PATH")
	var pathv []string
	if foundPath {
		pathv = strings.Split(path, ":")
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			fmt.Fprintln(os.Stdout, "error: ", err)
		}
		command, args, foundArgs := strings.Cut(line, " ")
		command = strings.TrimSpace(command)
		args = strings.TrimSpace(args)
		argv := strings.Fields(args)
		if len(command) > 0 {
			commandPath := ""
			for _, p := range pathv {
				if _, err := os.Stat(p + "/" + command); err == nil {
					commandPath = p + "/" + command
					break
				}
			}
			if commandPath != "" {
				cmd := exec.Command(commandPath, argv...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				continue
			}

			switch command {
			case "exit":
				exitCode := 0
				if foundArgs {
					exitCode, _ = strconv.Atoi(args)
				}
				os.Exit(exitCode)
			case "echo":
				fmt.Fprintln(os.Stdout, strings.TrimSpace(args))
			case "type":
				builtinSet := map[string]struct{}{"echo": {}, "exit": {}, "type": {}, "pwd": {}, "cd": {}}
				for _, arg := range argv {
					if _, contains := builtinSet[arg]; contains {
						fmt.Fprintln(os.Stdout, arg+" is a shell builtin")
						continue
					}
					commandPath := ""
					for _, p := range pathv {
						if _, err := os.Stat(p + "/" + arg); err == nil {
							commandPath = p + "/" + arg
							break
						}
					}
					if commandPath != "" {
						fmt.Fprintln(os.Stdout, arg+" is "+commandPath)
					} else {
						fmt.Fprintln(os.Stdout, arg+": not found")
					}
				}
			case "pwd":
				wd, err := os.Getwd()
				if err != nil {
					fmt.Fprintln(os.Stdout, "error: ", err)
				}
				fmt.Fprintln(os.Stdout, wd)
			case "cd":
				if args == "~" {
					home, err := os.UserHomeDir()
					if err != nil {
						fmt.Fprintln(os.Stdout, "error: ", err)
					}
					os.Chdir(home)
				} else {
					if _, err := os.Stat(args); err == nil {
						os.Chdir(args)
					} else {
						fmt.Fprintln(os.Stdout, "cd: "+args+": No such file or directory")
					}
				}
			default:
				fmt.Fprintln(os.Stdout, command+": command not found")
			}
		}

	}
}
