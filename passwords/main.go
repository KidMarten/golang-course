package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abwncvoeuvhnsapcFPHEHCIWPCNvo")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	updatedAt time.Time
	createdAt time.Time
	account
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login string, password string, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	newAcc := &account{
		login:    login,
		url:      urlString,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}

func newAccountWithTimeStamp(login string, password string, urlString string) (*accountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			url:      urlString,
			password: password},
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}

func main() {

	login := promptData("enter login")
	password := promptData("enter pass")
	url := promptData("enter url")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.outputPassword()

	myAccountWithTS, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccountWithTS.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
