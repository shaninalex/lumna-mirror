package auth

import (
	"log"
	"net/http"

	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type registerPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	payload, err := utils.BodyParser[registerPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	// NOTE: tmp, while services not ready
	log.Println(ctx, payload)

	// if _, err = s.userService.GetUserByEmail(ctx, payload.Email); err == nil {
	// 	web.Error(w, http.StatusBadRequest, errors.New("user with email already exists"))
	// 	return
	// }
	//
	// if _, err = s.userService.CreateUser(ctx, payload.Email, payload.Password); err != nil {
	// 	web.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	utils.Success(w, nil, "Registration Successful")
}
