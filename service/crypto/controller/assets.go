package controller

import (
	"base/pkg/log"
	"base/pkg/router"
	"base/service/crypto/crawler"
	"base/service/crypto/dao"
	"fmt"
	"net/http"
)

// type ExplorerAddressDTO struct {
// 	Address string     `json:"address"`
// 	Assets  dao.Assets `json:"assets"`
// }

func GetAssets(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if len(address) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing address param!")
		return
	}

	listAsset := dao.Assets{
		Address: address,
	}

	count, err := listAsset.CheckExist()
	if err != nil {
		log.Println(log.LogLevelError, "check", err.Error())
	}

	if count == 0 {
		fmt.Println("new address")
		assets, err := crawler.CallZapperNewAddress(address)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(assets)
		err = assets.InsertListAsset()
		if err != nil {
			fmt.Println(err)
		}
	}

	err = listAsset.GetAllAsset()
	if err != nil {
		log.Println(log.LogLevelError, "ExplorerAddress listAsset.GetAllAsset()", err.Error())
	}

	router.ResponseSuccessWithData(w, "B.200", "Get assets succeessful", listAsset)
}
