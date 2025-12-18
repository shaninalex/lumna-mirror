
Project structure principles:

`app/entities` - logical parts of application ( user, project, task...). 
Contain next parts:
- `model` - model itself and ngrx (reducer, actions, effects...)
- `api` - base CRUD operations
- `ui` - clear representative componentns, without logic ( all logic in `features` )

`app/features` - forms, actions... Everything where user can interact and change state. ( add task, change list, create project, change settings...)
It's not only forms. It's filters, searches, action buttons etc.
Features can contain multiple entities inside ( create task can contain entities/user/ui/badge if you tag someone, or comments, or permission form, or attachment form etc...)

`app/pages` - it's just pages. Simple components used in routes, have route specific logic ( parse url params for example ).
Pages can have multiple features

`app/shared` - it's everything that used everywhere else, that logically does not part of some entity or feature ( simple ui components, generics, types )

`app/store` - NGRX store, store types etc.