package crawler

import (
	"base/pkg/fcm"
	"base/service/crypto/dao"
	"fmt"

	"firebase.google.com/go/v4/messaging"
	"gopkg.in/robfig/cron.v2"
)

func NotificationCronjob() {
	c := cron.New()
	// c.AddFunc("0 0 9 * * *", ComparePriceChange)
	c.AddFunc("@every 0h0m10s", ComparePriceChange)
	c.Start()
}

func ComparePriceChange() {
	repo := &dao.CryptoRepo{}
	err := repo.GetCryptos()
	if err != nil {
		fmt.Println(err)
		return
	}

	changeRepo := &dao.CryptoChangeRepo{}
	err = changeRepo.GetCryptoChanges()
	if err != nil {
		fmt.Println(err)
		return
	}

	topCryptoLen := 10
	if len(repo.Cryptos) < 10 {
		topCryptoLen = len(repo.Cryptos)
	}

	cryptoNames := make([]string, 0, 10)
	for i := 0; i <= topCryptoLen; i++ {
		if isBigChange(repo.Cryptos[i], changeRepo) {
			cryptoNames = append(cryptoNames, repo.Cryptos[i].Name)
		}

		cryptoChange := &dao.CryptoChange{
			Id:       repo.Cryptos[i].Id,
			Name:     repo.Cryptos[i].Name,
			PriceUSD: repo.Cryptos[i].PriceUSD,
		}

		cnt, _ := cryptoChange.CheckExist()
		if cnt > 0 {
			err = cryptoChange.Update()
			if err != nil {
				fmt.Println("Insert change: ", err)
			}
		} else {
			err = cryptoChange.Insert()
			if err != nil {
				fmt.Println("Insert change: ", err)
			}
		}
	}

	title := "Crypto App"
	body := "big change in the exchange rate!"
	switch len(cryptoNames) {
	case 0:
		body = "Let check the exchange rate!"
	case 1:
		body = cryptoNames[0] + " has " + body
	default:
		body = getNames(cryptoNames) + " have " + body
	}

	tokenRepo := &dao.TokenRepo{}
	err = tokenRepo.GetTokens()
	if err != nil {
		fmt.Println(err)
		return
	}

	messages := make([]*messaging.Message, 0, 100)
	for _, token := range tokenRepo.Tokens {
		mess := messaging.Message{
			Token: token.DeviceToken,
			Data: map[string]string{
				"title": title,
				"body":  body,
			},
			Android: &messaging.AndroidConfig{
				Notification: &messaging.AndroidNotification{
					Title: title,
					Body:  body,
				},
			},
		}
		messages = append(messages, &mess)
	}
	fcm.PublishMessage(messages)
}

func isBigChange(crypto dao.Crypto, cryptoChangeRepo *dao.CryptoChangeRepo) bool {
	for _, cryptoChange := range cryptoChangeRepo.CryptoChanges {
		if cryptoChange.Id == crypto.Id {
			return (crypto.PriceUSD-cryptoChange.PriceUSD)/cryptoChange.PriceUSD >= 0.1
		}
	}
	return true
}

func getNames(names []string) string {
	res := names[0]
	for i := 1; i < len(names)-1; i++ {
		res += ", " + names[i]
	}
	res += " and " + names[len(names)-1]
	return res
}
