package main

import "fmt"

func main() {
	userInput := getUserInput()
	fmt.Print(userInput)
}

func getUserInput() string {
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

func convertValue(number float64, sourceCurrency, targetCurrency string) (convertedCurrency float64) {
	// TODO: реализовать логику конвертации
	return
}
