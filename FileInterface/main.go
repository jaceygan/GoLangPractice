package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tempBytes, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {

		io.Copy(os.Stdout, tempBytes)
	}

}
