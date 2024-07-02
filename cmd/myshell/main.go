package main

import (
	"bufio"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

func main() {
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
				builtinSet := map[string]struct{}{"echo": {}, "exit": {}, "type": {}}
				argv := strings.Fields(args)
				for _, arg := range argv {
					if _, contains := builtinSet[arg]; contains {
						fmt.Fprintln(os.Stdout, arg+" is a shell builtin")
					} else {
						fmt.Fprintln(os.Stdout, arg+": not found")
					}
				}
			default:
				fmt.Fprintln(os.Stdout, command+": command not found")
			}
		}

	}
}
