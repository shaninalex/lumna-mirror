start:
	docker compose \
		-f docker/base.yaml \
		-f docker/database.yaml \
		-f docker/kratos.yaml \
 		up -d --build

clear:
	docker compose \
		-f docker/base.yaml \
		-f docker/database.yaml \
		-f docker/kratos.yaml \
 		down -v

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