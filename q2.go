package main

import (
	"fmt"
	"strings"
)

// PrintXMLIndent prints the input XML as an indented XML string.
func PrintXMLIndent(input, indentation string) {
	fmt.Println(XMLIndent(input, indentation))
}

// XMLIndent adds indentation to all XML elements in input.
func XMLIndent(input, indentation string) (string, error) {
	if !isValidXML(input) {
		return "", fmt.Errorf("Malformed XML: %q", input)
	}

	// split the input string into two sublists of open and close tags.
	// then indent the sublist of open tags and unindent the sublist of
	// close tags.

	firstCloseTagIndex := strings.Index(input, "</")

	openTags := strings.SplitAfter(input[:firstCloseTagIndex], ">")
	indentedOpenTags := indent(openTags[:len(openTags)-1], indentation, 0)

	closeTags := strings.SplitAfter(input[firstCloseTagIndex:], ">")
	indentedCloseTags := unindent(closeTags[:len(closeTags)-1], indentation, len(openTags)-2)

	return (indentedOpenTags + indentedCloseTags), nil
}

func isValidXML(input string) bool {
	leftBracketCount, rightBracketCount := strings.Count(input, "<"), strings.Count(input, ">")
	return leftBracketCount > 0 && rightBracketCount > 0 && leftBracketCount == rightBracketCount
}

// indent increases the indentation each element in inputs r times.
func indent(inputs []string, indentation string, r int) string {
	if len(inputs) == 0 {
		return ""
	}

	result := strings.Repeat(indentation, r) + inputs[0] + "\n"
	return result + indent(inputs[1:], indentation, r+1)
}

// unindent decreases the indentation of each element in inputs r times.
func unindent(inputs []string, indentation string, r int) string {
	if len(inputs) == 0 {
		return ""
	}

	if r < 0 {
		r = 0
	}

	result := strings.Repeat(indentation, r) + inputs[0] + "\n"
	return result + unindent(inputs[1:], indentation, r-1)
}
