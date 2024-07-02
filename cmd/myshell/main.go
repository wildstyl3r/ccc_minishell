package main

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
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
		if len(command) > 0 {
			commandPath := ""
			for _, p := range pathv {
				if _, err := os.Stat(p + "/" + command); err == nil {
					commandPath = p + "/" + command
					break
				}
			}
			if commandPath != "" {
				cmd := exec.Command(commandPath, args)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				err = cmd.Run()
				if err != nil {
					fmt.Fprintln(os.Stdout, "error: ", err)
				}
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
				builtinSet := map[string]struct{}{"echo": {}, "exit": {}, "type": {}, "pwd": {}}
				argv := strings.Fields(args)
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
			default:
				fmt.Fprintln(os.Stdout, command+": command not found")
			}
		}

	}
}
