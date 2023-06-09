package service

import (
	"base/pkg/router"
	"base/service/crypto"
	"base/service/index"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {

	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/crypto_service", crypto.CryptosService)

}
