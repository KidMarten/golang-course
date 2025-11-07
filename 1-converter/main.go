package main

import "fmt"

func main() {
	const USDToEUR = 0.94
	const USDToRub = 90.0

	const EURtoRub = USDToRub / USDToEUR

	fmt.Print(EURtoRub)
}
