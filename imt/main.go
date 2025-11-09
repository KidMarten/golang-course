package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	for {
		height, weight := getUserInput()
		IMT, err := calculateIMT(height, weight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(IMT)
		decision := tryAgainRequest()
		if decision == "yes" {
			continue
		} else {
			break
		}
	}
}

func outputResult(imt float64) {
	switch {
	case imt < 16:
		fmt.Println("You are very skinny")
	case imt < 18.5:
		fmt.Println("You are skinny")
	case imt < 25:
		fmt.Println("You are OK")
	case imt < 30:
		fmt.Println("You are obese")
	default:
		fmt.Println("You are fat")
	}
	result := fmt.Sprintf("Your IMT: %.2f \n", imt)
	fmt.Print(result)
}

func calculateIMT(height float64, weight float64) (float64, error) {
	const IMTPower = 2

	if weight <= 0 || height <= 0 {
		return 0, errors.New("height or weight not entered")
	}

	IMT := weight / math.Pow(height/100, IMTPower)
	return IMT, nil
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
