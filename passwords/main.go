package main

import (
	"app/password/account"
	"app/password/files"
	"fmt"
)

func main() {
	fmt.Println("__Password manager__")
Menu:
	for {
		input := getMenu()
		switch input {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var input int
	fmt.Println("1. Create account")
	fmt.Println("2. Find account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	fmt.Scan(&input)
	return input
}

func createAccount() {

	files.ReadFile("data.json")

	login := promptData("enter login")
	password := promptData("enter pass")
	url := promptData("enter url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func findAccount() {

}

func deleteAccount() {

}
