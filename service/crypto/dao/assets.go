package dao

import (
	"base/pkg/db"
	"base/pkg/log"
)

type Assets struct {
	Address string   `json:"address"`
	Assets  []*Asset `json:"assets"`
	Total   float64  `json:"total"`
}

type Asset struct {
	Address       *string `json:"address"`
	TokenAddress  *string `json:"tokenAddress"`
	TokenName     *string `json:"tokenName"`
	TokenSymbol   *string `json:"tokenSymbol"`
	Image         *string `json:"image"`
	TokenPriceUSD *string `json:"tokenPriceUSD"`
	ChainName     *string `json:"chainName"`
	Amount        *string `json:"amount"`
}

func (assets *Assets) CheckExist() (int, error) {
	query := `select count(*) from assets where address = $1`
	count := 0
	err := db.PSQL.QueryRow(query, assets.Address).Scan(&count)
	return count, err
}

func (assets *Assets) GetAllAsset() error {
	query := `select crypto.address, crypto.name, crypto.symbol, crypto.image, crypto.priceUSD,
	wallet_assets.chainname, wallet_assets.amount from crypto join 
	(select tokenaddress, chainname, amount from assets where address = $1) as wallet_assets
	on crypto.chainname = wallet_assets.chainname and crypto.address = wallet_assets.tokenaddress
	`

	rows, err := db.PSQL.Query(query, assets.Address)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var asset = Asset{
			Address: &assets.Address,
		}
		err := rows.Scan(&asset.TokenAddress, &asset.TokenName, &asset.TokenSymbol, &asset.Image, &asset.TokenPriceUSD,
			&asset.ChainName, &asset.Amount)
		if err != nil {
			return err
		}
		assets.Assets = append(assets.Assets, &asset)
	}
	return nil
}

func (assets *Assets) InsertListAsset() error {
	query := `INSERT INTO assets (address, tokenaddress, chainname, amount) 
	VALUES($1, $2, $3, $4);`

	for _, asset := range assets.Assets {
		_, err := db.PSQL.Exec(query, asset.Address, asset.TokenAddress, asset.ChainName, asset.Amount)
		if err != nil {
			log.Println(log.LogLevelError, "asset.Insert()"+":"+*asset.Address+" "+*asset.TokenAddress+" "+*asset.ChainName+" "+*asset.Amount, err.Error())
		}
	}
	return nil
}
func (asset *Asset) Insert() error {
	query := `INSERT INTO assets (address, tokenaddress, chainname, amount) 
		VALUES($1, $2, $3, $4);`
	_, err := db.PSQL.Exec(query, asset.Address, asset.TokenAddress, asset.ChainName, asset.Amount)
	if err != nil {
		log.Println(log.LogLevelError, "asset.Insert()"+":"+*asset.Address+" "+*asset.TokenAddress+" "+*asset.ChainName+" "+*asset.Amount, err.Error())
	}
	return err
}

func (assets *Assets) GetMoreInfoAsset() error {

	query := `select name, symbol, image, priceUSD from crypto where
	address = $1 and chainname = $2`

	for _, asset := range assets.Assets {
		err := db.PSQL.QueryRow(query, asset.TokenAddress, asset.ChainName).Scan(
			&asset.TokenName, &asset.TokenSymbol, &asset.Image, &asset.TokenPriceUSD)
		if err != nil {
			log.Println(log.LogLevelError, "GetMoreInfoAsset db.PSQLWrite().QueryRow "+*asset.TokenAddress+" "+*asset.ChainName, err.Error())
		}
	}

	return nil
}
