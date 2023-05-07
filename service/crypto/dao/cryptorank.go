package dao

type CryptorankList struct {
	Data []struct {
		Rank             int    `json:"rank"`
		Key              string `json:"key"`
		Name             string `json:"name"`
		HasFundingRounds bool   `json:"hasFundingRounds"`
		Symbol           string `json:"symbol"`
		Type             string `json:"type"`
		RankHistory      struct {
			H24 int `json:"h24"`
			D7  int `json:"d7"`
			MTD int `json:"MTD"`
			D30 int `json:"d30"`
			D14 int `json:"d14"`
			Q1  int `json:"q1"`
			M3  int `json:"m3"`
			YTD int `json:"YTD"`
			Y1  int `json:"y1"`
			M6  int `json:"m6"`
		} `json:"rankHistory"`
		AthMarketCap struct {
			USD     float64 `json:"USD"`
			DateUSD string  `json:"dateUSD"`
		} `json:"athMarketCap"`
		LifeCycle       string  `json:"lifeCycle"`
		MaxSupply       int     `json:"maxSupply"`
		UnlimitedSupply bool    `json:"unlimitedSupply"`
		TotalSupply     float64 `json:"totalSupply"`
		Image           struct {
			Native string `json:"native"`
			Icon   string `json:"icon"`
			X60    string `json:"x60"`
			X150   string `json:"x150"`
		} `json:"image"`
		Category               string  `json:"category"`
		CategoryID             int     `json:"categoryId"`
		TagIds                 []int   `json:"tagIds"`
		IsTraded               bool    `json:"isTraded"`
		MarketDataNotAvailable bool    `json:"marketDataNotAvailable"`
		FullyDilutedMarketCap  float64 `json:"fullyDilutedMarketCap"`
		AvailableSupply        int     `json:"availableSupply"`
		MarketCap              float64 `json:"marketCap"`
		Volume24H              float64 `json:"volume24h"`
		NoData                 bool    `json:"noData"`
		Volatility             struct {
			USD float64 `json:"USD"`
			ETH float64 `json:"ETH"`
		} `json:"volatility"`
		Price struct {
			USD float64 `json:"USD"`
			BTC float64 `json:"BTC"`
			ETH float64 `json:"ETH"`
		} `json:"price"`
		HistPrices struct {
			Two4H struct {
				USD float64 `json:"USD"`
				BTC float64 `json:"BTC"`
				ETH float64 `json:"ETH"`
			} `json:"24H"`
			SevenD struct {
				USD float64 `json:"USD"`
				BTC float64 `json:"BTC"`
				ETH float64 `json:"ETH"`
			} `json:"7D"`
			OneY struct {
				USD float64 `json:"USD"`
				BTC float64 `json:"BTC"`
				ETH float64 `json:"ETH"`
			} `json:"1Y"`
			YTD struct {
				USD float64 `json:"USD"`
				BTC float64 `json:"BTC"`
				ETH float64 `json:"ETH"`
			} `json:"YTD"`
			SixM struct {
				USD float64 `json:"USD"`
				BTC float64 `json:"BTC"`
				ETH float64 `json:"ETH"`
			} `json:"6M"`
			ThreeM struct {
				USD float64 `json:"USD"`
				BTC float64 `json:"BTC"`
				ETH float64 `json:"ETH"`
			} `json:"3M"`
			Three0D struct {
				USD float64 `json:"USD"`
				BTC float64 `json:"BTC"`
				ETH float64 `json:"ETH"`
			} `json:"30D"`
		} `json:"histPrices"`
		AthPrice struct {
			BTC     float64 `json:"BTC"`
			ETH     float64 `json:"ETH"`
			USD     float64 `json:"USD"`
			Date    string  `json:"date"`
			DateBTC string  `json:"dateBTC"`
			DateETH string  `json:"dateETH"`
		} `json:"athPrice"`
		AtlPrice struct {
			BTC     float64 `json:"BTC"`
			ETH     float64 `json:"ETH"`
			USD     float64 `json:"USD"`
			DateBTC string  `json:"dateBTC"`
			DateETH string  `json:"dateETH"`
			DateUSD string  `json:"dateUSD"`
		} `json:"atlPrice"`
		Tokens []struct {
			PlatformName string `json:"platformName"`
			PlatformKey  string `json:"platformKey"`
			PlatformSlug string `json:"platformSlug"`
			ExplorerURL  string `json:"explorerUrl"`
			Address      string `json:"address"`
		} `json:"tokens"`
	} `json:"data"`
}
