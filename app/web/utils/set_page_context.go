package utils

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/csrf"
)

func BasePageData(c *gin.Context, title string) BasePage {
	base := BasePage{
		Title:     title,
		CsrfToken: csrf.GetToken(c),
	}

	if user, ok := c.Request.Context().Value(internal.ContextIdentity).(*models.Identity); ok {
		base.User = &PageUser{
			ID:    user.ID.String(),
			Email: user.Email,
			Icon:  "/assets/img/7.png", // TODO: get from identity details
		}
	}

	return base
}
