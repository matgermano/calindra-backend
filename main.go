package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func init() {
	flag.Parse()
}

var(
	key = flag.String("key", "", "apikey")
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// First address
	fmt.Println("Enter first address:")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	firstAddress := strings.TrimSuffix(text, "\r\n")

	// Second address
	fmt.Println("Enter second address:")
	text, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	secondAddress := strings.TrimSuffix(text, "\r\n")

	geo1, err := GetCoordinatesFromAddress(firstAddress)
	if err != nil {
		panic(err)
	}

	geo2, err := GetCoordinatesFromAddress(secondAddress)
	if err != nil {
		panic(err)
	}

	if len(geo1.Results) == 0 || len(geo2.Results) == 0 {
		panic(errors.New("one of the address (or both of them) was not found"))
	}

	mi, km := EuclidianDistance(geo1.Results[0].Geometry.Location, geo2.Results[0].Geometry.Location)

	fmt.Printf("\n\nFirst address: %s\nFirst Address Geometry: %v / %v\nSecond address: %s\nSecond Address Geometry: %v / %v\n\nDistance in Miles: %v\nDistance in KM: %v",
		firstAddress, geo1.Results[0].Geometry.Location.Latitude, geo1.Results[0].Geometry.Location.Longitude,
		secondAddress, geo2.Results[0].Geometry.Location.Latitude, geo2.Results[0].Geometry.Location.Longitude,
		mi, km)
}
