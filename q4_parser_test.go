package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	data := []byte("square,5\npentagon,3\ntriangle,2,5,4\nrectangle,4,5\ncircle,2")
	expected := []Geometry{
		&Square{side: 5},
		&Pentagon{side: 3},
		&Triangle{base: 2, height: 5, side: 4},
		&Rectangle{length: 4, width: 5},
		&Circle{radius: 2},
	}

	actual, err := Parse(data)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	if len(actual) != len(expected) {
		t.Errorf("Mismatch elements count. Expected %d, but got %d", len(expected), len(actual))
	}

	// expect ordering to be preserved.
	for index, p := range expected {
		if !reflect.DeepEqual(p, actual[index]) {
			t.Errorf("Mismatch element. Expected %+v, but got %+v", p, actual[index])
		}
	}
}

func TestParseInvalidData(t *testing.T) {
	var tests = []struct {
		data []byte
	}{
		{data: []byte("square")},
		{data: []byte("square,\npentagon,3")},
		{data: []byte("square,4,\ndiamond,3")},
		{data: []byte("square,-9")},
		{data: []byte("circle,abc")},
		{data: []byte("circle,abc,")},
		{data: []byte("rectangle,4")},
		{data: []byte("rectangle,4,")},
		{data: []byte("rectangle,-4,")},
		{data: []byte("triangle,2,3")},
		{data: []byte("triangle,2,3,")},
		{data: []byte("pentagon,")},
		{data: []byte("pentagon,4,")},
		{data: []byte("pentagon,-4")},
	}

	for _, test := range tests {
		_, err := Parse(test.data)
		if err == nil {
			t.Errorf("Expected error didn't occur for data %q ", test.data)
		}
	}
}
