package services

type UsdPriceResponse struct {
	Bitcoin struct {
		Price            int     `json:"usd"`
		MarketCap        float64 `json:"usd_market_cap"`
		MarketVolume     float64 `json:"usd_24h_vol"`
		MarketPercentage float64 `json:"usd_24h_change"`
		UpdatedAt        int     `json:"last_updated_at"`
	} `json:"bitcoin"`
}

type EurPriceResponse struct {
	Bitcoin struct {
		Price            int     `json:"eur"`
		MarketCap        float64 `json:"eur_market_cap"`
		MarketVolume     float64 `json:"eur_24h_vol"`
		MarketPercentage float64 `json:"eur_24h_change"`
		UpdatedAt        int     `json:"last_updated_at"`
	} `json:"bitcoin"`
}

type BrlPriceResponse struct {
	Bitcoin struct {
		Price            int     `json:"brl"`
		MarketCap        float64 `json:"brl_market_cap"`
		MarketVolume     float64 `json:"brl_24h_vol"`
		MarketPercentage float64 `json:"brl_24h_change"`
		UpdatedAt        int     `json:"last_updated_at"`
	} `json:"bitcoin"`
}
