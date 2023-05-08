package controller

import (
	"base/pkg/router"
	"base/service/crypto/dao"
	"fmt"
	"net/http"
	"strconv"
)

func GetListCryptos(w http.ResponseWriter, r *http.Request) {
	repo := &dao.CryptoRepo{}
	err := repo.GetCryptos()
	if err != nil {
		fmt.Println(err)
	}
	router.ResponseSuccessWithData(w, "B.200", "Get list cryptos succeessful", repo.Cryptos)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing address param!")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	crypto := &dao.Crypto{
		Id: idInt,
	}
	err = crypto.GetDetail()
	if err != nil {
		fmt.Println(err)
	}
	router.ResponseSuccessWithData(w, "B.200", "Get detail succeessful", crypto)
}
