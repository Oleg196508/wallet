package main

import (
	"fmt"
	"github.com/Oleg196508/wallet/pkg/wallet"
)

func main() {
	s := &wallet.Service{}
	account, err := s.RegisterAccount("+992000000001")
	if err !=nil {
		fmt.Println(err)
		return
	}
	err = s.Deposit(account.ID, 10)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")	
		}
		return
	}
	fmt.Println(account.Balance) //10
}

