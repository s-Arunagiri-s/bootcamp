package wallet

import (
	"errors"
	"math"
)

// Money stored in INR regardless of Input Currency.
type Wallet struct {
	balanceInr *MoneyObject
}

type MoneyObject struct {
	amount   float64
	currency Currency
}

// /*
type Currency struct {
	name                string
	country             string
	conversionRateToInr float64
}

func NewCurrency(name string, country string, conversionRateToInr float64) (*Currency, error) {
	if conversionRateToInr <= 0 {
		return nil, errors.New("invalid conversion rate")
	}
	return &Currency{
		name:                name,
		country:             country,
		conversionRateToInr: conversionRateToInr,
	}, nil
}

func (c *Currency) GetNameOfCurrency() string {
	return c.name
}

func (c *Currency) GetCountryOfCurrency() string {
	return c.country
}

func (c *Currency) GetConvRateToInrOfCurrency() float64 {
	return c.conversionRateToInr
}

func (c *Currency) SetConvRateToInrOfCurrenct(NewRate float64) {
	c.conversionRateToInr = NewRate
}

//*/

var CurrencyTable = map[string]Currency{}

func init() {
	AddCurrencyAndConvRateToInrInTable("USD", "USA", 82.47)
	AddCurrencyAndConvRateToInrInTable("INR", "INDIA", 1)
	AddCurrencyAndConvRateToInrInTable("SGD", "SINGAPORE", 61.87)
}

func (mb *MoneyObject) roundFloatInr() {
	val := mb.GetAmountOfMoneyObject()
	ratio := math.Pow(10, float64(2))
	mb.amount = math.Round(val*ratio) / ratio
}

func AddCurrencyAndConvRateToInrInTable(name string, country string, convRate float64) error {
	currency, err := NewCurrency(name, country, convRate)
	if err == nil {
		CurrencyTable[name] = *currency
		return nil
	} else {
		return err
	}
}

func NewMoneyObject(amount float64, currency Currency) (*MoneyObject, error) {

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

func (mb *MoneyObject) GetCurrencyOfMoneyObject() Currency {
	return mb.currency
}

// create wallet
func NewWalletBalanceInInr(moneyObject *MoneyObject) *Wallet {
	return &Wallet{
		balanceInr: moneyObject,
	}
}

// Get conversion rate to INR
func GetConRateFromInr(currency Currency) (float64, error) {

	if _, ok := CurrencyTable[currency.name]; ok {
		return float64(CurrencyTable[currency.name].conversionRateToInr), nil
	} else {
		return -1, errors.New("currency not supported")
	}
}

// Converter
func ConverterToInr(moneyObject MoneyObject) (*MoneyObject, error) {
	ConRate, _ := GetConRateFromInr(moneyObject.GetCurrencyOfMoneyObject())

	moneyObjectRet, errMoneyObject := NewMoneyObject(moneyObject.GetAmountOfMoneyObject()*ConRate, CurrencyTable["INR"])
	if errMoneyObject != nil {
		return nil, errMoneyObject
	}

	return moneyObjectRet, nil
}

// return balance
func (w *Wallet) GetBalanceMoneyObjectInInr() *MoneyObject {
	return w.balanceInr
}

// Add money into wallet
func (w *Wallet) AddMoney(moneyObject MoneyObject) error {
	MoneyObjectAddedInInr, _ := ConverterToInr(moneyObject)

	oldBalance := w.GetBalanceMoneyObjectInInr()
	newBalanceInr := oldBalance.GetAmountOfMoneyObject() + MoneyObjectAddedInInr.GetAmountOfMoneyObject()
	newBalanceInrObj, errMoneyObject := NewMoneyObject(newBalanceInr, CurrencyTable["INR"])
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

	oldBalance := w.GetBalanceMoneyObjectInInr()
	newBalanceInr := oldBalance.GetAmountOfMoneyObject() - MoneyObjectAddedInInr.GetAmountOfMoneyObject()
	if newBalanceInr < 0 {
		return errors.New("insufficient funds")
	}

	newBalanceInrObj, errMoneyObject := NewMoneyObject(newBalanceInr, CurrencyTable["INR"])
	if errMoneyObject != nil {
		return errMoneyObject
	}
	newBalanceInrObj.roundFloatInr()
	w.balanceInr = newBalanceInrObj
	return nil
}

// get balance in preferred currency.
func (wallet *Wallet) GetBalanceMoneyObjectInPreferredCurrency(currency Currency) (*MoneyObject, error) {
	ConRate, errConv := GetConRateFromInr(CurrencyTable[currency.name])
	if errConv == nil {
		amountInInr := wallet.GetBalanceMoneyObjectInInr().GetAmountOfMoneyObject()
		amountInPrefCurr := amountInInr / ConRate
		balanceMoneyObjectinPrefCurr, _ := NewMoneyObject(amountInPrefCurr, CurrencyTable["INR"])
		balanceMoneyObjectinPrefCurr.roundFloatInr()
		return balanceMoneyObjectinPrefCurr, nil
	} else {
		return nil, errConv
	}
}

// Comparision of USD and INR
func CheckInrToUsdConRate(UserGuessConRate float64) bool {
	conRate, _ := GetConRateFromInr(CurrencyTable["USD"])
	if UserGuessConRate == conRate {
		return true
	} else {
		return false
	}
}
