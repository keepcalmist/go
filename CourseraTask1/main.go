package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "strings"
)

var deep int = 0

func main() {
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("Usage go run main.go ./ [-f]")
	}
	path := os.Args[1]
	out := os.Stdout
	flag := len(os.Args) == 3 && os.Args[2] == "-f"
	fmt.Printf("go run main.go \n")
	if DirTree(path, out, flag) != nil {
		log.Fatal("Something did wrong...")
	}
}

func DirTree(path string, out *os.File, flag bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for i, _ := range files {
		if files[i].IsDir() {
			deep++
			for i := 0; i < deep-1; i++ {
				fmt.Fprintf(out, "\t")
			}
			fmt.Fprintf(out, "└───%s\n", files[i].Name())
			_ = DirTree(fmt.Sprintf("%s\\%s", path, files[i].Name()), out, flag)
		}
		if (!files[i].IsDir()) && flag == true {
			deep++
			for i := 0; i < deep-1; i++ {
				fmt.Fprintf(out, "\t")
			}
			if flag == true {
				fmt.Fprintf(out, "└───%s(.%db)\n", files[i].Name(), files[i].Size())
			}
		}
		deep--
	}

	return nil
}
