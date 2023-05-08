package crawler

import (
	"base/pkg/log"
	"base/pkg/utils"
	"base/service/crypto/dao"
	"base/service/crypto/dto"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func CallZapperNewAddress(address string) (dao.Assets, error) {
	assets := dao.Assets{
		Address: address,
	}

	response, err := http.Get("https://zapper.xyz")
	if err != nil {
		log.Println(log.LogLevelError, "http.Get https://zapper.xyz", err.Error())
	}

	defer response.Body.Close()

	cookie := http.Cookie{}

	for _, ele := range response.Cookies() {
		// fmt.Printf("Name: %s\n", ele.Name)
		// fmt.Printf("Value: %s\n", ele.Value)
		cookie.Name = ele.Name
		cookie.Value = ele.Value
	}

	err = CallZapperPostAddress(address, cookie)
	if err != nil {
		log.Println(log.LogLevelError, "CallZapperNewAddress CallZapperPostAddress", err.Error())
	}

	//todo: wait for the job to run after 10s
	time.Sleep(10 * time.Second)

	assets, err = CallZapperGetAssets(address, cookie)
	if err != nil {
		log.Println(log.LogLevelError, "CallZapperNewAddress CallZapperGetAssets", err.Error())
	}

	return assets, nil
}

func CallZapperPostAddress(address string, cookie http.Cookie) error {
	request, err := http.NewRequest("POST", `https://zapper.xyz/z/v2/balances/tokens?addresses[]=`+address+`&networks[]=ethereum&networks[]=polygon&networks[]=optimism&networks[]=gnosis&networks[]=binance-smart-chain&networks[]=fantom&networks[]=avalanche&networks[]=arbitrum&networks[]=celo&networks[]=moonriver&networks[]=bitcoin&networks[]=aurora`, nil)

	request.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	request.Header.Add("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7`)
	request.Header.Add("accept-encoding", "deflate, br")
	request.Header.Add("accept-language", "en-US,en;q=0.9,vi;q=0.8,vi-VN;q=0.7")
	request.Header.Add("cache-control", "max-age=0")
	request.Header.Add("sec-ch-ua", `"Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("sec-ch-ua-platform", "macOS")
	request.Header.Add("sec-fetch-dest", "document")
	request.Header.Add("sec-fetch-mode", "navigate")
	request.Header.Add("sec-fetch-site", "none")
	request.Header.Add("sec-fetch-user", "?1")
	request.Header.Add("upgrade-insecure-requests", "1")

	request.AddCookie(&cookie)

	if err != nil {
		return err
	}

	// for _, ele := range request.Cookies() {
	// 	log.Println(log.LogLevelError, "Cookie:", ele.Name+"="+ele.Value)
	// }

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	// fmt.Println("status code post zapper", response.Status)s

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("CallZapperPostAddress post request failed statusCode: %d", response.StatusCode)
	}
	return nil
}

func CallZapperGetAssets(address string, cookie http.Cookie) (dao.Assets, error) {
	assets := dao.Assets{
		Address: address,
	}

	request, err := http.NewRequest("GET", `https://zapper.xyz/z/v2/balances/tokens?addresses[]=`+address+`&networks[]=ethereum&networks[]=polygon&networks[]=optimism&networks[]=gnosis&networks[]=binance-smart-chain&networks[]=fantom&networks[]=avalanche&networks[]=arbitrum&networks[]=celo&networks[]=moonriver&networks[]=bitcoin&networks[]=aurora`, nil)

	request.Header.Add("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7`)
	request.Header.Add("accept-encoding", "deflate, br")
	request.Header.Add("accept-language", "en-US,en;q=0.9,vi;q=0.8,vi-VN;q=0.7")
	request.Header.Add("cache-control", "max-age=0")
	request.Header.Add("cookie", "AMP_MKTG_8637806ba8=JTdCJTdE; AMP_8637806ba8=JTdCJTIyb3B0T3V0JTIyJTNBZmFsc2UlMkMlMjJkZXZpY2VJZCUyMiUzQSUyMjhhNzdkNzJjLThjNDktNGZmYi1iODNjLTBmYTExN2JmMWUyMCUyMiUyQyUyMmxhc3RFdmVudFRpbWUlMjIlM0ExNjgwODYwMDQ5NDE5JTJDJTIyc2Vzc2lvbklkJTIyJTNBMTY4MDg2MDA0OTM0MiUyQyUyMnVzZXJJZCUyMiUzQSUyMjIwODkzNTQyLTYwMzYtNDE5MS04ZmI4LTgzMTYzOTk3YTkxMiUyMiU3RA==; _gid=GA1.2.1701039656.1680860050; _ga_JJ1RCE5CXC=GS1.1.1680860049.9.0.1680860049.0.0.0; _ga=GA1.1.377045207.1680251888; __cf_bm=ynKFLiQ2.pqEiJA_4_iNfvvxjAsoZ2aYXitcczXHA3A-1680940391-0-ASYThzAJo6Nm1x0bNm81p9XK87Dj0xQdGPa1EILZIpdRVjQTf1V/gQKo7DA3CaiFYcNO7/ley9vFvGJGRr4RME0=")
	request.Header.Add("if-none-match", `W/"1llj64i"`)
	request.Header.Add("sec-ch-ua", `"Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("sec-ch-ua-platform", "macOS")
	request.Header.Add("sec-fetch-dest", "document")
	request.Header.Add("sec-fetch-mode", "navigate")
	request.Header.Add("sec-fetch-site", "none")
	request.Header.Add("sec-fetch-user", "?1")
	request.Header.Add("upgrade-insecure-requests", "1")
	request.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	request.AddCookie(&cookie)

	if err != nil {
		return assets, err
	}

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}}

	response, err := client.Do(request)
	if err != nil {
		return assets, err
	}

	// fmt.Println("status code CallZapperGetAssets", response.Status)

	if response.StatusCode != http.StatusOK {
		return assets, fmt.Errorf("CallZapperGetAssets get request failed statusCode: %d", response.StatusCode)

	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return assets, err
	}
	defer response.Body.Close()

	var data = make(map[string]any)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return assets, err
	}

	isLowerCaseAddress := false
	if data[address] == nil {
		if data[strings.ToLower(address)] == nil {
			return assets, errors.New("data uncached")
		} else {
			isLowerCaseAddress = true
		}
	}

	tokenBalances := []dto.TokenBalance{}
	if isLowerCaseAddress {
		err = utils.Mapping(data[strings.ToLower(address)], &tokenBalances)
	} else {
		err = utils.Mapping(data[address], &tokenBalances)
	}
	if err != nil {
		return assets, err
	}

	if len(tokenBalances) == 0 {
		return assets, errors.New("no data found")
	}

	// fmt.Println("len tokenBalances: ", len(tokenBalances))
	for _, balance := range tokenBalances {
		asset := &dao.Asset{
			Address:      balance.Address,
			TokenAddress: balance.Token.Address,
			TokenName:    balance.Token.Name,
			TokenSymbol:  balance.Token.Symbol,
			ChainName:    balance.Network,
			Amount:       balance.Token.BalanceRaw,
		}

		decimalString := fmt.Sprintf("%d", balance.Token.Decimals)
		asset.TokenDecimal = &decimalString

		priceString := fmt.Sprintf("%f", *balance.Token.Price)
		asset.TokenPriceUSD = &priceString

		binanceChainName := "binance"
		if *balance.Network == "binance-smart-chain" {
			asset.ChainName = &binanceChainName
		}
		// fmt.Println(i, asset)

		assets.Assets = append(assets.Assets, asset)
	}
	return assets, nil
}
