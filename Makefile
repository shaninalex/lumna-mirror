migrate_create:
	~/go/bin/migrate create -ext sql -dir ./app/persistance/migrations -format "20060102150405" $(name)

migrate_up:
	~/go/bin/migrate \
		-path ./app/migrations \
		-database "sqlite3://lumna.db" \
		-verbose up

# usage:
# 	make migrate_down N=1 - for one migration down
migrate_down:
	~/go/bin/migrate \
		-path ./app/migrations \
		-database "sqlite3://lumna.db" \
		-verbose down $(N)

build:
	cd app/frontend && \
	yarn build && \
	cd ../../ && \
	go build -tags embed -o bin/lumna ./app/cmd/web/


dir_clear:
	rm -rf ~/.local/share/lumna/
	rm -rf ~/.local/state/lumna/
	rm ~/.local/bin/lumna
	rm -rf ~/.config/lumna

dir_list:
	ls -la ~/.local/share/lumna/
	ls -la ~/.local/state/lumna/
	ls -la ~/.config/lumna

run:
	go run ./app

debug:
	dlv exec ./bin/lumna

clear_port:
	@pid=$$(sudo lsof -t -i :8000); \
	if [ -n "$$pid" ]; then \
		echo "Killing process on port 8000 (PID: $$pid)"; \
		sudo kill -9 $$pid; \
	else \
		echo "No process is listening on port 8000."; \
	fi

test:
	go test -tags noembed ./... -v
