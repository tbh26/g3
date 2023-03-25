package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println(" hello :")
		for _, a := range args[1:] {
			fmt.Printf("\t %s \n", a)
		}
	} else {
		fmt.Println("hello world")
	}
}
