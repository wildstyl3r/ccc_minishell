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
		argv := strings.Fields(line)
		if len(argv) > 0 {
			switch argv[0] {
			case "exit":
				exitCode := 0
				if len(argv) > 1 {
					exitCode, _ = strconv.Atoi(argv[1])
				}
				os.Exit(exitCode)
			default:
				fmt.Fprintln(os.Stdout, argv[0]+": command not found")
			}
		}

	}
}
