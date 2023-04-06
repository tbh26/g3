package main

import (
	"fmt"
	"os"
)

func hello(sr *string) (result string) {
	if sr == nil || *sr == "" {
		result = "hello world"
	} else {
		result = fmt.Sprintf("hello %s", *sr)
	}
	return
}

func main() {
	args := os.Args
	if len(args) > 1 {
		for _, a := range args[1:] {
			fmt.Println(hello(&a))
		}
	} else {
		fmt.Println(hello(nil))
	}
}
