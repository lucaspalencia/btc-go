package config

type Config struct {
	API_URL string

	DATETIME_FORMAT string

	VALID_CURRENCIES []string
}

func New() Config {
	return Config{
		API_URL:          "https://api.coingecko.com/api/v3",
		DATETIME_FORMAT:  "01/02/06 15:04:05",
		VALID_CURRENCIES: []string{"usd", "eur", "brl"},
	}
}
