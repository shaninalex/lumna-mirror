package board

import (
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type BoardPageData struct {
	utils.BasePage
	Board models.Board
	Tasks []string // TODO: task model
}
