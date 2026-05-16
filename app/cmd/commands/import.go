package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/auth"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/pkg/persistence"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type MockDbDataSchema struct {
	Identities []struct {
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		Credentials []struct {
			Provider string `json:"provider"`
			Password string `json:"password"`
		} `json:"credentials"`
	} `json:"identities"`
	Workspaces []struct {
		Title    string `json:"title"`
		Email    string `json:"email"`
		Active   bool   `json:"active"`
		Projects []struct {
			Title  string `json:"title"`
			Boards []struct {
				Title   string `json:"title"`
				Columns []struct {
					Title string `json:"title"`
					Order int    `json:"order"`
					Tasks []struct {
						Title string `json:"title"`
						Order int    `json:"order"`
						Body  string `json:"body"`
					} `json:"tasks"`
				} `json:"columns"`
			} `json:"boards"`
		} `json:"projects"`
	} `json:"workspaces"`
}

func NewImportRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import [path_to_file]",
		Short: "Import db",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			// read resources/mock_db.json
			var payload MockDbDataSchema
			data, err := os.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
			if err := json.Unmarshal(data, &payload); err != nil {
				panic(err)
			}

			if err := c.Invoke(importDB(payload)); err != nil {
				panic(err)
			}
		},
	}
	return cmd
}

func importDB(payload MockDbDataSchema) func(db *gorm.DB) {
	return func(database *gorm.DB) {
		// create identity+credential
		for idx, idn := range payload.Identities {
			identity := models.Identity{
				Email:    idn.Email,
				FullName: idn.FullName,
			}
			if result := database.Create(&identity); result.Error != nil {
				panic(result.Error)
			}
			fmt.Printf("%d. Identity: %s\n", idx, identity.Email)
			for cidx, crd := range idn.Credentials {
				pwd, _ := auth.CreatePasswordHash(crd.Password)
				credential := models.Credential{
					IdentityID:   identity.ID,
					Provider:     crd.Provider,
					PasswordHash: utils.Pointer(pwd),
					Email:        &idn.Email,
				}
				if result := database.Create(&credential); result.Error != nil {
					panic(result.Error)
				}
				fmt.Printf("\t%d. credential for: %s\n", cidx, identity.Email)
			}
		}

		for _, wpd := range payload.Workspaces {
			wp := models.Workspace{
				Title:      wpd.Title,
				OwnerEmail: wpd.Email,
				Active:     wpd.Active,
			}
			if result := database.Create(&wp); result.Error != nil {
				panic(result.Error)
			}

			// create projects
			for pidx, _project := range wpd.Projects {
				project := models.Project{Title: _project.Title, WorkspaceID: wp.ID}
				if result := database.Create(&project); result.Error != nil {
					panic(result.Error)
				}
				fmt.Printf("%d. Project: %s\n", pidx, project.Title)
				for bi, _board := range _project.Boards {
					board := models.List{
						Title:       _board.Title,
						ProjectID:   project.ID,
						WorkspaceID: wp.ID,
					}
					if result := database.Create(&board); result.Error != nil {
						panic(result.Error)
					}
					fmt.Printf("\t%d. Board: %s\n", bi, board.Title)
					for li, _list := range _board.Columns {
						column := models.Status{
							Title:       _list.Title,
							ListID:      board.ID,
							Order:       uint(li),
							ProjectID:   project.ID,
							WorkspaceID: wp.ID,
						}
						if result := database.Create(&column); result.Error != nil {
							panic(result.Error)
						}
						fmt.Printf("\t\t%d. List: %s\n", li, column.Title)
						for ti, _task := range _list.Tasks {
							task := models.Task{
								Title:       _task.Title,
								ColumnID:    column.ID,
								Order:       uint(ti),
								Body:        _task.Body,
								ProjectID:   project.ID,
								ListID:      board.ID,
								WorkspaceID: wp.ID,
							}
							if result := database.Create(&task); result.Error != nil {
								panic(result.Error)
							}
							fmt.Printf("\t\t\t%d. Task: %s\n", ti, task.Title)
						}
					}
				}
			}
		}
	}
}
