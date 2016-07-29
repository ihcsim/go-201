package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse parses b into a slice of geometries.
func Parse(b []byte) ([]Geometry, error) {
	results := make([]Geometry, 0)
	for _, s := range strings.Split(string(b), "\n") {
		if s == "" {
			continue
		}

		attr := strings.Split(s, ",")
		shape := strings.ToLower(attr[0])

		switch shape {
		case "square":
			s, err := parseSquare(attr)
			if err != nil {
				return results, err
			}
			results = append(results, s)
		case "pentagon":
			p, err := parsePentagon(attr)
			if err != nil {
				return results, err
			}
			results = append(results, p)
		case "triangle":
			t, err := parseTriangle(attr)
			if err != nil {
				return results, err
			}
			results = append(results, t)
		case "rectangle":
			r, err := parseRectangle(attr)
			if err != nil {
				return results, err
			}
			results = append(results, r)
		case "circle":
			c, err := parseCircle(attr)
			if err != nil {
				return results, err
			}
			results = append(results, c)
		default:
			return results, fmt.Errorf("Unsupported shape %q", shape)
		}
	}

	return results, nil
}

func parseSquare(attr []string) (*Square, error) {
	if len(attr) != 2 {
		return nil, fmt.Errorf("Bad square data: %v", attr)
	}

	side, err := strconv.ParseFloat(attr[1], 64)
	if err != nil {
		return nil, err
	}
	if side < 0 {
		return nil, fmt.Errorf("Side of a square cannot be negative")
	}

	return &Square{side: side}, nil
}

func parsePentagon(attr []string) (*Pentagon, error) {
	if len(attr) != 2 {
		return nil, fmt.Errorf("Bad pentagon data: %v", attr)
	}

	side, err := strconv.ParseFloat(attr[1], 64)
	if err != nil {
		return nil, err
	}
	if side < 0 {
		return nil, fmt.Errorf("Side of a pentagon cannot be negative")
	}

	return &Pentagon{side: side}, nil
}

func parseTriangle(attr []string) (*Triangle, error) {
	if len(attr) != 4 {
		return nil, fmt.Errorf("Bad triangle data: %v", attr)
	}

	base, err := strconv.ParseFloat(attr[1], 64)
	if err != nil {
		return nil, err
	}
	if base < 0 {
		return nil, fmt.Errorf("Base of a triangle cannot be negative")
	}

	height, err := strconv.ParseFloat(attr[2], 64)
	if err != nil {
		return nil, err
	}
	if height < 0 {
		return nil, fmt.Errorf("Height of a triangle cannot be negative")
	}

	side, err := strconv.ParseFloat(attr[3], 64)
	if err != nil {
		return nil, err
	}
	if side < 0 {
		return nil, fmt.Errorf("Side of a triangle cannot be negative")
	}

	return &Triangle{base: base, height: height, side: side}, nil
}

func parseRectangle(attr []string) (*Rectangle, error) {
	if len(attr) != 3 {
		return nil, fmt.Errorf("Bad rectangle data: %v", attr)
	}

	length, err := strconv.ParseFloat(attr[1], 64)
	if err != nil {
		return nil, err
	}
	if length < 0 {
		return nil, fmt.Errorf("Length of a rectangle cannot be negative")
	}

	width, err := strconv.ParseFloat(attr[2], 64)
	if err != nil {
		return nil, err
	}
	if width < 0 {
		return nil, fmt.Errorf("Width of a rectangle cannot be negative")
	}

	return &Rectangle{length: length, width: width}, nil
}

func parseCircle(attr []string) (*Circle, error) {
	if len(attr) != 2 {
		return nil, fmt.Errorf("Bad circle data: %v", attr)
	}

	radius, err := strconv.ParseFloat(attr[1], 64)
	if err != nil {
		return nil, err
	}
	if radius < 0 {
		return nil, fmt.Errorf("Radius of a circle cannot be negative")
	}

	return &Circle{radius: radius}, nil
}
