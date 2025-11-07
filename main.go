package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func roundToTwoDecimalPlaces(f float64) float64 {
	return math.Round(f*100) / 100
}

func main() {
	radiusStr := os.Args[1]
	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		os.Exit(1)
	}

	if radius < 0.1 || radius > 100.0 {
		fmt.Printf("Error. Radius must be between 0.1 and 100")
		os.Exit(1)
	}

	var area = math.Pi * math.Pow(radius, 2)
	var formattedArea = roundToTwoDecimalPlaces(area)
	fmt.Print(formattedArea)
}
