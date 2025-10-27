migrate_create:
	~/go/bin/migrate create -ext sql -dir ./app/internal/db/migrations -format "20060102150405" $(name)

migrate_up:
	~/go/bin/migrate \
		-path ./app/internal/db/migrations \
		-database "sqlite3://lumna.db" \
		-verbose up

# usage:
# 	make migrate_down N=1 - for one migration down
migrate_down:
	~/go/bin/migrate \
		-path ./app/internal/db/migrations \
		-database "sqlite3://lumna.db" \
		-verbose down $(N)

build:
	cd app/frontend && \
	yarn build && \
	cd ../../ && \
	go build -o bin/lumna ./app/cmd/standalone/


clear_local:
	echo "remove local app configuration" && \
	rm -rf ~/.local/share/lumna/ && \
	rm -rf ~/.local/state/lumna/ && \
	rm ~/.local/bin/lumna && \
	rm -rf ~/.config/lumna
