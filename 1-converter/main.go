package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("__ Currency converter __")
	sourceCurrency, value, targetCurrency := getUserInput()
	convertedValue := convertValue(sourceCurrency, value, targetCurrency)
	fmt.Printf("Exchange rate %s/%s is %.2f\n", sourceCurrency, targetCurrency, convertedValue)
}

func getUserInput() (string, float64, string) {
	sourceCurrency := getSourceCurrency()
	value := getValue()
	targetCurrency := getTargetCurrency(sourceCurrency)
	return sourceCurrency, value, targetCurrency
}

func getSourceCurrency() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter source currency (USD, Rub, EUR): ")
		if !scanner.Scan() {
			os.Exit(1)
		}
		currency := strings.TrimSpace(scanner.Text())
		if checkAvailableSourceCurrency(currency) {
			return currency
		}
		fmt.Println("Available values: USD, Rub, EUR")
	}
}

func checkAvailableSourceCurrency(sourceCurrency string) bool {
	switch sourceCurrency {
	case "EUR", "USD", "Rub":
		return true
	}
	return false
}

func getValue() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter value: ")
		if !scanner.Scan() {
			fmt.Println("\nInput ended. Exiting.")
			os.Exit(1)
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			fmt.Println("Input cannot be empty. Try again.")
			continue
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error: input is not a valid number. Try again.")
			continue
		}

		if value <= 0 {
			fmt.Println("Value must be greater than 0. Try again.")
			continue
		}

		return value
	}
}

func getTargetCurrency(sourceCurrency string) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter target currency (USD, Rub, EUR): ")
		if !scanner.Scan() {
			os.Exit(1)
		}
		currency := strings.TrimSpace(scanner.Text())
		if currency == sourceCurrency {
			fmt.Println("Currencies are the same. Choose a different currency.")
			continue
		}
		if checkAvailableSourceCurrency(currency) {
			return currency
		}
		fmt.Println("Available values: USD, Rub, EUR")
	}
}

func convertValue(sourceCurrency string, value float64, targetCurrency string) float64 {
	const USDToEUR = 0.94
	const USDToRub = 90.0

	toRates := map[string]float64{
		"USD": 1.0,
		"EUR": USDToEUR,
		"Rub": USDToRub,
	}

	fromRates := map[string]float64{
		"USD": 1.0,
		"EUR": 1.0 / USDToEUR,
		"Rub": 1.0 / USDToRub,
	}

	fromRate := fromRates[sourceCurrency]
	toRate := toRates[targetCurrency]

	return fromRate * value * toRate
}
