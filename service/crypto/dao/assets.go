package dao

import (
	"base/pkg/db"
	"fmt"
)

type AssetRepo struct {
	Total  float64 `json:"total"`
	UserId string  `json:"userId"`
	Assets []Asset `json:"assets"`
}

type Asset struct {
	UserId string  `json:"userId"`
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
	Image  string  `json:"image"`
	Name   string  `json:"name"`
}

func (asset *Asset) AddAsset() error {
	fmt.Println(asset)
	query := `INSERT INTO public.asset (user_id, symbol, amount) VALUES($1, $2, $3);`
	_, err := db.PSQL.Exec(query, asset.UserId, asset.Symbol, asset.Amount)
	return err
}

func (asset *Asset) Update() error {
	query := `UPDATE public.asset SET amount = $1 where user_Id = $2 and symbol = $3;`
	_, err := db.PSQL.Exec(query, asset.Amount, asset.UserId, asset.Symbol)
	return err
}

func (repo *AssetRepo) GetAssets() error {
	query := `SELECT symbol, amount FROM public.asset where user_id = $1;`
	rows, err := db.PSQL.Query(query, repo.UserId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		asset := &Asset{}
		err := rows.Scan(&asset.Symbol, &asset.Amount)
		if err != nil {
			return err
		}
		repo.Assets = append(repo.Assets, *asset)
	}
	return nil
}
