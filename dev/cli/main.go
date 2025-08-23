package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/database/seed"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		panic("not enough args")
	}
	userIDStr := args[1]

	userID := uuid.MustParse(userIDStr)
	db := database.InitDB("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	err := seed.Seed(db, userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Seed completed!")
}
