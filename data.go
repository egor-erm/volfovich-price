package main

// Root представляет корневую структуру JSON
type Root struct {
	SchemaVersion string `json:"schemaVersion"`
	Pairs         []Pair `json:"pairs"`
}

// Pair содержит информацию о каждой паре
type Pair struct {
	ChainID       string             `json:"chainId"`
	DexID         string             `json:"dexId"`
	URL           string             `json:"url"`
	PairAddress   string             `json:"pairAddress"`
	PriceNative   string             `json:"priceNative"`
	PriceUSD      string             `json:"priceUsd"`
	FDV           int                `json:"fdv"`
	MarketCap     int                `json:"marketCap"`
	PairCreatedAt int                `json:"pairCreatedAt"`
	Labels        []string           `json:"labels"`
	Volume        map[string]float64 `json:"volume"`
	PriceChange   map[string]float64 `json:"priceChange"`
	BaseToken     Token              `json:"baseToken"`
	QuoteToken    Token              `json:"quoteToken"`
	Liquidity     Liquidity          `json:"liquidity"`
	Boosts        Boosts             `json:"boosts"`
	Txns          Txns               `json:"txns"`
	Info          Info               `json:"info"`
}

// Token описывает токен
type Token struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

// Liquidity описывает ликвидность пары
type Liquidity struct {
	USD   float64 `json:"usd"`
	Base  float64 `json:"base"`
	Quote float64 `json:"quote"`
}

// Boosts описывает boost параметры
type Boosts struct {
	Active int `json:"active"`
}

// Txns описывает транзакции
type Txns struct {
	AnyAdditionalProperty struct {
		Buys  int `json:"buys"`
		Sells int `json:"sells"`
	} `json:"ANY_ADDITIONAL_PROPERTY"`
}

// Info содержит дополнительную информацию о паре
type Info struct {
	ImageURL string    `json:"imageUrl"`
	Websites []Website `json:"websites"`
	Socials  []Social  `json:"socials"`
}

// Website описывает веб-сайт
type Website struct {
	URL string `json:"url"`
}

// Social описывает социальные сети
type Social struct {
	Platform string `json:"platform"`
	Handle   string `json:"handle"`
}
