package dao

import (
	"base/pkg/db"
	"base/pkg/utils"
)

type CryptoRepo struct {
	Cryptos []Crypto `json:"cryptos"`
}

type Crypto struct {
	Id          int     `json:"id"`
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Chainname   string  `json:"chainname"`
	Address     string  `json:"address"`
	Type        string  `json:"type"`
	TotalSupply float64 `json:"totalSupply"`
	Image       string  `json:"image"`
	MarketCap   float64 `json:"marketCap"`
	Des         string  `json:"des"`
	Volume24H   float64 `json:"volume24h"`
	PriceUSD    float64 `json:"priceUSD"`
	CreatedDate string  `json:"createdDate"`
	UpdatedDate string  `json:"updatedDate"`
}

func (crypto *Crypto) Insert() error {
	query := `INSERT INTO crypto (id, key, name, symbol, chainname, address, type, totalSupply, image,
		marketcap, volume24h, priceUSD, createddate, updateddate) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	_, err := db.PSQL.Exec(query, crypto.Id, crypto.Key, crypto.Name, crypto.Symbol, crypto.Chainname, crypto.Address, crypto.Type, crypto.TotalSupply, crypto.Image,
		crypto.MarketCap, crypto.Volume24H, crypto.PriceUSD, utils.Timestamp(), utils.Timestamp())

	return err
}

func (crypto *Crypto) Update() error {
	query := `UPDATE public.crypto SET totalsupply = $1, marketcap = $2, volume24h = $3, priceusd = $4 where key = $5;`
	_, err := db.PSQL.Exec(query, crypto.TotalSupply, crypto.MarketCap, crypto.Volume24H, crypto.PriceUSD, crypto.Key)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CryptoRepo) GetCryptos() error {
	query := `SELECT id, "name", image, symbol, priceusd FROM public.crypto order by marketCap desc;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		crypto := &Crypto{}
		err := rows.Scan(&crypto.Id, &crypto.Name, &crypto.Image, &crypto.Symbol, &crypto.PriceUSD)
		if err != nil {
			return err
		}
		repo.Cryptos = append(repo.Cryptos, *crypto)
	}
	return nil
}

func (crypto *Crypto) GetDetail() error {
	query := `SELECT id, "name", symbol, "type", totalsupply, 
	image, marketcap, volume24h, priceusd, des FROM public.crypto where id = $1;`
	err := db.PSQL.QueryRow(query, crypto.Id).Scan(
		&crypto.Id, &crypto.Name, &crypto.Symbol, &crypto.Type, &crypto.TotalSupply,
		&crypto.Image, &crypto.MarketCap, &crypto.Volume24H, &crypto.PriceUSD, &crypto.Des)
	if err != nil {
		return err
	}
	return nil
}
