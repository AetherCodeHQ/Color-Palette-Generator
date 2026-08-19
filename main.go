package main

import (
	"fmt"
	"os"
)

// color_palette_generator - Generate color palettes
func color_palette_generator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Color-Palette-Generator")
	fmt.Println("  Generate color palettes")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	color_palette_generator(path)
}
