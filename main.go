package main

import (
	"fmt"
	"github.com/Kotaro666-dev/my_git/srcs/my_init"
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
	case "init":
		my_init.Init()
	case "add":
		fmt.Println("my_git add")
	case "commit":
		fmt.Println("my_git commit")
	default:
		fmt.Printf("my_git: '%s' is not a git command. See 'git --help'.\n", arg)
		os.Exit(1)
	}
}
