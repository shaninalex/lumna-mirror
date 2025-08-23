package database

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"gitlab.com/shaninalex/jajirra/models"
	"gorm.io/gorm"
)

func GenerateUniqueUserCode(ctx context.Context, db *gorm.DB, maxAttempts int) (string, error) {
	for i := 0; i < maxAttempts; i++ {
		base := strings.ToLower(randomdata.SillyName())
		code := fmt.Sprintf("%s%d", base, rand.IntN(100_000))

		var existing models.User
		err := db.WithContext(ctx).Where("code = ?", code).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code, nil
		}

		if err != nil {
			return "", err
		}

		// Code exists. Retry
	}

	fallback := fmt.Sprintf("user-%s", shortID())
	return fallback, nil
}

const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

func shortID() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}
