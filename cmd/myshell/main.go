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
	cmd := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
	}

	for true {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Printf("Something went wrong when reading file: %#v\n", err)
			return
		}

		command = strings.TrimSpace(command)
		chainedCmd := strings.Split(command, " ")

		if cmd[chainedCmd[0]] {
			val := executeCmd(chainedCmd[0], chainedCmd[1:], cmd)
			if val {
				break
			}
			continue
		}

		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func executeCmd(cmd string, attr []string, cmdMap map[string]bool) bool {
	switch cmd {
	case "echo":
		fmt.Fprintf(os.Stdout, strings.Join(attr, " "))
		fmt.Fprintf(os.Stdout, "\n")
		return false
	case "exit":
		return true
	case "type":
		builtinCmd := attr[0]
		if cmdMap[builtinCmd] {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", builtinCmd)
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", builtinCmd)
		}

	}

	return false
}
