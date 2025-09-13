## 0.8.0 (2025-09-14)

### Feat

- user settings form + settings api

## 0.7.0 (2025-09-13)

### Feat

- task board view, create form, update status
- **ui**: projects list + simple create form
- create org on user registration 2
- create org on user registration
- app errors effects + join or create org page
- dispatch app errors

### Fix

- **ui**: project detail route + template
- **ui**: auth flows
- tests
- tests
- user settings
- create org on user registration

## 0.6.0 (2025-09-09)

### Feat

- **ui**: add account settings page
- **ui**: add auth pages
- remove Kratos forms api
- user service domain + uncompleted tests
- **api**: project task list return tasks only
- **ui**: task entity + redux
- **ui**: project card + project state + project resolver change
- **ui**: minimalistic primary layout and components
- **ui**: registration, login, verification and main component
- **ui**: simplified login form
- **auth**: registration hook middleware

### Fix

- task redux

### Refactor

- **ui**: bunch of small little things
- **ui**: refactoring preloader auth layout
- **ui**: replace features/auth with page data resolvers

## 0.5.0 (2025-08-30)

### Feat

- **ui**: branding

### Fix

- **tests**: test db and tdata.Manager

### Refactor

- split controllers + define managers interfaces
- use similar structure for services
- service structure

## 0.4.0 (2025-08-29)

### Feat

- rename app to Flowreon
- update task
- **ui**: task detail form modal dialog
- **ui**: task detail form modal dialog
- **model**: add task code
- save status and sort
- **models**: remove issue type and add completed prop
- **ui**: develop board [in progress]
- get project board

### Fix

- small issues and improvement
- **ui**: tasks drag and drop list

### Refactor

- **api**: project api handlers
- **ui**: replace the word "issue" with "task"
- rename issue to task
- **ui**: board view
- **models**: define interfaces
- **db**: introduce models interfaces + reorganize code

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
