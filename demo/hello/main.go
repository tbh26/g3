package main

import (
	"fmt"
	"github.com/tbh26/g3/demo/hello"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		hello.HelloLoop(args[1:])
	} else {
		fmt.Println(hello.Hello(nil))
	}
}
