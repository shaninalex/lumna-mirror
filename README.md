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
go run ./app migrate --config=config/config.yaml  # ./bin/migrate.sh
go run ./app import resources/dev_db.json --config=config/config.yaml # ./bin/import.sh
go run ./app serve  --config=config/config.yaml # ./bin/dev.sh

# login with test@test.com:111 in http://localhost:8000/auth/login
```
