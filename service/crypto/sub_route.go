package crypto

import (
	"base/pkg/log"
	"base/service/crypto/controller"
	"base/service/crypto/dao"
	"fmt"

	"github.com/go-chi/chi"
)

var CryptosService = chi.NewRouter()

func init() {

	listAsset := dao.Assets{
		Address: "0xfe5fb84cfe723f6c809e4a0e02212e955217e2b0",
	}

	count, err := listAsset.CheckExist()
	if err != nil {
		log.Println(log.LogLevelError, "check", err.Error())
	}

	fmt.Println(count, err)

	CryptosService.Group(func(r chi.Router) {
		CryptosService.Get("/list", controller.GetListCryptos)

		CryptosService.Get("/assets", controller.GetAssets)
	})
}
