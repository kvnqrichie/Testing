package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	output := [15]rune{
		'x', ' ', '=', ' ',
		'0', '0',
		',', ' ', 'y', ' ', '=', ' ',
		'0', '0',
		'\n',
	}

	for i := 0; i < points.x/10; i++ {
		output[4]++
	}
	for i := 0; i < points.x%10; i++ {
		output[5]++
	}
	for i := 0; i < points.y/10; i++ {
		output[12]++
	}
	for i := 0; i < points.y%10; i++ {
		output[13]++
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
