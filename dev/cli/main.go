// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/database/seed"
)

// DevCliAction - dev cli action.
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

// RunSeeder - runs the seeder.
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

// RunTypeScriptify - runs the type scriptify.
func RunTypeScriptify() {
	fmt.Print("Generate appErrors... ")
	resultPath := "frontend/projects/client/src/shared/common/errors.ts"
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
	if err := os.WriteFile(resultPath, []byte(ts), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", resultPath)
}
