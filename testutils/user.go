package testutils

import (
	"fmt"
	"log"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/services/auth"
	"gorm.io/gorm"
)

func User(user models.Identity, db *gorm.DB) *models.Identity {
	if result := db.Create(&user); result.Error != nil {
		panic(result.Error)
	}

	pwdHash, err := auth.CreatePasswordHash("111")
	if err != nil {
		log.Fatal(err)
	}

	credential := models.Credential{
		IdentityID:     user.ID,
		Provider:       "local",
		ProviderUserID: utils.Pointer(fmt.Sprintf("%d", user.ID)),
		Email:          &user.Email,
		PasswordHash:   utils.Pointer(pwdHash),
	}
	if result := db.Create(&credential); result.Error != nil {
		panic(result.Error)
	}
	return &user
}
