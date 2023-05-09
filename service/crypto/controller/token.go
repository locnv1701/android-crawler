package controller

import (
	"base/pkg/router"
	"base/service/crypto/dao"
	"fmt"
	"net/http"
)

func SaveDeviceToken(w http.ResponseWriter, r *http.Request) {
	dt := r.URL.Query().Get("token")
	if len(dt) <= 0 {
		router.ResponseBadRequest(w, "B.400", "Missing device token!")
		return
	}

	token := &dao.Token{
		DeviceToken: dt,
	}

	exist, err := token.CheckExist()
	if err != nil {
		fmt.Println(err)
	}
	if exist == 0 {
		err = token.Insert()
		if err != nil {
			fmt.Println(err)
		}
	}

	router.ResponseSuccessWithData(w, "B.200", "Save token succeessful", token)
}
