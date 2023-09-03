package main //main keyword creates executeable when we run go build.

import (
	"fmt"

	"rsc.io/quote"
)

func maine() { //main func is needed in every main package
	fmt.Println(quote.Glass())
}
