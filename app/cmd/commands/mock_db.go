package commands

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/auth/local"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/internal/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

type MockDbDataSchema struct {
	Identities []struct {
		ID          string `json:"id"`
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		Credentials []struct {
			Provider string `json:"provider"`
			Password string `json:"password"`
		} `json:"credentials"`
	} `json:"identities"`
	Projects []struct {
		Title  string `json:"title"`
		Boards []struct {
			Title string `json:"title"`
			Lists []struct {
				Title string `json:"title"`
				Order int    `json:"order"`
				Tasks []struct {
					Title string `json:"title"`
					Order int    `json:"order"`
				} `json:"tasks"`
			} `json:"lists"`
		} `json:"boards"`
	} `json:"projects"`
}

func NewMockDBRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mockdb [path_to_file]",
		Short: "Create mock data",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				panic(err)
			}
			// read resources/mock_db.json
			var payload MockDbDataSchema
			data, err := os.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
			if err := json.Unmarshal(data, &payload); err != nil {
				panic(err)
			}

			// create identity+credential
			for _, idn := range payload.Identities {
				identity := models.Identity{
					ID:       uuid.MustParse(idn.ID),
					Email:    idn.Email,
					FullName: idn.FullName,
				}
				if result := c.DB().Create(&identity); result.Error != nil {
					panic(result.Error)
				}
				for _, crd := range idn.Credentials {
					pwd, _ := local.CreatePasswordHash(crd.Password)
					credential := models.Credential{
						IdentityID:   identity.ID,
						Provider:     crd.Provider,
						PasswordHash: utils.Pointer(pwd),
						Email:        &idn.Email,
					}
					if result := c.DB().Create(&credential); result.Error != nil {
						panic(result.Error)
					}
				}
			}

			// create projects
			for _, _project := range payload.Projects {
				project := models.Project{Title: _project.Title}
				if result := c.DB().Create(&project); result.Error != nil {
					panic(result.Error)
				}
				for bi, _board := range _project.Boards {
					board := models.Board{
						Title:     _board.Title,
						ProjectID: project.ID,
						Order:     uint(bi),
					}
					if result := c.DB().Create(&board); result.Error != nil {
						panic(result.Error)
					}

					for li, _list := range _board.Lists {
						boardList := models.BoardList{
							Title:   _list.Title,
							BoardID: board.ID,
							Order:   uint(li),
						}
						if result := c.DB().Create(&boardList); result.Error != nil {
							panic(result.Error)
						}
						for ti, _task := range _list.Tasks {
							task := models.Task{
								Title:       _task.Title,
								BoardListID: boardList.ID,
								Order:       uint(ti),
								ProjectID:   project.ID,
							}
							if result := c.DB().Create(&task); result.Error != nil {
								panic(result.Error)
							}
						}
					}
				}
			}
		},
	}
	return cmd
}
