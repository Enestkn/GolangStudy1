package Currency

import (
	"github.com/icobani/GOTCMBCurrencyHelper"
	"time"
)

type Currency struct {
	CurrencyCode string  `json:"currency_code"`
	ExchangeRate float64 `json:"exchange_rate"`
}

// TODO: Ilgili döviz bugünkü kurunu tcmb adresinden alıp bize verecek
func (r *Currency) Get() {
	today := time.Now()
	journal := GOTCMBCurrencyHelper.GetArchive(today)

	for _, currency := range journal.Currencies {
		if currency.Code == r.CurrencyCode {
			r.ExchangeRate = currency.ForexBuying
			break
		}
	}
}
