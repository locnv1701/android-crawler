package dao

import (
	"base/pkg/db"
	"base/pkg/log"
)

type Assets struct {
	Address string   `json:"address"`
	Assets  []*Asset `json:"assets"`
}

type Asset struct {
	Address       *string `json:"address"`
	TokenAddress  *string `json:"tokenAddress"`
	TokenName     *string `json:"tokenName"`
	TokenSymbol   *string `json:"tokenSymbol"`
	SmallLogo     *string `json:"smallLogo"`
	TokenDecimal  *string `json:"tokenDecimal"`
	TokenPriceUSD *string `json:"tokenPriceUSD"`
	ChainName     *string `json:"chainName"`
	Amount        *string `json:"amount"`
}

func (assets *Assets) GetAllAsset() error {
	query := `select tokens.address, tokens.name, tokens.symbol, tokens.smalllogo, tokens."decimal", tokens.priceUSD,
	wallet_assets.chainname, wallet_assets.amount from tokens join 
	(select tokenaddressid, chainname, amount from assets where addressid = ('x'||md5($1))::bit(32)::int8) as wallet_assets 
	on tokens.chainname = wallet_assets.chainname and tokens.addressid = wallet_assets.tokenaddressid
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
		err := rows.Scan(&asset.TokenAddress, &asset.TokenName, &asset.TokenSymbol, &asset.SmallLogo, &asset.TokenDecimal, &asset.TokenPriceUSD,
			&asset.ChainName, &asset.Amount)
		if err != nil {
			return err
		}
		assets.Assets = append(assets.Assets, &asset)
	}
	return nil
}

func (assets *Assets) InsertListAsset() error {
	query := `INSERT INTO assets (addressid, tokenaddressid, chainname, amount) 
	VALUES(('x'||md5($1))::bit(32)::int8, ('x'||md5($2))::bit(32)::int8, $3, $4);`

	for _, asset := range assets.Assets {
		_, err := db.PSQL.Exec(query, asset.Address, asset.TokenAddress, asset.ChainName, asset.Amount)
		if err != nil {
			log.Println(log.LogLevelError, "asset.Insert()"+":"+*asset.Address+" "+*asset.TokenAddress+" "+*asset.ChainName+" "+*asset.Amount, err.Error())
		}
	}
	return nil
}
func (asset *Asset) Insert() error {
	query := `INSERT INTO assets (addressid, tokenaddressid, chainname, amount) 
		VALUES(('x'||md5($1))::bit(32)::int8, ('x'||md5($2))::bit(32)::int8, $3, $4);`
	_, err := db.PSQL.Exec(query, asset.Address, asset.TokenAddress, asset.ChainName, asset.Amount)
	if err != nil {
		log.Println(log.LogLevelError, "asset.Insert()"+":"+*asset.Address+" "+*asset.TokenAddress+" "+*asset.ChainName+" "+*asset.Amount, err.Error())
	}
	return err
}

func (assets *Assets) GetMoreInfoAsset() error {

	query := `select name, symbol, smalllogo, "decimal", priceUSD from tokens where
	addressid = ('x'||md5($1))::bit(32)::int8 and chainname = $2`

	for _, asset := range assets.Assets {
		err := db.PSQL.QueryRow(query, asset.TokenAddress, asset.ChainName).Scan(
			&asset.TokenName, &asset.TokenSymbol, &asset.SmallLogo, &asset.TokenDecimal, &asset.TokenPriceUSD)
		if err != nil {
			log.Println(log.LogLevelError, "GetMoreInfoAsset db.PSQLWrite().QueryRow "+*asset.TokenAddress+" "+*asset.ChainName, err.Error())
		}
	}

	return nil
}
