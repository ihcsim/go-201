package main

import "testing"

func TestPrintXMLIndent(t *testing.T) {
	var tests = []struct {
		input    []int
		expected string
	}{
		{input: []int{1}, expected: "<1>\n</1>\n"},
		{input: []int{1, 2}, expected: "<1>\n\t<2>\n\t</2>\n</1>\n"},
		{input: []int{1, 2}, expected: "<1>\n\t<2>\n\t</2>\n</1>\n"},
		{input: []int{1, 2, 3}, expected: "<1>\n\t<2>\n\t\t<3>\n\t\t</3>\n\t</2>\n</1>\n"},
		{input: []int{1, 2, 3, 4}, expected: "<1>\n\t<2>\n\t\t<3>\n\t\t\t<4>\n\t\t\t</4>\n\t\t</3>\n\t</2>\n</1>\n"},
		{input: []int{1, 2, 3, 4, 5, 6}, expected: "<1>\n\t<2>\n\t\t<3>\n\t\t\t<4>\n\t\t\t\t<5>\n\t\t\t\t\t<6>\n\t\t\t\t\t</6>\n\t\t\t\t</5>\n\t\t\t</4>\n\t\t</3>\n\t</2>\n</1>\n"},
	}

	indent := "\t"
	for _, test := range tests {
		actual := toXMLIndent(test.input, indent)
		if test.expected != actual {
			t.Errorf("Mismatch: Expected %q, but got %q", test.expected, actual)
		}
	}
}

func TestOpenTag(t *testing.T) {
	var tests = []struct {
		input      int
		indentSize int
		expected   string
	}{
		{input: 1, indentSize: 0, expected: "<1>"},
		{input: 1, indentSize: 1, expected: " <1>"},
		{input: 1, indentSize: 5, expected: "     <1>"},
		{input: 2, indentSize: 0, expected: "<2>"},
		{input: 2, indentSize: 1, expected: " <2>"},
		{input: 2, indentSize: 5, expected: "     <2>"},
	}

	for _, test := range tests {
		actual := openTag(test.input, test.indentSize, " ")
		if test.expected != actual {
			t.Errorf("Bad input %d, indent size %d: Expected %q, but got %q", test.input, test.indentSize, test.expected, actual)
		}
	}
}

func TestCloseTag(t *testing.T) {
	var tests = []struct {
		input      int
		indentSize int
		expected   string
	}{
		{input: 1, indentSize: 0, expected: "</1>"},
		{input: 1, indentSize: 1, expected: " </1>"},
		{input: 1, indentSize: 5, expected: "     </1>"},
		{input: 2, indentSize: 0, expected: "</2>"},
		{input: 2, indentSize: 1, expected: " </2>"},
		{input: 2, indentSize: 5, expected: "     </2>"},
	}

	for _, test := range tests {
		actual := closeTag(test.input, test.indentSize, " ")
		if test.expected != actual {
			t.Errorf("Bad input %d, indent size %d: Expected %q, but got %q", test.input, test.indentSize, test.expected, actual)
		}
	}
}
