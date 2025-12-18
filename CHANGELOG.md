## 0.21.0 (2025-12-19)

### Feat

- **db**: expresions and conditions
- **test**: board+board_list
- **auth**: authentication + tests
- **api**: user handlers
- **api**: auth controller
- embed, db
- **build**: static handlers and embed
- **web**: router and test for that
- **repository**: user

### Fix

- restore authentication

### Refactor

- use typed options
- domain repositories

## 0.20.1 (2025-11-15)

### Fix

- embed dir path
- **ui**: project overview component layout
- **ui**: dark mode inputs
- **config**: fix configuration issues

### Refactor

- rename internal to pkg
- struct definition duplications
- fix duplications - badges and comments

## 0.20.0 (2025-11-11)

### Feat

- comment separate entity and api resource
- **ui**: comment form (wrong)
- **ui**: comment list
- comment manager
- add comments migrations and models

### Fix

- **ui**: set task title input as focus on form open

### Refactor

- **ui**: fix comments creations and list, place components in a proper package

## 0.19.1 (2025-11-01)

### Fix

- **ci**: migrate github actions to gitlab pipelines pipeline #53 %2
- **ci**: migrate github actions #53

## 0.19.0 (2025-11-01)

### Feat

- move to gitlab #51
- do not use absolute path #52
- add port number to configuration file

### Fix

- **ci**: build tags for embed directory
- **ui**: minor ui fixes

### Refactor

- **auth**: provide UserService

## 0.18.0 (2025-10-25)

### Feat

- **ci**: config tests github action workflow #10
- **ci**: config tests github action workflow #10
- on start pipeline #12
- **ui**: move delete column button into the settings page #23
- **config**: viper, switch to config file instead of env #15
- create working directories #5
- create working directories #5

### Fix

- **ui**: project settings feature styles size
- **ui**: task-detail layout
- remove status completed field #22
- **ui**: task-detail width
- **ui**: hovers on btn's #26
- init config directory #27
- **ui**: card styles
- **ui**: board view columns and scrollbar
- **ui**: colors and overview
- **ui**: create card form
- **routes**: print routes only in dev mode #6

## 0.17.0 (2025-10-20)

### Feat

- **ui**: changed task detail page to board modal
- **ui**: ngx-editor
- **ui**: response toast manager
- delete task button
- move task dropdown component

### Fix

- **ui**: drag&drop on new items
- add view mode for project

### Refactor

- component prefixes
- move task detail logic to it's feature component

## 0.16.0 (2025-10-10)

### Feat

- **ui**: light/dark theme
- patch task detail description
- task detail description
- task detail page

### Refactor

- component sufixes

## 0.15.0 (2025-10-05)

### Feat

- **ui**: sort statuses + breadcrumbs
- **ui**: ui for project settings page

### Fix

- **ui**: styles
- **ui**: styles
- **ui**: create task form position
- **ui**: icons + sidebar collapse
- **ui**: styles

## 0.14.0 (2025-10-05)

### Feat

- remove material design again ( use only angular cdk )

### Fix

- complete tasks via status
- **ui**: task list index
- **ui**: task list index
- **ui**: change status

## 0.13.0 (2025-10-05)

### Feat

- **ui**: project statuses
- **api**: project statuses

### Fix

- **ui**: reset form status form after submit + properly select statuses
- **db**: drop constraint unique status name

### Refactor

- **ui**: use material design and angular cdk
- move user service to root domain

## 0.12.0 (2025-10-04)

### Feat

- task services
- **api**: task handler placeholders
- **api**: badge handler
- **db**: badge table + type + repo
- task handler
- **api**: user token (list/delete/revoke)
- **ui**: access/refresh JWT token
- **api**: access/refresh JWT token
- jwt token utility package
- **api**: headers middleware
- **db**: company table and relations

### Fix

- **db**: remove company table
- **api**: logout
- **api**: api endpoints
- auth with updated user tokens services and table
- **ui**: drag&drop board view

### Refactor

- root domain
- root domain
- project service

## 0.11.0 (2025-09-24)

### Feat

- **ui**: theme switcher icon + theme
- **ui**: theme switcher

### Fix

- **api**: use http-only cookies instead of localStorage
- **ui**: select themes

## 0.10.0 (2025-09-22)

### Feat

- **db**: remove GORM
- **db**: add tasks table
- include project and fixed it to start at least
- jwt based authentication
- login/logout/register pages + functionality
- **api**: replace fiber with std http router

### Fix

- **api**: task create
- **api**: additional fixes and manual testings
- jwt validation
- **api**: change task status

## 0.9.1 (2025-09-14)

### Fix

- **ui**: version on auth layout

## 0.9.0 (2025-09-14)

### Feat

- **CI**: build images

### Fix

- **CI**: service /_health endpoint

## 0.8.1 (2025-09-14)

### Fix

- remove redundant types

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
