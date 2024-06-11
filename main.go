package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	art "ascii/functions"
)

var banners = map[string]string{
	"standard":   "standard.txt",
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no arguments passed")
		return
	}
	if args[0] == "" {
		return
	}
	if args[0] == "\\n" {
		fmt.Println()
		return
	}
	var banner string
	if len(args) > 1 {
		banner = args[1]
	} else {
		banner = "standard"
	}
	filename, ok := banners[banner]
	if !ok {
		log.Fatal("invalid banner file")
	}
	input := args[0]
	 printable := art.NonPrintable(input)
	lines := strings.Split(printable,"\ng")
	result := art.ProcessLine(lines)
	result2 := art.ProcessString(result,filename)
	art.PrintStrings(result2)
	fmt.Println()
}
