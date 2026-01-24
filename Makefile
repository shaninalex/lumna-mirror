build:
	go build -ldflags "-X lumna.Version=$(git describe --tags --dirty)" -o bin/lumna ./app

build_embed:
	go build -tags embed -o bin/lumna_embed ./app
