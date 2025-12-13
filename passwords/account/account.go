package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abwncvoeuvhnsapcFPHEHCIWPCNvo")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.Login)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
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
		Login:    login,
		Url:      urlString,
		Password: password,
	}
	if password == "" {
		newAcc.generatePassword(20)
	}
	return newAcc, nil
}
