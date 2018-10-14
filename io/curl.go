package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init(){
	if len := len(os.Args); len != 2 {
		fmt.Println("params counts error, %d", len)
		os.Exit(-1)
	}
}

func main(){
	content, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("error %+v", err)
		return
	}

	io.Copy(os.Stdout, content.Body)
	if err := content.Body.Close(); err != nil {
		fmt.Println("aaaaaaa")
	}


}