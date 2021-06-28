package add

import (
	"fmt"
	"log"
	"os"
)

//func doesFileExist(name string) bool {
//	if _, err := os.Stat(name); err != nil {
//		if os.IsNotExist(err) {
//			return false
//		}
//	}
//	return true
//}

func writeIndex() {
	indexPath := ".my_git/index"
	file, err := os.OpenFile(indexPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.Write([]byte("test\n")); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func Add(args []string) {
	fmt.Println("my_git add")
	writeIndex()
	var i int = 2
	for i < len(args) {
		//fmt.Printf("args[%d] -> %s\n", i, args[i])
		// ファイルを読み込み

		//
		i++
	}
}
