package main

import (
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length == 1 {
		fmt.Println("git help")
		os.Exit(1)
	}
	arg := os.Args[1]
	switch arg {
		case "add":
			fmt.Println("my_git add")
		case "commit":
			fmt.Println("my_git commit")
		default:
			fmt.Printf("my_git: '%s' is not a git command. See 'git --help'.\n", arg)
			os.Exit(1)
	}
}
