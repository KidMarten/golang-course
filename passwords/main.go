package main

import (
	"app/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Password manager__")
	vault := account.NewVault()
Menu:
	for {
		input := getMenu()
		switch input {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
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

func createAccount(vault *account.Vault) {

	login := promptData("enter login")
	password := promptData("enter pass")
	url := promptData("enter url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func findAccount(vault *account.Vault) {
	// URL
	url := promptData("Enter url")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("No accounts found")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccount() {

}
