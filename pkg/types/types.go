package types

// Money представляет собой денежную сумму в минимальных единицах
type Money int64

// PaymentCategory представляет собой категорию, в которой был совершён платёж ( авто, аптеки, рестораны и т.д.).
type PaymentCategory string

// PaymentStatus представляет собой статус платежа
type PaymentStatus string

// Предопределённые статусы платежей.
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Payment представляет информацию о платеже
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

// Favorite шаблон для создания платежа
type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}

type Phone string

// Account представляет информацию о счёте пользователя.
type Account struct {
	ID          int64 
	Phone       Phone 
	Balance     Money 
}