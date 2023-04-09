package main

import (
	"fmt"
	"os"
	"strconv"
)

func hello(sr *string) (result string) {
	if sr == nil || *sr == "" {
		result = "hello world"
	} else {
		result = fmt.Sprintf("hello %s", *sr)
	}
	return
}

func helloInt(n int) (result string) {
	result = fmt.Sprintf("hello number %d", n)
	return
}

func helloLoop(args []string) {
	for _, a := range args {
		n, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(hello(&a))
		} else {
			fmt.Println(helloInt(n))
		}
	}
}

func main() {
	args := os.Args
	if len(args) > 1 {
		helloLoop(args[1:])
	} else {
		fmt.Println(hello(nil))
	}
}
