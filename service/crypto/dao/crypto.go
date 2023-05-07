package dao

import (
	"base/pkg/db"
	"base/pkg/utils"
)

type Crypto struct {
	Id          int     `json:"rank"`
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Type        string  `json:"type"`
	TotalSupply int     `json:"totalSupply"`
	Image       string  `json:"image"`
	MarketCap   float64 `json:"marketCap"`
	Volume24H   float64 `json:"volume24h"`
	PriceUSD    float64 `json:"priceUSD"`
	CreatedDate string  `json:"createdDate"`
	UpdatedDate string  `json:"updatedDate"`
}

func (crypto *Crypto) Insert() error {
	query := `INSERT INTO crypto (id, key, name, symbol, type, totalSupply, image,
		marketcap, volume24h, priceUSD, createddate, updateddate) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := db.PSQL.Exec(query, crypto.Id, crypto.Key, crypto.Name, crypto.Symbol, crypto.Type, crypto.TotalSupply, crypto.Image,
		crypto.MarketCap, crypto.Volume24H, crypto.PriceUSD, utils.Timestamp(), utils.Timestamp())

	return err
}
 