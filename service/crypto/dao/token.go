package dao

import (
	"base/pkg/db"
	"base/pkg/log"
	"fmt"
)

type TokenRepo struct {
	Tokens []Token `json:"device_tokens"`
}

type Token struct {
	DeviceToken string `json:"device_token"`
}

func (token *Token) CheckExist() (int, error) {
	query := `select count(*) from tokens where device_token = $1`
	count := 0
	err := db.PSQL.QueryRow(query, token.DeviceToken).Scan(&count)
	return count, err
}

func (token *Token) Insert() error {
	query := `INSERT INTO tokens VALUES($1);`
	_, err := db.PSQL.Exec(query, token.DeviceToken)
	if err != nil {
		log.Println(log.LogLevelError, "token.Insert()"+": "+token.DeviceToken, err.Error())
	}
	return err
}

func (repo *TokenRepo) GetTokens() error {
	query := `SELECT device_token FROM public.tokens;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		token := &Token{}
		err := rows.Scan(&token.DeviceToken)
		fmt.Println("select dt from tokens where ")
		if err != nil {
			return err
		}
		repo.Tokens = append(repo.Tokens, *token)
	}
	return nil
}
