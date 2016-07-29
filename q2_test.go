package lendesk

import "testing"

func TestXMLIndent(t *testing.T) {
	indentation := "\t"
	var tests = []struct {
		input    string
		expected string
	}{
		{input: "<1></1>", expected: "<1>\n</1>\n"},
		{input: "<1><2></2></1>", expected: "<1>\n\t<2>\n\t</2>\n</1>\n"},
		{input: "<1><2><3></3></2></1>", expected: "<1>\n\t<2>\n\t\t<3>\n\t\t</3>\n\t</2>\n</1>\n"},
		{input: "<1><2><3><4></4></3></2></1>", expected: "<1>\n\t<2>\n\t\t<3>\n\t\t\t<4>\n\t\t\t</4>\n\t\t</3>\n\t</2>\n</1>\n"},
		{input: "<1><2><3><4><5><6></6></5></4></3></2></1>", expected: "<1>\n\t<2>\n\t\t<3>\n\t\t\t<4>\n\t\t\t\t<5>\n\t\t\t\t\t<6>\n\t\t\t\t\t</6>\n\t\t\t\t</5>\n\t\t\t</4>\n\t\t</3>\n\t</2>\n</1>\n"},
	}

	for _, test := range tests {
		actual, err := XMLIndent(test.input, indentation)
		if err != nil {
			t.Fatal("Unexpected error occurred: ", err)
		}
		if test.expected != actual {
			t.Errorf("Mismatched. Expected %q, but got %q", test.expected, actual)
		}
	}
}

func TestXMLIndentWithInvalidData(t *testing.T) {
	indentation := "\t"
	var tests = []struct {
		input string
	}{
		{input: "1"},
		{input: "<1"},
		{input: "<1></1"},
		{input: "<1</1>"},
		{input: "1></1>"},
	}

	for _, test := range tests {
		_, err := XMLIndent(test.input, indentation)
		if err == nil {
			t.Errorf("Expected error caused by malformed XML didn't occur. Input %q", test.input)
		}
	}
}

func TestIsValidXML(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{input: "<1>", expected: true},
		{input: "</1>", expected: true},
		{input: "1", expected: false},
		{input: "<1", expected: false},
		{input: "1>", expected: false},
	}

	for _, test := range tests {
		actual := isValidXML(test.input)
		if test.expected != actual {
			t.Errorf("Mismatched. Expected %t, but got %t", test.expected, actual)
		}
	}
}

func TestIndent(t *testing.T) {
	indentation := "\t"
	var tests = []struct {
		input    []string
		depth    int
		expected string
	}{
		{input: []string{"<1>"}, depth: 0, expected: "<1>\n"},
		{input: []string{"<1>", "<2>"}, depth: 0, expected: "<1>\n\t<2>\n"},
		{input: []string{"<1>", "<2>", "<3>"}, depth: 0, expected: "<1>\n\t<2>\n\t\t<3>\n"},
	}

	for _, test := range tests {
		actual := indent(test.input, indentation, test.depth)
		if actual != test.expected {
			t.Errorf("Mismatched. Expected %q, but got %q", test.expected, actual)
		}
	}
}

func TestUnindent(t *testing.T) {
	indentation := "\t"
	var tests = []struct {
		input    []string
		depth    int
		expected string
	}{
		{input: []string{"</1>"}, depth: 0, expected: "</1>\n"},
		{input: []string{"</1>", "</2>"}, depth: 1, expected: "\t</1>\n</2>\n"},
	}

	for _, test := range tests {
		actual := unindent(test.input, indentation, test.depth)
		if actual != test.expected {
			t.Errorf("Mismatched. Expected %q, but got %q", test.expected, actual)
		}
	}
}
