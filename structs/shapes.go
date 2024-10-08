package structs

import "math"

func Perimeter(rectangle Rectangle) float64 {
  return 2 * (rectangle.Width + rectangle.Height) 
}

type Rectangle struct {
  Width float64
  Height float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

type Circle struct {
  Radius float64
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
  Base float64
  Heigth float64
}

func (t Triangle) Area() float64 {
  return (t.Base * t.Heigth) / 2
}

type Shape interface {
  Area() float64
}
