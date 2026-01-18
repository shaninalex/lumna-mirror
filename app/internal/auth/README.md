### internal/auth

Package structure:

```bash
internal/auth/
  service.go          // main auth service
  provider.go         // Provider interface
  middleware.go       // HTTP middleware
  errors.go
  repository.go       // DB access (credentials, accounts)

  local/
    provider.go
    password.go
    provider_test.go

  google/
    provider.go
    oauth.go

  github/
    provider.go
```
