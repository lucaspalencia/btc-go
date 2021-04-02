package app

func getCurrencySymbol(currency string) string {
	var symbol string = "$"

	switch currency {
	case "eur":
		symbol = "€"
	case "brl":
		symbol = "R$"
	}

	return symbol
}
