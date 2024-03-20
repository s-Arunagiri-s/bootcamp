package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// req 1
func TestIf1DollarIs82Rs(t *testing.T) {
	want := false

	got := CheckInrToUsdConRate(82)

	assert.Equal(t, want, got)
}

func TestIf1DollarIs82point47Rs(t *testing.T) {
	want := true

	got := CheckInrToUsdConRate(82.47)

	assert.Equal(t, want, got)
}

// Money stored in INR regardless of Input Currency.
// new wallet
func TestNewWalletWithValidBalance(t *testing.T) {
	want := 100.0
	moneyObject, _ := NewMoneyObject(100, "INR")

	wallet := NewWalletBalanceInInr(moneyObject)
	got := wallet.GetMoneyObjectInInr()

	assert.NotNil(t, wallet)
	assert.Equal(t, want, got.amount)

}

func TestNewWalletWith0Balance(t *testing.T) {
	want := 0.0
	moneyObject, _ := NewMoneyObject(0, "INR")

	wallet := NewWalletBalanceInInr(moneyObject)
	got := wallet.GetMoneyObjectInInr()
	assert.NotNil(t, wallet)
	assert.Equal(t, want, got.amount)

}

// New Money Object
func TestNewMoneyObjectWithValidAmountAndCurrency(t *testing.T) {
	amount := 100.0
	currency := "INR"

	moneyobject, err := NewMoneyObject(amount, currency)

	assert.NotNil(t, moneyobject)
	assert.Nil(t, err)
	assert.Equal(t, amount, moneyobject.GetAmountOfMoneyObject())
	assert.Equal(t, currency, moneyobject.GetCurrencyOfMoneyObject())

}

func TestNewMoneyObjectWith0Balance(t *testing.T) {
	amount := 0.0
	currency := "INR"

	moneyobject, err := NewMoneyObject(amount, currency)

	assert.NotNil(t, moneyobject)
	assert.Nil(t, err)
	assert.Equal(t, amount, moneyobject.GetAmountOfMoneyObject())
	assert.Equal(t, currency, moneyobject.GetCurrencyOfMoneyObject())

}

func TestMoneyObjectWithInValidBalance(t *testing.T) {
	amount := -10.0
	currency := "INR"
	moneyobject, err := NewMoneyObject(amount, currency)

	assert.NotNil(t, err)
	assert.Nil(t, moneyobject)
	assert.EqualError(t, err, "invalid amount for money object")
}

func TestMoneyObjectWithUnsupportedCurrency(t *testing.T) {
	amount := 10.0
	currency := "EURO"
	moneyobject, err := NewMoneyObject(amount, currency)

	assert.NotNil(t, err)
	assert.Nil(t, moneyobject)
	assert.EqualError(t, err, "currency not supported")
}

// Converter
func TestConUsdtoInr(t *testing.T) {
	inputUsd := 100.0
	currency := "USD"
	moneyobject, _ := NewMoneyObject(inputUsd, currency)
	want := 8247.0

	got, err := ConverterToInr(*moneyobject)

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}

// GetConRateFromInr
func TestGetConRateForUSD(t *testing.T) {
	want := 82.47

	got, err := GetConRateFromInr("USD")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestGetConRateForYEN(t *testing.T) {
	want := -1.0

	got, err := GetConRateFromInr("YEN")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "currency not supported")
	assert.Equal(t, want, got)
}

// Money stored in INR regardless of Input Currency.
// AddMoney
func TestMoneyAddedInUsd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(100, "INR")
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoAdd := 200.0
	Currency := "USD"
	moneyObject2, _ := NewMoneyObject(MoneytoAdd, Currency)
	want := 16594.0

	err := Wallet.AddMoney(*moneyObject2)
	got := Wallet.GetMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}

func TestMoneyAddedInInr(t *testing.T) {
	moneyObject, _ := NewMoneyObject(100, "INR")
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoAdd := 200.0
	Currency := "INR"
	moneyObject2, _ := NewMoneyObject(MoneytoAdd, Currency)
	want := 300.0

	err := Wallet.AddMoney(*moneyObject2)
	got := Wallet.GetMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}

//Rounding

// debit money
func TestMoneydebitedInUsd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(10000, "INR")
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoDebit := 20.5
	Currency := "USD"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, Currency)
	want := 8309.37

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}

func TestMoneydebitedInInr(t *testing.T) {
	moneyObject, _ := NewMoneyObject(1000, "INR")
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoDebit := 200.0
	Currency := "INR"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, Currency)
	want := 800.0

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}

func TestMoneydebitedgreaterThanBalance(t *testing.T) {
	moneyObject, _ := NewMoneyObject(1000, "INR")
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoDebit := 2000.0
	Currency := "INR"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, Currency)
	want := 1000.0

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetMoneyObjectInInr()

	assert.NotNil(t, err)
	assert.EqualError(t, err, "insufficient funds")
	assert.Equal(t, want, got.amount)
}

// Add currency
func TestGetConRateForSGD(t *testing.T) {
	want := 61.87
	//AddCurrencyAndConvRateToInr("SGD", 61.87)
	got, err := GetConRateFromInr("SGD")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestConSgdtoInr(t *testing.T) {
	inputSgd := 100.0
	currency := "SGD"
	//AddCurrencyAndConvRateToInr("SGD", 50)
	moneyobject, _ := NewMoneyObject(inputSgd, currency)
	want := 6187.0

	got, err := ConverterToInr(*moneyobject)

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}
func TestMoneyAddedInSgd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(100, "INR")
	Wallet := NewWalletBalanceInInr(moneyObject)
	//AddCurrencyAndConvRateToInr("SGD", 61.87)
	MoneytoAdd := 200.0
	Currency := "SGD"
	moneyObject2, _ := NewMoneyObject(MoneytoAdd, Currency)
	want := 12474.0

	err := Wallet.AddMoney(*moneyObject2)
	got := Wallet.GetMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}

func TestMoneydebitedInSgd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(10000, "INR")
	Wallet := NewWalletBalanceInInr(moneyObject)
	//AddCurrencyAndConvRateToInr("SGD", 61.87)
	MoneytoDebit := 20.5
	Currency := "SGD"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, Currency)
	want := 8731.67

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.amount)
}
