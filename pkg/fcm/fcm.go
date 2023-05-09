package fcm

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

const CredentialsFileName = "login-test-34d0f-firebase-adminsdk-teu7z-8299694c7b.json"

func PublishMessage(messages []*messaging.Message) {
	opt := option.WithCredentialsFile(CredentialsFileName)
	app, err := firebase.NewApp(context.TODO(), nil, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fcmClient, err := app.Messaging(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, mess := range messages {
		_, err = fcmClient.Send(context.TODO(), mess)
		if err != nil {
			fmt.Println(err)
		}
	}
}
