start_all:
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

start:
	docker compose \
		-f docker/base.yml \
		-f docker/database.yml \
 		up -d --build

types:
	go run ./dev/cli/main.go --action=types

# usage:
# 	make seed id=8b8a6994-c474-4bf9-bc2c-a2eedcc4cb1d
seed:
	go run ./dev/cli/main.go --action=seed --userID=$(id)

migrate_create:
	~/go/bin/migrate create -ext sql -dir ./migrations -format "20060102150405" $(name)

migrate_up:
	~/go/bin/migrate \
		-path ./migrations/ \
		-database "sqlite3://flowreon.db" \
		-verbose up

# usage:
# 	make migrate_down N=1 - for one migration down
migrate_down:
	~/go/bin/migrate \
		-path ./migrations/ \
		-database "sqlite3://flowreon.db" \
		-verbose down $(N)

start_auth:
	go run cmd/cloud/auth/main.go ./config/development.local.yml

start_project:
	go run cmd/cloud/project/main.go ./config/development.local.yml

start_org:
	go run cmd/cloud/org/main.go ./config/development.local.yml

start_user:
	go run cmd/cloud/user/main.go ./config/development.local.yml

db_backup:
	docker exec -t flowreon-postgres-1 pg_dump -U postgres postgres > ./database/data/db_backup.sql

db_restore:
	# docker exec -it flowreon-postgres-1 psql -U postgres -c "DROP DATABASE postgres;"
	# docker exec -it flowreon-postgres-1 psql -U postgres -c "CREATE DATABASE postgres;"
	cat ./database/data/db_backup.sql | docker exec -i flowreon-postgres-1 psql -U postgres -d postgres

tests:
	go clean -testcache
	CONFIG_PATH=$$(pwd)/config/development.test.yml go test -v -p 1 ./...
