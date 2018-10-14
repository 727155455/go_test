package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goinaction/code/chapter3/words"
)

func main() {
	fmt.Printf("os.arg is %+v", os.Args)
	filename := os.Args[1]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("content is %d", content)
	test := string(content)

	count := words.CountWords(test)
	fmt.Printf("%d words\n", count)
}
