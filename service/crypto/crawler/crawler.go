package crawler

import (
	"base/service/crypto/dao"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CallApiCryptorank() {
	url := "https://api.cryptorank.io/v0/coins?locale=en"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	// Read response body
	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// Unmarshal the response body into the asset struct

	cryptorankList := dao.CryptorankList{}

	err = json.Unmarshal(body, &cryptorankList)

	fmt.Println("len", len(cryptorankList.Data))

	for _, cryptorank := range cryptorankList.Data[:100] {
		crypto := dao.Crypto{
			Id:          cryptorank.Rank,
			Key:         cryptorank.Key,
			Name:        cryptorank.Name,
			Symbol:      cryptorank.Symbol,
			Type:        cryptorank.Type,
			TotalSupply: cryptorank.TotalSupply,
			Image:       cryptorank.Image.Native,
			MarketCap:   cryptorank.MarketCap,
			Volume24H:   cryptorank.Volume24H,
			PriceUSD:    cryptorank.Price.USD,
		}

		err := crypto.Insert()

		if err != nil {
			fmt.Println("err insert", err)
		}
	}

	fmt.Println(err)
}
