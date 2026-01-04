package main

import (
	"app/password/account"
	"app/password/files"
	"app/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Password manager__")
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		input := promptData([]string{
			"1. Create account",
			"2. Find account",
			"3. Delete account",
			"4. Exit",
			"Choose option",
		})
		switch input {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAccount(vault *account.VaultWithDb) {

	login := promptData([]string{"enter login"})
	password := promptData([]string{"enter pass"})
	url := promptData([]string{"enter url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError(err)
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData[T any](prompt []T) string {
	for i, element := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", element)
		} else {
			fmt.Println(element)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Enter url"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("No accounts found")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Enter url"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Deleted account")
	} else {
		output.PrintError("Unable to find an account")
	}
}
