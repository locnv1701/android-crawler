package dao

import (
	"base/pkg/db"
)

type CryptoChangeRepo struct {
	CryptoChanges []CryptoChange `json:"crypto_changes"`
}

type CryptoChange struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	PriceUSD float64 `json:"priceUSD"`
}

func (repo *CryptoChangeRepo) GetCryptoChanges() error {
	query := `SELECT id, "name", priceusd FROM public.crypto_changes;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		cryptoChange := &CryptoChange{}
		err := rows.Scan(&cryptoChange.Id, &cryptoChange.Name, &cryptoChange.PriceUSD)
		if err != nil {
			return err
		}
		repo.CryptoChanges = append(repo.CryptoChanges, *cryptoChange)
	}
	return nil
}

func (cryptoChange *CryptoChange) CheckExist() (int, error) {
	query := `select count(*) from crypto_changes where id = $1`
	count := 0
	err := db.PSQL.QueryRow(query, cryptoChange.Id).Scan(&count)
	return count, err
}

func (cryptoChange *CryptoChange) Insert() error {
	query := `INSERT INTO crypto_changes (id, name, priceUSD) values ($1, $2, $3);`

	_, err := db.PSQL.Exec(query, cryptoChange.Id, cryptoChange.Name, cryptoChange.PriceUSD)

	return err
}

func (cryptoChange *CryptoChange) Update() error {
	query := `UPDATE public.crypto_changes SET name = $1, priceUSD = $2 where id = $3;`
	_, err := db.PSQL.Exec(query, cryptoChange.Name, cryptoChange.PriceUSD, cryptoChange.Id)
	if err != nil {
		return err
	}
	return nil
}
