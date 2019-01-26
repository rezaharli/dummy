package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kharism/dummy/dblayer"
)

var (
	CarStorage dblayer.IStorage
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	executor := NewCommandExecutor()
	for {
		//fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		//fmt.Println(text)
		out := executor.ExecuteCommand(text)
		fmt.Println(out)
	}

}
