package wallet

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Oleg196508/wallet/pkg/types"
	"github.com/google/uuid"
)

type testService struct{
	*Service
}
func newTestService() *testService {
	return &testService{Service: &Service{}}	
}

type testAccount struct {
	phone      types.Phone
	balance    types.Money
	payments   []struct {
		amount    types.Money
		category  types.PaymentCategory
	}
}
//создаём значения по умолчанию
var defaultTestAccount = testAccount{
	phone:         "+992988888881",
	balance:       10_000_00,
	payments:      []struct {
		amount    types.Money
		category  types.PaymentCategory
	}{
		{amount: 5_000_00, category: "auto" },
	},
}
//создаём метод, который генерирует тестовый аккаунт и платёж
func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error) {
    //регистрируем там пользователя
	account, err := s.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}
	//пополняем его счёт
	err = s.Deposit(account.ID, data.balance)
	if err !=nil {
		return nil, nil, fmt.Errorf("can't deposit account, error = %v", err)
	}
	//выполняем платежи
	//можем создать слайс сразу нужной длины, поскольку знаем размер
	payments := make([]*types.Payment, len(data.payments))
	for i, payment := range data.payments{
		//тогда здесь работаем через index, а не через append
		payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
		if err !=nil {
			return nil, nil, fmt.Errorf("can't make payment, error = %v", err)
		}
	}
	return account, payments, nil	
}

func TestService_FindPaymentByID_success(t *testing.T) {
	//создаём сервис
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}	
	//пробуем найти платёж
	payment := payments[0]
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

func TestService_FindPaymentByID_notFound(t *testing.T) {
	//создаём сервис
	s := newTestService()
	_, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}	
	//пробуем найти несуществующий платёж
	_, err = s.FindPaymentByID(uuid.New().String())
	if err == nil {
		t.Errorf("FindPaymentByID(): must return error, returned nil")
		return
	}
	if err != ErrPaymentNotFound {
		t.Errorf("FindPaymentByID(): must return ErrPaymentNotFound, returned =%v", err)
		return
	}
}

func TestService_Repeat_success(t *testing.T) {
	//создаём сервис
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	//пробуем повторить платёж
	payment := payments[0]
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
// создаёт избранное из конкретного платежа совершает платёж
func TestService_FavoritePayment_success(t *testing.T) {
	//создаём сервис
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	payment := payments[0]
	favorite, err := s.FavoritePayment(payment.ID, "autom")
	if err != nil{
		t.Error(err)
		return
	}
	payment, err = s.PayFromFavorite(favorite.ID)
	if err != nil{
		t.Errorf("FavoritePayment(): error = %v", err)
		return
	}
}


