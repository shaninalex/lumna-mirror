package token

import (
	"gitlab.com/shaninalex/lumna/app/web"
)

func RegisterTokenController(router *web.Router) {
	h := NewTokenHandler()

	// list all user tokens
	router.GET("/api/v1/token", h.HandleTokenList)

	// delete a specific token
	router.DELETE("/api/v1/token/{tokenID}", h.HandleDeleteToken)

	// revoke a token
	router.GET("/api/v1/token/{tokenID}/revoke", h.HandleRevokeToken)
}
