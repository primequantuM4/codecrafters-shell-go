package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for true {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Printf("Something went wrong when reading file: %#v\n", err)
			return
		}

		command = strings.TrimSpace(command)
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
