#start_all:
#	docker compose \
#		-f docker/base.yml \
#		-f docker/database.yml \
#		-f docker/kratos.yml \
# 		up -d --build
#
#clear:
#	docker compose \
#		-f docker/base.yml \
#		-f docker/database.yml \
#		-f docker/kratos.yml \
# 		down -v
#
#start:
#	docker compose \
#		-f docker/base.yml \
#		-f docker/database.yml \
# 		up -d --build

migrate_create:
	~/go/bin/migrate create -ext sql -dir ./internal/db/migrations -format "20060102150405" $(name)

migrate_up:
	~/go/bin/migrate \
		-path ./internal/db/migrations \
		-database "sqlite3://lumna.db" \
		-verbose up

# usage:
# 	make migrate_down N=1 - for one migration down
migrate_down:
	~/go/bin/migrate \
		-path ./internal/db/migrations \
		-database "sqlite3://lumna.db" \
		-verbose down $(N)

build:
	cd frontend && \
	yarn build && \
	cd ../ && \
	go build -o bin/standalone ./cmd/standalone/