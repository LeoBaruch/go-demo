package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance1 := Point.Distance

	fmt.Println(distance1(p, q))
	fmt.Println((p.Distance(q)))

	fmt.Printf("%T\n", distance1)

	fmt.Println("==============")

	distance2 := (*Point).Distance
	fmt.Println(distance2(&p, q))
	fmt.Printf("%T\n", distance2)

}
