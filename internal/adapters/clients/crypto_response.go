package clients

type exchangeResponse struct {
	Buenbit marketData `json:"buenbit"`
	Ripio   marketData `json:"ripio"`
	Binance marketData `json:"binance"`
}

type marketData struct {
	Ask      float64 `json:"ask"`
	TotalAsk float64 `json:"totalAsk"`
}

func (e *exchangeResponse) giveMeTheCheaper() {

}
