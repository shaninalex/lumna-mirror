package user

import "net/http"

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type UserHandler struct {
}

func (s *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {

}

func (s *UserHandler) HandleUpdateSettings(w http.ResponseWriter, r *http.Request) {

}

func (s *UserHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {

}
