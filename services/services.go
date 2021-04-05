package services

import (
	"btc-go/config"
	"btc-go/httpclient"
	"encoding/json"
)

type BitcoinData struct {
	Price            int
	MarketPercentage float64
	UpdatedAt        int
}

func GetBitcoinData(currency string) (BitcoinData, error) {
	urlParams := "?ids=bitcoin&vs_currencies=" + currency + "&include_24hr_change=true&include_last_updated_at=true"
	url := "" + config.New().API_URL + "/simple/price" + urlParams + ""

	httpClient := httpclient.New(url)

	resp, err := httpClient.Get()

	if err != nil {
		return BitcoinData{}, err
	}

	bitcoinData := mapBitcoinData(currency, resp)

	return bitcoinData, nil
}

func mapBitcoinData(currency string, resp []byte) BitcoinData {
	switch currency {
	case "usd":
		usdResponse := getUsdResponse(resp)

		return BitcoinData{
			Price:            usdResponse.Bitcoin.Price,
			MarketPercentage: usdResponse.Bitcoin.MarketPercentage,
			UpdatedAt:        usdResponse.Bitcoin.UpdatedAt,
		}
	case "eur":
		eurResponse := getEurResponse(resp)

		return BitcoinData{
			Price:            eurResponse.Bitcoin.Price,
			MarketPercentage: eurResponse.Bitcoin.MarketPercentage,
			UpdatedAt:        eurResponse.Bitcoin.UpdatedAt,
		}
	case "brl":
		brlResponse := getBrlResponse(resp)

		return BitcoinData{
			Price:            brlResponse.Bitcoin.Price,
			MarketPercentage: brlResponse.Bitcoin.MarketPercentage,
			UpdatedAt:        brlResponse.Bitcoin.UpdatedAt,
		}
	}

	return BitcoinData{}
}

func getUsdResponse(resp []byte) UsdPriceResponse {
	usdPriceResponse := UsdPriceResponse{}
	json.Unmarshal(resp, &usdPriceResponse)

	return usdPriceResponse
}

func getEurResponse(resp []byte) EurPriceResponse {
	eurPriceResponse := EurPriceResponse{}
	json.Unmarshal(resp, &eurPriceResponse)

	return eurPriceResponse
}

func getBrlResponse(resp []byte) BrlPriceResponse {
	brlPriceResponse := BrlPriceResponse{}
	json.Unmarshal(resp, &brlPriceResponse)

	return brlPriceResponse
}
