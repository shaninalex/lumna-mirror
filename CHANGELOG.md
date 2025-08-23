## 0.3.0 (2025-08-23)

### Feat

- **models**: create user code field
- **client**: load projects
- **frontend**: create library lib + move errors in it
- **api**: get project tasks
- **auth**: login with google
- **db**: replace uuid_generate_v4 with gen_random_uuid
- organization api + test
- get projects [in progress]
- **models**: add new models
- **ui**: drag and drop tasks list
- **ui**: project detail page [in progress]
- **ui**: pages
- **ui**: setup base layouts

### Fix

- **models**: fix relation between user and organization
- **db**: project key + set org id in ctx
- **api**: project api
- **test**: test db creation+testing

### Refactor

- **models**: replacing repositories with services
- **errors**: change response obj + app errors code generation
- **api**: use web.Error
- rename "services" with "apps"

## 0.2.0 (2025-08-16)

### Feat

- **auth**: basic registration
- **frontend**: setup
- **infrastructure**: prepare for frontend
- back to kratos

### Fix

- user model GORM types

## 0.1.0 (2025-08-14)

### Feat

- **db**: users table migrations + model + repository
