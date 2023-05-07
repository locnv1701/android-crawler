package controller

import (
	"base/pkg/log"
	"base/pkg/router"
	"base/service/crypto/dao"
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

	err := listAsset.GetAllAsset()
	if err != nil {
		log.Println(log.LogLevelError, "ExplorerAddress listAsset.GetAllAsset()", err.Error())
	}

	router.ResponseSuccessWithData(w, "B.200", "Get assets succeessful", listAsset)
}
