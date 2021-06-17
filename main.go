package main

import (
	"fmt"
	"regexp"
	"os"
	"io/ioutil"
)

func main() {
	// read source code from stdin
	source, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	find := os.Args[1]
	replace := os.Args[2] 

	// replace first command line argument with the second
	var re = regexp.MustCompile(find)
	s := re.ReplaceAllString(string(source), replace)
	fmt.Print(s)
}

