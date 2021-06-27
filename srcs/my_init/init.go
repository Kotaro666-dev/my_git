package my_init

import (
	"fmt"
	"os"
)

func Init() {
	currentDir, _ := os.Getwd()
	if err := os.Mkdir(".my_git", 0777); err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll(".my_git/index", 0777); err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll(".my_git/refs/main", 0777); err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll(".my_git/objects", 0777); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Initialized empty Git repository in %s/.my_git/\n", currentDir)
}
