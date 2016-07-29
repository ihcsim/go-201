package main

import (
	"fmt"
	"math"
)

// Geometry can perform the basic geometry operations of calculating the perimeter and area of polygons and circle.
type Geometry interface {
	// Perimeter calculates the perimeter of the implementor.
	Perimeter() float64

	// Area calculates the area of the implementor.
	Area() float64
}

// Circle represents a circle.
type Circle struct {
	Geometry
	radius float64
}

// Perimeter calculates the perimeter of c.
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Area calculates the area of c.
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// String returns the string representation of c.
func (c *Circle) String() string {
	return fmt.Sprintf("circle, with radius of %.2f, having a perimeter of %.2f and an area of %.2f", c.radius, c.Perimeter(), c.Area())
}

// Triangle represents a triangle.
type Triangle struct {
	Geometry
	side   float64
	base   float64
	height float64
}

// Perimeter calculates the perimeter of t.
func (t *Triangle) Perimeter() float64 {
	return t.base + 2*t.side
}

// Area calculates the area of t.
func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// String returns the string representation of t.
func (t *Triangle) String() string {
	return fmt.Sprintf("triangle, with base of %.2f, height of %.2f and side of %.2f, having a perimeter of %.2f and an area of %.2f", t.base, t.height, t.side, t.Perimeter(), t.Area())
}

// Rectangle represents a rectangle.
type Rectangle struct {
	Geometry
	length float64
	width  float64
}

// Perimeter calculates the perimeter of r.
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Area calculates the area of r.
func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

// String returns the string representation of r.
func (r *Rectangle) String() string {
	return fmt.Sprintf("rectangle, with length of %.2f and width of %.2f, having a perimeter of %.2f and an area of %.2f", r.length, r.width, r.Perimeter(), r.Area())
}

// Square represents a square.
type Square struct {
	Geometry
	side float64
}

// Perimeter calculates the perimeter of s.
func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

// Area calculates the area of s.
func (s *Square) Area() float64 {
	return math.Pow(s.side, 2)
}

// String returns the string representation of s.
func (s *Square) String() string {
	return fmt.Sprintf("square, with side of %.2f, having a perimeter of %.2f and an area of %.2f", s.side, s.Perimeter(), s.Area())
}

// Pentagon represents a pentagon.
type Pentagon struct {
	Geometry
	side float64
}

// Perimeter calculates the perimeter of p.
func (p *Pentagon) Perimeter() float64 {
	return 5 * p.side
}

// Area calculates the area of p.
func (p *Pentagon) Area() float64 {
	return 0.25 * math.Sqrt(5*(5+2*math.Sqrt(5))) * math.Pow(p.side, 2)
}

// String returns the string representation of p.
func (p *Pentagon) String() string {
	return fmt.Sprintf("pentagon, with side of %.2f, having a perimeter of %.2f and an area of %.2f", p.side, p.Perimeter(), p.Area())
}
