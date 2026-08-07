# resources/ is embedded into the binary: build the frontend into
# resources/assets first, or the result serves no UI.
build:
	go build -o bin/lumna ./app

migrate:
	go run app/main.go --config=./config/config.yaml migrate apply

run:
	go run app/main.go --config=./config/config.yaml serve

