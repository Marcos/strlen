package main

import (
	"bytes"
	"flag"
	"fmt"
	"sort"
	"strconv"
)

func showVerboseResult(str string) {
	chars := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		char := str[i]
		totalCount, hasChar := chars[char]
		count := 1
		if hasChar {
			count = totalCount + 1
		}
		chars[char] = count
	}

	keys := make([]string, 0, len(chars))
	for k := range chars {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k + ": " + strconv.Itoa(chars[[]byte(k)[0]]))
	}
}

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
		showVerboseResult(str)
		buffer.WriteString(str + ": ")
	}

	buffer.WriteString(size)
	fmt.Println(buffer.String())
}
