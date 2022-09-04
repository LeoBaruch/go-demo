package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-q.X)
}

func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y}
}

func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

func (p Point) Print() {
	fmt.Printf("{%f, %f}\n", p.X, p.Y)
}

type Path []Point

func (path Path) TranslateBy(another Point, add bool) {
	var op func(p, q Point) Point

	if add == true {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], another)
		path[i].Print()
	}
}

func main() {
	points := Path{
		{10, 10},
		{11, 11},
	}

	anotherPoint := Point{5, 5}

	points.TranslateBy(anotherPoint, false)

	fmt.Println("==================")

	points.TranslateBy(anotherPoint, true)
}
