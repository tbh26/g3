package main

import (
	"fmt"
	"strconv"
)

func Hello(sr *string) (result string) {
	if sr == nil || *sr == "" {
		result = "hello world"
	} else {
		result = fmt.Sprintf("hello %s", *sr)
	}
	return
}

func HelloInt(n int) (result string) {
	result = fmt.Sprintf("hello number %d", n)
	return
}

func HelloLoop(args []string) {
	for _, a := range args {
		n, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(Hello(&a))
		} else {
			fmt.Println(HelloInt(n))
		}
	}
}
