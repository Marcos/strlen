package main

import (
	"bytes"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	verbose := flag.Bool("v", false, "Display verbose size")
	flag.Parse()
	args := flag.Args()

	var str, size string = "", "0"
	if len(args) > 0 {
		str = args[0]
		size = strconv.Itoa(len(str))
	}

	var buffer bytes.Buffer
	if *verbose {
		buffer.WriteString(str)
		buffer.WriteString(": ")
	}

	buffer.WriteString(size)
	fmt.Println(buffer.String())
}
