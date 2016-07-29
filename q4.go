package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Printf("Error: Missing CSV file.\nUsage:\n\tgo-201 <path-to-csv>")
		os.Exit(1)
	}

	dataFile := flag.Arg(0)
	b, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	geometries, err := Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	report(geometries)
}

func report(gemoetries []Geometry) {
	for index, g := range gemoetries {
		fmt.Printf("Shape %d is a %s\n", index, g)
	}
}
