package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var height float64
	var weight float64

	fmt.Print("Enter your height (sm): ")
	fmt.Scan(&height)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&weight)

	IMT := weight / math.Pow(height/100, IMTPower)
	result := fmt.Sprintf("Your IMT: %.2f \n", IMT)

	fmt.Print(result)

}
