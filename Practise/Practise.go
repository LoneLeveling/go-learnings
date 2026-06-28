package main

import "fmt"

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

func main() {
	c := Circle{x: 10, y: 20, radius: 80, area: 0}

	fmt.Printf("%+v\n", c)

	calcArea(&c)

	fmt.Printf("%+v\n", c)
}

func calcArea(c *Circle) {
	var area float64

	area = float64(c.x) * float64(c.y) * c.radius

	c.area = area // same as (*c).area = area

	// or:
	// (*c).area = area
}