build:
	go build -o bin/lumna ./app

build_embed:
	go build -tags embed -o bin/lumna_embed ./app

migrate:
	go run app/main.go --config=./config/config.yaml migrate apply

run:
	go run app/main.go --config=./config/config.yaml serve

