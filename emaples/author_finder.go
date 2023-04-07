package main

import (
	"fmt"

	"github.com/tomba-io/go/tomba"
)

func main() {
	tomba := tomba.New("ta_xxxxx", "ts_xxxxx")

	result, err := tomba.AuthorFinder("https://clearbit.com/blog/company-name-to-domain-api")
	if err == nil {
		fmt.Println(result)
	}
}