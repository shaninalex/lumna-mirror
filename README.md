![Lumna](assets/logo-h.svg "Lumna")

# Lumna

Lumna - is a self hosted, all in one, project management system that can be
hosted in your own vps with minimal setup and configuration.

### build

```bash
make build
make build_embed
```

### development

```bash
go mod tidy
cp ./config/config.test.yaml ./config/config.yaml

# modify path to database

cd frontend
yarn build

go run ./app migrate apply --config=config/config.yaml
go run ./app serve  --config=config/config.yaml

# go to http://localhost:8000/setup and setup your workspace
```
