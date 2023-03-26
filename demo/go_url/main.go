package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func main() {
	args := os.Args
	en := path.Base(args[0])
	if len(args) == 2 {
		arg := args[1]
		p, e := url.ParseRequestURI(arg)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Printf("%s, url; %s \n", en, p)
		r, e2 := http.Get(arg)
		if e2 != nil {
			log.Fatal(e2)
		}
		fmt.Printf("%s, response status; %d \n", en, r.StatusCode)
		defer r.Body.Close()
		body, e3 := io.ReadAll(r.Body)
		if e3 != nil {
			log.Fatal(e3)
		}
		// fmt.Printf("%s, response bytes; \n%v \n", en, body)
		fmt.Printf("%s, response body; \n%s \n", en, body)
	} else {
		fmt.Printf("Usage: %s `url-to-parse` \n", en)
	}
}
