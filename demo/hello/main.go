package main

import (
	"fmt"
	"github.com/tbh26/g3/demo/lib"
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
