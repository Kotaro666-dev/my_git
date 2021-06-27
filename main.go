package main

import (
	"fmt"
	"github.com/Kotaro666-dev/my_git/srcs/add"
	"github.com/Kotaro666-dev/my_git/srcs/commit"
	"github.com/Kotaro666-dev/my_git/srcs/log"
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
		add.Add()
	case "commit":
		commit.Commit()
	case "log":
		log.Log()
	default:
		fmt.Printf("my_git: '%s' is not a git command. See 'git --help'.\n", arg)
		os.Exit(1)
	}
}
