package clients

type ExchangeResponse struct {
	ArgenBTC      MarketData `json:"argenbtc"`
	Buenbit       MarketData `json:"buenbit"`
	Ripio         MarketData `json:"ripio"`
	RipioExchange MarketData `json:"ripioexchange"`
	Satoshitango  MarketData `json:"satoshitango"`
	CryptoMKT     MarketData `json:"cryptomkt"`
	Decrypto      MarketData `json:"decrypto"`
	Latamex       MarketData `json:"latamex"`
	Letsbit       MarketData `json:"letsbit"`
	BinanceP2P    MarketData `json:"binancep2p"`
	Fiwind        MarketData `json:"fiwind"`
	LemonCash     MarketData `json:"lemoncash"`
	OKEX_P2P      MarketData `json:"okexp2p"`
	Bitmonedero   MarketData `json:"bitmonedero"`
	PaxfulP2P     MarketData `json:"paxfulp2p"`
	Belo          MarketData `json:"belo"`
	HuobiP2P      MarketData `json:"huobip2p"`
	TiendaCrypto  MarketData `json:"tiendacrypto"`
	BybitP2P      MarketData `json:"bybitp2p"`
	KucoinP2P     MarketData `json:"kucoinp2p"`
	Saldo         MarketData `json:"saldo"`
	KriptonMarket MarketData `json:"kriptonmarket"`
	TrubitP2P     MarketData `json:"trubitp2p"`
	BitgetP2P     MarketData `json:"bitgetp2p"`
	PlusCrypto    MarketData `json:"pluscrypto"`
	Bybit         MarketData `json:"bybit"`
	Binance       MarketData `json:"binance"`
	PaydeceP2P    MarketData `json:"paydecep2p"`
	EldoradoP2P   MarketData `json:"eldoradop2p"`
	Trubit        MarketData `json:"trubit"`
	BingxP2P      MarketData `json:"bingxp2p"`
	BitsoAlpha    MarketData `json:"bitsoalpha"`
	LemonCashP2P  MarketData `json:"lemoncashp2p"`
}

type MarketData struct {
	TotalAsk float64 `json:"totalAsk"`
	TotalBid float64 `json:"totalBid"`
}

func (ex *ExchangeResponse) PropertiesGlossary() string {
	return "Total Ask significa el valor de compra de la moneda. Total bid significa el valor de venta"
}
