start:
	docker compose \
		-f docker/base.yml \
		-f docker/database.yml \
		-f docker/kratos.yml \
 		up -d --build

clear:
	docker compose \
		-f docker/base.yml \
		-f docker/database.yml \
		-f docker/kratos.yml \
 		down -v

start_base:
	docker compose \
		-f docker/base.yml \
		-f docker/database.yml \
 		up -d --build

types:
	go run ./dev/cli/main.go --action=types

migrate_create:
	~/go/bin/migrate create -ext sql -dir ./database/migrations -format "20060102150405" $(name)

migrate_up:
	~/go/bin/migrate \
		-path ./database/migrations/ \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		-verbose up

# usage:
# 	make migrate_down N=1 - for one migration down
migrate_down:
	~/go/bin/migrate \
		-path ./database/migrations/ \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		-verbose down $(N)

start_auth:
	go run apps/auth/cmd/main.go ./config/development.local.yml

start_project:
	go run apps/project/cmd/main.go ./config/development.local.yml

start_org:
	go run apps/org/cmd/main.go ./config/development.local.yml