package crypto

import (
	"base/service/crypto/controller"

	"github.com/go-chi/chi"
)

var CryptosService = chi.NewRouter()

func init() {

	CryptosService.Group(func(r chi.Router) {
		CryptosService.Get("/list", controller.GetListCryptos)

		CryptosService.Post("/add-asset", controller.AddAsset)
		CryptosService.Patch("/update-asset", controller.UpdateAmount)
		CryptosService.Get("/assets", controller.GetAssets)
	})
}
