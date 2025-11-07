package main

import (
	"fmt"
	"math"
)

func main() {
	var height float64
	var weight float64

	fmt.Print("Enter your height (sm): ")
	fmt.Scan(&height)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&weight)

	IMT := calculateIMT(height, weight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Your IMT: %.2f \n", imt)
	fmt.Print(result)
}

func calculateIMT(height float64, weight float64) float64 {
	const IMTPower = 2
	IMT := weight / math.Pow(height/100, IMTPower)
	return IMT
}
