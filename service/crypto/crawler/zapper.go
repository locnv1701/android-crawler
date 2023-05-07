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

func CrawlAssetByZapper() {
	addresses := []string{
		"0xfe5fb84cfe723f6c809e4a0e02212e955217e2b0",
		"0x2ce41ec427664971608ee87397e06e88bbe36a47",
		"0x1fb0b7f567d4bff19b3b3ba22538ed7b0934cbe2",
		"0x764b0adfa038721cf27b15e2eb94c22341fe5a2d",
		"0x602dbf79e0c4f8e90bb97bdd0a2f068b59924f4d",
		"0xf59c58337f66438de81570d8f76bfc67aa0788a5",
		"0xb7240af2af90b33c08ae9764103e35dce3638428",
		"0x8e9227f1d526d03789f8c09c2b5944eb53e6d5ec",
		"0x4a16ecf42d72528264b8313b604493eafef5d845",
		"0xcc6fcfb8b3988043e382a481d6bf482d68897024",
		"0x13288f3f4d784377a1d94eea529e5955c57a54e5",
		"0x04d9cc35d5bf408a7d442fb45d235667144e4d92",
		"0xa178ff321335bb777a7e21a56376592f69556b9c",
		"0xca8d6f69d8f32516a109df68b623452cc9f5e64d",
		"0x55152e90293b52ed08f4cdc9b60aa4ebefa3c635",
		"0x525fc5d4fca3bac67b150cd7821816dccf4471b7",
		"0x0bd95534c6d270d3f867f17958a67a3f89d84104",
		"0x301407427168fb51bcc927b9fb76dcd88fe45681",
		"0x985124fa22125fcd4480876c1e8191167f4efcb4",
		"0x856442d44ffde0fc090e7a4d1b92c9dac57af4b7",
		"0xc92d500dcbe98df16a0906248a031cca80676593",
		"0xaab72305897ddf54ef56040e9181f29c668f1c6e",
		"0x92f3da97f4a17708233452e6a66cde76a80a1938",
		"0x7422f5f528cc36ea48f6cd8cda5a9816e4573ede",
		"0x21e90dc3a909c72aaf97678dbb063fbbc32aef0c",
		"0x42c09c68d5edf5ddf89175bcb9d6847d3dfb6669",
		"0xfe3348559721ba6801accd438c9576b1513e7c52",
		"0xfcba81c19b20d2f2d7cea6e10647803ede7697c0",
		"0x0fc19b69f40373c2abd10eeb37952c1197a3302e",
		"0x453a1313270865a7cb0acd21e31db44f1a64e908",
		"0x72c2a6ab872b51ec6e1d4314febf8f55587bbff5",
		"0x6dd5a9f47cfbc44c04a0a4452f0ba792ebfbcc9a",
		"0xd225a255ab6fa11b9673ad0516b579f4cdbaffe5",
		"0xb2698c2d99ad2c302a95a8db26b08d17a77cedd4",
		"0xb4d1fa72c69099c554848de5be2735da315678c3",
		"0x6edf7f5283725c953ee64317f66188af1184b033",
		"0x7153d2ef9f14a6b1bb2ed822745f65e58d836c3f",
		"0xe5968797468ef767101b761d431fce14abffdbb4",
		"0xc2a39291acf6ae4ed29c6fcc15f1d54141f8d833",
		"0x16a0772b17ae004e6645e0e95bf50ad69498a34e",
		"0x1dfbbb649e4c89fd8354f01303ccd668957ce9bf",
		"0xb7330c1290ac1e8cbdcf60f82b423216e53742f7",
		"0x69ab22c316c32eb03494f6998be36bdc2379f87a",
		"0xf52284a180d11257a87d528009c12d59ff109a52",
		"0x382591e7217b435e8e884cdbf415fe377a6fe29e",
		"0x3113b42b97de26116b2957288ea94120d5c3e84b",
		"0x550d74e4f43d447ba104662f44bdb47036c0d2c9",
		"0xd07957a1183d956e804fb125c70f777ccffcb140",
	}

	for _, address := range addresses {
		assets, err := CallZapperNewAddress(address)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(assets)
		err = assets.InsertListAsset()
		if err != nil {
			fmt.Println(err)
		}
	}

}

func CallZapperNewAddress(address string) (dao.Assets, error) {
	assets := dao.Assets{
		Address: address,
	}

	err := CallZapperPostAddress(address)
	if err != nil {
		log.Println(log.LogLevelError, "CallZapperNewAddress CallZapperPostAddress", err.Error())
	}

	//todo: wait for the job to run after 10s
	time.Sleep(10 * time.Second)

	assets, err = CallZapperGetAssets(address)
	if err != nil {
		log.Println(log.LogLevelError, "CallZapperNewAddress CallZapperGetAssets", err.Error())
	}

	return assets, nil
}

func CallZapperPostAddress(address string) error {
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

func CallZapperGetAssets(address string) (dao.Assets, error) {
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
