package utils

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/models"
)

func GetBasePage(ctx context.Context) BasePage {
	page := BasePage{}

	if user, ok := ctx.Value(internal.ContextIdentity).(*models.Identity); ok {
		page.User = &PageUser{
			ID:    user.ID.String(),
			Email: user.Email,
			Icon:  "/assets/img/7.png",
		}
	}

	return page
}
