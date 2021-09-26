package wallet

import (
	"reflect"
	"testing"

	"github.com/Oleg196508/wallet/pkg/types"
)


func TestService_Repeat_success(t *testing.T) {
	//создаём сервис
	s := &Service{}

	//регистрируем там пользователя
	phone := types.Phone("+992988888881")
	account, err := s.RegisterAccount(phone)
	if err !=nil {
		t.Errorf("Repeat(): can't register account, error = %v", err)
		return
	}

	//пополняем его счёт
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("Repeat(): can't deposit account, error = %v", err)
		return
	}

	//осущствляем платёж на его счёт
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("Repeat(): can't create payment, error = %v", err)
		return	
	}

	//пробуем повторить платёж
	payment, err = s.Repeat(payment.ID)
	if err != nil {
		t.Errorf("Repeat(): error = %v", err)
		return
	}		
	//пробуем отменить платёж
	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Repeat(): error = %v", err)
		return
	}		
}

func TestService_FindPaymentByID_success(t *testing.T) {
	//создаём сервис
	s := &Service{}

	//регистрируем там пользователя
	phone := types.Phone("+992988888881")
	account, err := s.RegisterAccount(phone)
	if err !=nil {
		t.Errorf("Repeat(): can't register account, error = %v", err)
		return
	}

	//пополняем его счёт
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("Repeat(): can't deposit account, error = %v", err)
		return
	}

	//осущствляем платёж на его счёт
	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("Repeat(): can't create payment, error = %v", err)
		return	
	}

	//пробуем найти платёж
	got, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	}
	//сравниваем платежи
	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID():wrong payment returned = %v", err)
		return
	}
}

func TestService_FindAccountByID_notFound(t *testing.T) {
	s := Service{
		accounts: []*types.Account{
			{ ID: 10, Phone: "9929888881", Balance: 1000,},
			{ ID: 15, Phone: "9929999991", Balance: 1500,},		
		},
	}

	expected := ErrAccountNotFound

	result, _ := s.FindAccountByID(5)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid Result: excpected: %v, actual: %v ", expected, result)
	}
}

func TestService_Reject_paymentNotFound(t *testing.T) {
	s := Service{
		payments: []*types.Payment{
			{ ID: "1020", AccountID: 99, Amount: 1000,},
			{ ID: "1515", AccountID: 100, Amount: 1500,},		
		},
		accounts: []*types.Account{
			{ID: 100,},
		},
	}

	expected := ErrPaymentNotFound

	result := s.Reject("1600")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid Result: excpected: %v, actual: %v ", expected, result)
	}
}

