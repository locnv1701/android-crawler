package dto

type TokenBalance struct {
	Address *string `json:"address"`
	Network *string `json:"network"`
	Token   struct {
		Address    *string  `json:"address"`
		Name       *string  `json:"name"`
		Symbol     *string  `json:"symbol"`
		Decimals   *int     `json:"decimals"`
		Price      *float64 `json:"price"`
		BalanceRaw *string  `json:"balanceRaw"`
	} `json:"token"`
}
