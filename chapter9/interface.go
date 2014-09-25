package main

import "fmt"
import "math"

type Shape interface {
     area() float64
}

type Circle struct {
     x, y, r float64
}

type Rectangle struct {
     x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
     return math.Pi * c.r*c.r
}

func distance(x1, y1, x2, y2 float64) (area float64) {
     a := x2 - x1
     b := y2 - y1
     return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
     l := distance(r.x1, r.y1, r.x1, r.y2)
     w := distance(r.x1, r.y1, r.x2, r.y1)
     return l * w
}

func totalArea(shapes ...Shape) float64 {
     var area float64
     for _, s := range shapes {
     	 area += s.area()
     }
     return area
}

func main() {
     r := Rectangle{0, 0, 10, 10}
     c := Circle{0, 0, 5}
     fmt.Println(totalArea(&c, &r))
}