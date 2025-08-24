package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/database/seed"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
)

type DevCliAction string

var (
	//DevCliActionSeed this action seed data into development database
	DevCliActionSeed DevCliAction = "seed"

	// DevCliActionTypes this action generate types for frontend
	DevCliActionTypes DevCliAction = "types"
)

func main() {
	var _action string
	var userIDStr string
	flag.StringVar(&_action, "action", "", "Cli action to execute. Options: seed, types")
	flag.StringVar(&userIDStr, "userID", "", "UserID for seeder")
	required := []string{"action"}
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })

	for _, req := range required {
		if !seen[req] {
			fmt.Printf("missing required flag: %s\n", req)
			flag.Usage()
			os.Exit(1)
		}
	}

	action := DevCliAction(_action)
	if !slices.Contains([]DevCliAction{DevCliActionSeed, DevCliActionTypes}, action) {
		fmt.Printf("invalid action: %s\n", action)
		flag.Usage()
		os.Exit(1)
	}

	switch action {
	case DevCliActionSeed:
		RunSeeder(userIDStr)
	case DevCliActionTypes:
		RunTypeScriptify()
	}
}

func RunSeeder(strID string) {
	fmt.Print("Start seed... ")
	userID := uuid.MustParse(strID)
	db := database.InitDB("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	err := seed.Seed(db, userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Seed completed!")
}

func RunTypeScriptify() {
	fmt.Print("Generate appErrors... ")
	errorsMap := map[string]apperrors.AppError{}
	for _, e := range apperrors.AllErrors {
		errorsMap[e.Key] = e
	}

	// pretty JSON
	data, err := json.MarshalIndent(errorsMap, "", "\t")
	if err != nil {
		panic(err)
	}

	// wrap into TS code
	ts := `export interface AppError {
	id: string
	key: string
	message: string
    data?: any
}

export const ERRORS: Record<string, AppError> = ` + string(data) + "\n"

	// write to file
	if err := os.WriteFile("frontend/projects/lib/src/lib/errors.ts", []byte(ts), 0644); err != nil {
		panic(err)
	}

	fmt.Println("Generated frontend/projects/lib/src/lib/errors.ts")
}
