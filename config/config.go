package config

type Config struct {
	APP_NAME         string
	APP_USAGE        string
	APP_VERSION      string
	API_URL          string
	DATETIME_FORMAT  string
	VALID_CURRENCIES []string
}

func New() Config {
	return Config{
		APP_NAME:         "Bitcoin Go",
		APP_USAGE:        "Get bitcoin market data in USD/EUR/BRL",
		APP_VERSION:      "0.0.1",
		API_URL:          "https://api.coingecko.com/api/v3",
		DATETIME_FORMAT:  "01/02/06 15:04:05",
		VALID_CURRENCIES: []string{"usd", "eur", "brl"},
	}
}
