package main

import (
	"bufio"
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
			fmt.Fprintln(os.Stdout, argv[0]+": command not found")
		}

	}
}
