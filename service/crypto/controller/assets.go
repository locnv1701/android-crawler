package controller

import (
	"base/pkg/log"
	"base/pkg/router"
	"base/service/crypto/crawler"
	"base/service/crypto/dao"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
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

	var total float64
	for _, asset := range listAsset.Assets {

		amountInt := new(big.Int)
		amountInt.SetString(*asset.Amount, 10)

		// Tạo big.Int với giá trị bằng 10^18
		divisor := new(big.Int)
		divisor.Exp(big.NewInt(10), big.NewInt(18), nil)

		// Chia số num cho 10^18
		amountInt.Div(amountInt, divisor)

		// In kết quả

		priceInt, err := strconv.ParseFloat(*asset.TokenPriceUSD, 10)
		if err != nil {
			fmt.Println(err)
		}

		total += float64(amountInt.Int64()) * priceInt
	}

	listAsset.Total = total

	router.ResponseSuccessWithData(w, "B.200", "Get assets succeessful", listAsset)
}
