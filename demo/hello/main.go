package main

import (
	"fmt"
	"lib"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		lib.HelloLoop(args[1:])
	} else {
		fmt.Println(lib.Hello(nil))
	}
}
