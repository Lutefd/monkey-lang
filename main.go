package main

import (
	"fmt"
	"strings"

	"github.com/lutefd/monkey-lang/lexer"
)

func main () {
	testStr := "test ==	 aa"
	l := lexer.New(testStr)
	testFields := strings.Fields(testStr)
	for _, test := range testFields{
		l.NextToken()
		fmt.Println(test)
	}
}