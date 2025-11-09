package main

import (
	"fmt"
	"math"
)

func main() {

	for {
		height, weight := getUserInput()
		IMT := calculateIMT(height, weight)
		outputResult(IMT)

		switch {
		case IMT < 16:
			fmt.Println("You are very skinny")
		case IMT < 18.5:
			fmt.Println("You are skinny")
		case IMT < 25:
			fmt.Println("You are OK")
		case IMT < 30:
			fmt.Println("You are obese")
		default:
			fmt.Println("You are fat")
		}
		decision := tryAgainRequest()
		if decision == "yes" {
			continue
		} else {
			break
		}
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

func tryAgainRequest() (decision string) {
	var response string
	fmt.Print("Do you want to try again? (yes/no)")
	fmt.Scan(&response)
	return response
}
