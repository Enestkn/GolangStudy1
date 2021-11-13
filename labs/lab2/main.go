package main

import (
	"fmt"
	"github.com/enestkn/golangstudy1/Currency"
)

func main() {
	var currency Currency.Currency
	currency.CurrencyCode = "EUR"
	currency.Get()
	fmt.Printf("Bugünkü %s döviz kuru %f olmuştur.\n",
		currency.CurrencyCode, currency.ExchangeRate)
}
