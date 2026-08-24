
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: Color-Palette-Generator <bytes|lines|reverse|upper|lower>")
		os.Exit(1)
	}
	mode := os.Args[1]
	in := ""
	if len(os.Args) > 2 {
		in = os.Args[2]
	}
	lines := strings.Split(in, "\n")
	switch mode {
	case "lines":
		fmt.Println(len(lines))
	case "bytes":
		fmt.Println(len(in))
	case "upper":
		fmt.Println(strings.ToUpper(in))
	case "lower":
		fmt.Println(strings.ToLower(in))
	case "reverse":
		runes := []rune(in)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		fmt.Println(string(runes))
	default:
		fmt.Fprintln(os.Stderr, "unknown mode:", mode)
		os.Exit(1)
	}
}
