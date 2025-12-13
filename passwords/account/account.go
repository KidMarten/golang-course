package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abwncvoeuvhnsapcFPHEHCIWPCNvo")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	updatedAt time.Time
	createdAt time.Time
	Account
}

func (acc Account) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccount(login string, password string, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	newAcc := &Account{
		login:    login,
		url:      urlString,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}

func NewAccountWithTimeStamp(login string, password string, urlString string) (*AccountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    login,
			url:      urlString,
			password: password},
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}
