package main

import (
	"fmt"
	"testing"
)

func TestParameter(t *testing.T) {
	var tests = []struct {
		fixture  Geometry
		expected float64
	}{
		{fixture: &Circle{radius: 4.0}, expected: 25.13},
		{fixture: &Circle{radius: 10.0}, expected: 62.83},
		{fixture: &Triangle{base: 4.0, side: 4.0}, expected: 12.00},
		{fixture: &Triangle{base: 10.0, side: 12.5}, expected: 35.00},
		{fixture: &Rectangle{length: 1.6, width: 2.9}, expected: 9.00},
		{fixture: &Rectangle{length: 20.9, width: 17.5}, expected: 76.80},
		{fixture: &Square{side: 5.68}, expected: 22.72},
		{fixture: &Square{side: 64.87}, expected: 259.48},
		{fixture: &Pentagon{side: 15.68}, expected: 78.40},
		{fixture: &Pentagon{side: 72.95}, expected: 364.75},
	}

	for _, test := range tests {
		if actual := test.fixture.Perimeter(); !floatEquals(test.expected, actual) {
			t.Errorf("Mismatched. Expected perimeter to be %.3f, but got %.3f", test.expected, actual)
		}
	}
}

func TestArea(t *testing.T) {
	var tests = []struct {
		fixture  Geometry
		expected float64
	}{
		{fixture: &Circle{radius: 4.0}, expected: 50.26},
		{fixture: &Circle{radius: 10.0}, expected: 314.16},
		{fixture: &Triangle{base: 4.0, height: 4.0}, expected: 8.00},
		{fixture: &Triangle{base: 10.0, height: 12.5}, expected: 62.50},
		{fixture: &Rectangle{length: 1.6, width: 2.9}, expected: 4.64},
		{fixture: &Rectangle{length: 20.9, width: 17.5}, expected: 365.75},
		{fixture: &Square{side: 5.68}, expected: 32.26},
		{fixture: &Square{side: 64.87}, expected: 4208.12},
		{fixture: &Pentagon{side: 15.68}, expected: 423},
		{fixture: &Pentagon{side: 72.95}, expected: 9155.87},
	}

	for _, test := range tests {
		if actual := test.fixture.Area(); !floatEquals(test.expected, actual) {
			t.Errorf("Mismatched. Expected area to be %.3f, but got %.3f", test.expected, actual)
		}
	}
}

func TestStringCircle(t *testing.T) {
	fixture := &Circle{radius: 4.0}
	expectedPerimeter := 25.13
	expectedArea := 50.27

	actual := fmt.Sprintf("%s", fixture)
	expected := fmt.Sprintf("circle, with radius of %.2f, having a perimeter of %.2f and an area of %.2f", fixture.radius, expectedPerimeter, expectedArea)
	if expected != actual {
		t.Errorf("Mismatched. Expected %q, but got %q", expected, actual)
	}
}

func TestStringTriangle(t *testing.T) {
	fixture := &Triangle{base: 4.0, height: 4.0, side: 4.0}
	expectedPerimeter := 12.00
	expectedArea := 8.00

	actual := fmt.Sprintf("%s", fixture)
	expected := fmt.Sprintf("triangle, with base of %.2f, height of %.2f and side of %.2f, having a perimeter of %.2f and an area of %.2f", fixture.base, fixture.height, fixture.side, expectedPerimeter, expectedArea)
	if expected != actual {
		t.Errorf("Mismatched. Expected %q, but got %q", expected, actual)
	}
}

func TestStringRectangle(t *testing.T) {
	fixture := &Rectangle{length: 1.6, width: 2.9}
	expectedPerimeter := 9.00
	expectedArea := 4.64

	actual := fmt.Sprintf("%s", fixture)
	expected := fmt.Sprintf("rectangle, with length of %.2f and width of %.2f, having a perimeter of %.2f and an area of %.2f", fixture.length, fixture.width, expectedPerimeter, expectedArea)
	if expected != actual {
		t.Errorf("Mismatched. Expected %q, but got %q", expected, actual)
	}
}

func TestStringSquare(t *testing.T) {
	fixture := &Square{side: 5.68}
	expectedPerimeter := 22.72
	expectedArea := 32.26

	actual := fmt.Sprintf("%s", fixture)
	expected := fmt.Sprintf("square, with side of %.2f, having a perimeter of %.2f and an area of %.2f", fixture.side, expectedPerimeter, expectedArea)
	if expected != actual {
		t.Errorf("Mismatched. Expected %q, but got %q", expected, actual)
	}
}

func TestStringPentagon(t *testing.T) {
	fixture := &Pentagon{side: 72.95}
	expectedPerimeter := 364.75
	expectedArea := 9155.87

	actual := fmt.Sprintf("%s", fixture)
	expected := fmt.Sprintf("pentagon, with side of %.2f, having a perimeter of %.2f and an area of %.2f", fixture.side, expectedPerimeter, expectedArea)
	if expected != actual {
		t.Errorf("Mismatched. Expected %q, but got %q", expected, actual)
	}
}

var EPSILON float64 = 0.01

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}
