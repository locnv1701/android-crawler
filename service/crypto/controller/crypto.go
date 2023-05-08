package controller

import (
	"base/pkg/router"
	"base/service/crypto/dao"
	"fmt"
	"net/http"
)

func GetListCryptos(w http.ResponseWriter, r *http.Request) {
	repo := &dao.CryptoRepo{}
	err := repo.GetCryptos()
	if err != nil {
		fmt.Println(err)
	}
	router.ResponseSuccessWithData(w, "B.200", "Get list cryptos succeessful", repo.Cryptos)
}
