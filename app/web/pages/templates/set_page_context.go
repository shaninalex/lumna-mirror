package templates

import (
	"context"
	"fmt"

	"github.com/a-h/templ"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/models"
)

type CmpFunc func(ctx PageContext) templ.Component

func InitTemplate(ctx context.Context, fn CmpFunc) templ.Component {
	pageContext := PageContext{}
	a := ctx.Value(internal.ContextIdentity)
	fmt.Println(a)
	if user, ok := ctx.Value(internal.ContextIdentity).(*models.Identity); ok {
		pageContext.User = &PageUser{
			ID:    user.ID.String(),
			Email: user.Email,
			Icon:  "/assets/img/7.png",
		}
	}

	return fn(pageContext)
}
