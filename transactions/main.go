package main

import "fmt"

func main() {

	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalance(transactions)
	fmt.Println(balance)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Enter sum (or n to exit): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	var totalSum float64

	for _, value := range transactions {
		totalSum += value
	}
	return totalSum
}
