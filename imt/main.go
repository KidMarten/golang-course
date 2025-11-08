package main

import (
	"fmt"
	"math"
)

func main() {
	height, weight := getUserInput()
	IMT := calculateIMT(height, weight)
	outputResult(IMT)
	if IMT < 16 {
		fmt.Println("You are very skinny")
	} else if IMT >= 16 && IMT < 18.5 {
		fmt.Println("You are skinny")
	} else if IMT >= 18.5 && IMT < 25 {
		fmt.Println("You are OK")
	} else if IMT >= 25 && IMT < 30 {
		fmt.Println("You are obese")
	} else {
		fmt.Println("You are fat")
	}
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

func getUserInput() (float64, float64) {
	var height float64
	var weight float64

	fmt.Print("Enter your height (sm): ")
	fmt.Scan(&height)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&weight)

	return height, weight
}
