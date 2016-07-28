package lendesk

import (
	"fmt"
	"strconv"
	"strings"
)

// PrintAsXML prints the inputs slice as a string of indented XML tags.
func PrintAsXML(inputs []int, indentation string) {
	fmt.Println(toXMLIndent(inputs, indentation))
}

// toXMLIndent converts the inputs slice into a string of indented XML tags.
func toXMLIndent(inputs []int, indentation string) string {
	return toXML(inputs, 0, indentation)
}

// toXML recursively convert each element in inputs into a pair of open and close tags.
func toXML(inputs []int, indentSize int, indentation string) string {
	if len(inputs) == 0 {
		return ""
	}

	result := openTag(inputs[0], indentSize, indentation) + "\n"
	result = result + toXML(inputs[1:], indentSize+1, indentation)
	result = result + closeTag(inputs[0], indentSize, indentation) + "\n"

	return result
}

// openTag creates an indented XML open tag with input as the tag. The indentation is repeated r times.
func openTag(input, r int, indentation string) string {
	return strings.Repeat(indentation, r) + "<" + strconv.Itoa(input) + ">"
}

// closeTag creates an indented XML close tag with input as the tag. The indentation is repeated r times.
func closeTag(input, r int, indentation string) string {
	return strings.Repeat(indentation, r) + "</" + strconv.Itoa(input) + ">"
}
