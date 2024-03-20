package wallet

import (
	"errors"
	"math"
)

// Money stored in INR regardless of Input Currency.
type Wallet struct {
	balanceInr *MoneyObject
}

/*
	type Currency struct{
		Name string
		conversionRateToInr float64
	}
*/
type MoneyObject struct {
	amount   float64
	currency string
}

var exchangeRates = map[string]float64{}

func init() {
	AddCurrencyAndConvRateToInr("USD", 82.47)
	AddCurrencyAndConvRateToInr("INR", 1)
	AddCurrencyAndConvRateToInr("SGD", 61.87)
}

func (mb *MoneyObject) roundFloatInr() {
	val := mb.amount
	ratio := math.Pow(10, float64(2))
	mb.amount = math.Round(val*ratio) / ratio
}

func AddCurrencyAndConvRateToInr(currency string, convRate float64) {
	exchangeRates[currency] = convRate
}

func NewMoneyObject(amount float64, currency string) (*MoneyObject, error) {

	_, errConv := GetConRateFromInr(currency)
	if amount < 0 {
		return nil, errors.New("invalid amount for money object")
	} else if errConv != nil {
		return nil, errors.New("currency not supported")
	}
	return &MoneyObject{
		amount:   amount,
		currency: currency,
	}, nil
}

func (mb *MoneyObject) GetAmountOfMoneyObject() float64 {
	return mb.amount
}

func (mb *MoneyObject) GetCurrencyOfMoneyObject() string {
	return mb.currency
}

// create wallet
func NewWalletBalanceInInr(moneyObject *MoneyObject) *Wallet {
	return &Wallet{
		balanceInr: moneyObject,
	}
}

// Get conversion rate to INR
func GetConRateFromInr(currency string) (float64, error) {

	if _, ok := exchangeRates[currency]; ok {
		return float64(exchangeRates[currency]), nil
	} else {
		return -1, errors.New("currency not supported")
	}
}

// Converter
func ConverterToInr(moneyObject MoneyObject) (*MoneyObject, error) {
	ConRate, _ := GetConRateFromInr(moneyObject.currency)

	moneyObjectRet, errMoneyObject := NewMoneyObject(moneyObject.amount*ConRate, "INR")
	if errMoneyObject != nil {
		return nil, errMoneyObject
	}

	return moneyObjectRet, nil
}

// return balance
func (w *Wallet) GetMoneyObjectInInr() *MoneyObject {
	return w.balanceInr
}

// Add money into wallet
func (w *Wallet) AddMoney(moneyObject MoneyObject) error {
	MoneyObjectAddedInInr, _ := ConverterToInr(moneyObject)

	oldBalance := w.GetMoneyObjectInInr()
	newBalanceInr := oldBalance.amount + MoneyObjectAddedInInr.amount
	newBalanceInrObj, errMoneyObject := NewMoneyObject(newBalanceInr, "INR")
	if errMoneyObject != nil {
		return errMoneyObject
	}
	newBalanceInrObj.roundFloatInr()
	w.balanceInr = newBalanceInrObj
	return nil

}

// Remove money from the wallet
func (w *Wallet) DebitMoney(moneyObject MoneyObject) error {
	MoneyObjectAddedInInr, _ := ConverterToInr(moneyObject)

	oldBalance := w.GetMoneyObjectInInr()
	newBalanceInr := oldBalance.amount - MoneyObjectAddedInInr.amount
	if newBalanceInr < 0 {
		return errors.New("insufficient funds")
	}

	newBalanceInrObj, errMoneyObject := NewMoneyObject(newBalanceInr, "INR")
	if errMoneyObject != nil {
		return errMoneyObject
	}
	newBalanceInrObj.roundFloatInr()
	w.balanceInr = newBalanceInrObj
	return nil
}

// Comparision of USD and INR
func CheckInrToUsdConRate(UserGuessConRate float64) bool {
	conRate, _ := GetConRateFromInr("USD")
	if UserGuessConRate == conRate {
		return true
	} else {
		return false
	}
}
