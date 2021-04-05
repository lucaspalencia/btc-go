package services

type UsdPriceResponse struct {
	Bitcoin struct {
		Price            int     `json:"usd"`
		MarketPercentage float64 `json:"usd_24h_change"`
		UpdatedAt        int     `json:"last_updated_at"`
	} `json:"bitcoin"`
}

type EurPriceResponse struct {
	Bitcoin struct {
		Price            int     `json:"eur"`
		MarketPercentage float64 `json:"eur_24h_change"`
		UpdatedAt        int     `json:"last_updated_at"`
	} `json:"bitcoin"`
}

type BrlPriceResponse struct {
	Bitcoin struct {
		Price            int     `json:"brl"`
		MarketPercentage float64 `json:"brl_24h_change"`
		UpdatedAt        int     `json:"last_updated_at"`
	} `json:"bitcoin"`
}
