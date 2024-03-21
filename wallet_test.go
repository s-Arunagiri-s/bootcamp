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
	moneyObject, _ := NewMoneyObject(100, CurrencyTable["INR"])

	wallet := NewWalletBalanceInInr(moneyObject)
	got := wallet.GetBalanceMoneyObjectInInr()

	assert.NotNil(t, wallet)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())

}

func TestNewWalletWith0Balance(t *testing.T) {
	want := 0.0
	moneyObject, _ := NewMoneyObject(0, CurrencyTable["INR"])

	wallet := NewWalletBalanceInInr(moneyObject)
	got := wallet.GetBalanceMoneyObjectInInr()
	assert.NotNil(t, wallet)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())

}

// New Money Object
func TestNewMoneyObjectWithValidAmountAndCurrency(t *testing.T) {
	amount := 100.0
	currency := "INR"

	moneyobject, err := NewMoneyObject(amount, CurrencyTable[currency])
	curr := moneyobject.GetCurrencyOfMoneyObject()

	assert.NotNil(t, moneyobject)
	assert.Nil(t, err)
	assert.Equal(t, amount, moneyobject.GetAmountOfMoneyObject())
	assert.Equal(t, currency, curr.GetNameOfCurrency())

}

func TestNewMoneyObjectWith0Balance(t *testing.T) {
	amount := 0.0
	currency := "INR"

	moneyobject, err := NewMoneyObject(amount, CurrencyTable[currency])
	curr := moneyobject.GetCurrencyOfMoneyObject()

	assert.NotNil(t, moneyobject)
	assert.Nil(t, err)
	assert.Equal(t, amount, moneyobject.GetAmountOfMoneyObject())
	assert.Equal(t, currency, curr.GetNameOfCurrency())

}

func TestMoneyObjectWithInValidBalance(t *testing.T) {
	amount := -10.0
	currency := "INR"
	moneyobject, err := NewMoneyObject(amount, CurrencyTable[currency])

	assert.NotNil(t, err)
	assert.Nil(t, moneyobject)
	assert.EqualError(t, err, "invalid amount for money object")
}

func TestMoneyObjectWithUnsupportedCurrency(t *testing.T) {
	amount := 10.0
	currency := "EURO"
	moneyobject, err := NewMoneyObject(amount, CurrencyTable[currency])

	assert.NotNil(t, err)
	assert.Nil(t, moneyobject)
	assert.EqualError(t, err, "currency not supported")
}

func TestNewCurrencyYen(t *testing.T) {
	name := "YEN"
	country := "JAPAN"
	convRate := 0.55

	curr, err := NewCurrency(name, country, convRate)

	assert.NotNil(t, curr)
	assert.Nil(t, err)
	assert.Equal(t, name, curr.GetNameOfCurrency())
	assert.Equal(t, country, curr.GetCountryOfCurrency())
	assert.Equal(t, convRate, curr.GetConvRateToInrOfCurrency())

}

// Converter
func TestConUsdtoInr(t *testing.T) {
	inputUsd := 100.0
	currency := "USD"
	moneyobject, _ := NewMoneyObject(inputUsd, CurrencyTable[currency])
	want := 8247.0

	got, err := ConverterToInr(*moneyobject)

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

// GetConRateFromInr
func TestGetConRateForUSD(t *testing.T) {
	want := 82.47

	got, err := GetConRateFromInr(CurrencyTable["USD"])

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestGetConRateForYEN(t *testing.T) {
	want := -1.0

	got, err := GetConRateFromInr(CurrencyTable["YEN"])

	assert.NotNil(t, err)
	assert.EqualError(t, err, "currency not supported")
	assert.Equal(t, want, got)
}

// Money stored in INR regardless of Input Currency.
// AddMoney
func TestMoneyAddedInUsd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(100, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoAdd := 200.0
	currency := "USD"
	moneyObject2, _ := NewMoneyObject(MoneytoAdd, CurrencyTable[currency])
	want := 16594.0

	err := Wallet.AddMoney(*moneyObject2)
	got := Wallet.GetBalanceMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

func TestMoneyAddedInInr(t *testing.T) {
	moneyObject, _ := NewMoneyObject(100, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoAdd := 200.0
	currency := "INR"
	moneyObject2, _ := NewMoneyObject(MoneytoAdd, CurrencyTable[currency])
	want := 300.0

	err := Wallet.AddMoney(*moneyObject2)
	got := Wallet.GetBalanceMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

//Rounding

// debit money
func TestMoneydebitedInUsd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(10000, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoDebit := 20.5
	currency := "USD"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, CurrencyTable[currency])
	want := 8309.37

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetBalanceMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

func TestMoneydebitedInInr(t *testing.T) {
	moneyObject, _ := NewMoneyObject(1000, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoDebit := 200.0
	currency := "INR"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, CurrencyTable[currency])
	want := 800.0

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetBalanceMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

func TestMoneydebitedgreaterThanBalance(t *testing.T) {
	moneyObject, _ := NewMoneyObject(1000, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoDebit := 2000.0
	currency := "INR"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, CurrencyTable[currency])
	want := 1000.0

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetBalanceMoneyObjectInInr()

	assert.NotNil(t, err)
	assert.EqualError(t, err, "insufficient funds")
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

// Add currency
func TestGetConRateForSGD(t *testing.T) {
	want := 61.87
	//AddCurrencyAndConvRateToInr("SGD", 61.87)
	got, err := GetConRateFromInr(CurrencyTable["SGD"])

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestConSgdtoInr(t *testing.T) {
	inputSgd := 100.0
	currency := "SGD"
	//AddCurrencyAndConvRateToInr("SGD", 50)
	moneyobject, _ := NewMoneyObject(inputSgd, CurrencyTable[currency])
	want := 6187.0

	got, err := ConverterToInr(*moneyobject)

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}
func TestMoneyAddedInSgd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(100, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	//AddCurrencyAndConvRateToInr("SGD", 61.87)
	MoneytoAdd := 200.0
	currency := "SGD"
	moneyObject2, _ := NewMoneyObject(MoneytoAdd, CurrencyTable[currency])
	want := 12474.0

	err := Wallet.AddMoney(*moneyObject2)
	got := Wallet.GetBalanceMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

func TestMoneydebitedInSgd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(10000, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	//AddCurrencyAndConvRateToInr("SGD", 61.87)
	MoneytoDebit := 20.5
	currency := "SGD"
	moneyObject2, _ := NewMoneyObject(MoneytoDebit, CurrencyTable[currency])
	want := 8731.67

	err := Wallet.DebitMoney(*moneyObject2)
	got := Wallet.GetBalanceMoneyObjectInInr()

	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

// Balance in preferred currency
func TestBalanceInUsd(t *testing.T) {
	moneyObject, _ := NewMoneyObject(164.94, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	want := 2.0
	currency := "USD"
	got, err := Wallet.GetBalanceMoneyObjectInPreferredCurrency(CurrencyTable[currency])
	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}

func TestBalanceInEuro(t *testing.T) {
	moneyObject, _ := NewMoneyObject(164.94, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	currency := "EUR"
	got, err := Wallet.GetBalanceMoneyObjectInPreferredCurrency(CurrencyTable[currency])
	assert.NotNil(t, err)
	assert.Nil(t, got)
	assert.EqualError(t, err, "currency not supported")

}

func TestBalanceInUsdAfterMultipleTransaction(t *testing.T) {
	moneyObject, _ := NewMoneyObject(0, CurrencyTable["INR"])
	Wallet := NewWalletBalanceInInr(moneyObject)
	MoneytoAdd := 82.47
	currency := "INR"
	moneyObject2, _ := NewMoneyObject(MoneytoAdd, CurrencyTable[currency])
	Wallet.AddMoney(*moneyObject2)
	MoneytoAdd = 1.0
	currency = "USD"
	moneyObject2, _ = NewMoneyObject(MoneytoAdd, CurrencyTable[currency])
	Wallet.AddMoney(*moneyObject2)
	MoneytoAdd = 164.94
	currency = "INR"
	moneyObject2, _ = NewMoneyObject(MoneytoAdd, CurrencyTable[currency])
	Wallet.AddMoney(*moneyObject2)

	want := 4.0
	got, err := Wallet.GetBalanceMoneyObjectInPreferredCurrency(CurrencyTable["USD"])
	assert.Nil(t, err)
	assert.Equal(t, want, got.GetAmountOfMoneyObject())
}
