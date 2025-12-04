package token

import "net/http"

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{}
}

type TokenHandler struct {
}

func (s *TokenHandler) HandleTokenList(w http.ResponseWriter, r *http.Request) {

}

func (s *TokenHandler) HandleDeleteToken(w http.ResponseWriter, r *http.Request) {

}

func (s *TokenHandler) HandleRevokeToken(w http.ResponseWriter, r *http.Request) {

}
