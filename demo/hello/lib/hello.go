package lib

import (
	"fmt"
	"strconv"
)

func Hello(sr *string) (result string) {
	if sr == nil || *sr == "" {
		result = "lib world"
	} else {
		result = fmt.Sprintf("lib %s", *sr)
	}
	return
}

func HelloInt(n int) (result string) {
	result = fmt.Sprintf("lib number %d", n)
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
