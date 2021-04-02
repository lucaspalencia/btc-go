package app

import "btc-go/config"

func isValidCurrency(currency string) bool {
	for _, validCurrency := range config.New().VALID_CURRENCIES {
		if currency == validCurrency {
			return true
		}
	}

	return false
}
