package wallet

import (
	"testing"
	"reflect"
	"github.com/Oleg196508/wallet/pkg/types"
)

func TestService_FindAccountByID_registered(t *testing.T) {
	s := Service{
		accounts: []*types.Account{
			{ ID: 10, Phone: "9929888881", Balance: 1000,},
			{ ID: 15, Phone: "9929999991", Balance: 1500,},		
		},
	}

	expected := &types.Account{
		ID: 10, Phone: "9929888881", Balance: 1000,	
	}

	result, _ := s.FindAccountByID(10)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Invalid Result: excpected: %v, actual: %v ", expected, result)
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
func TestService_Reject_paymentFound(t *testing.T) {
	s := Service{
		payments: []*types.Payment{
			{ ID: "1020", AccountID: 99, Amount: 1000,},
			{ ID: "1515", AccountID: 100, Amount: 1500,},		
		},
		accounts: []*types.Account{
			{ID: 100,},
		},
	}


	result := s.Reject("1515")

	if !reflect.DeepEqual(nil, result) {
		t.Errorf("Invalid Result: excpected: %v, actual: %v ", nil, result)
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


