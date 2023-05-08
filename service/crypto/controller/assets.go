package controller

import (
	"base/pkg/router"
	"base/service/crypto/dao"
	"fmt"
	"net/http"
)

func AddAsset(w http.ResponseWriter, r *http.Request) {
	asset := &dao.Asset{}
	err := asset.AddAsset()
	if err != nil {
		fmt.Println(err)
	}
	router.ResponseSuccess(w, "B.200", "Add assets succeessful")
}

func GetAssets(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if len(userId) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing userId param!")
		return
	}

	repo := &dao.AssetRepo{
		UserId: userId,
	}

	err := repo.GetAssets()
	if err != nil {
		fmt.Println(err)
	}

	total := float64(0)
	for i, asset := range repo.Assets {
		crypto := &dao.Crypto{
			Symbol: asset.Symbol,
		}

		err = crypto.GetDetail()
		if err != nil {
			fmt.Println(err)
		}

		asset.Image = crypto.Image
		asset.Name = crypto.Name
		repo.Assets[i] = asset

		total += crypto.PriceUSD * asset.Amount
	}
	repo.Total = total
	router.ResponseSuccessWithData(w, "B.200", "Get assets succeessful", repo)
}

func UpdateAmount(w http.ResponseWriter, r *http.Request) {
	asset := &dao.Asset{}
	err := asset.Update()
	if err != nil {
		fmt.Println(err)
	}
	router.ResponseSuccess(w, "B.200", "Update assets succeessful")
}
