package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	base := 200.0
	if len(os.Args) > 1 {
		if v, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			base = v
		}
	}
	var b strings.Builder
	b.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"500\" height=\"100\">\n")
	for i := 0; i < 5; i++ {
		h := base + float64(i)*30
		if h >= 360 {
			h -= 360
		}
		r, g, blue := hslToRgb(h, 0.7, 0.5)
		fmt.Fprintf(&b, "<rect x=\"%d\" y=\"0\" width=\"100\" height=\"100\" fill=\"rgb(%d,%d,%d)\"/>\n", i*100, r, g, blue)
	}
	b.WriteString("</svg>\n")
	fmt.Print(b.String())
}

func hslToRgb(h, s, l float64) (int, int, int) {
	c := (1.0 - abs(2.0*l-1.0)) * s
	x := c * (1.0 - abs(fmod(h/60.0, 2.0)-1.0))
	m := l - c/2.0
	var r, g, b float64
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}
	return int((r+m)*255 + 0.5), int((g+m)*255 + 0.5), int((b+m)*255 + 0.5)
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func fmod(a, b float64) float64 {
	for a >= b {
		a -= b
	}
	return a
}
