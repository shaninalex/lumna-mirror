Project structure principles:

`app/core` - singleton services (auth, logger, analytics), guards, interceptors, HTTP error handling, app-wide providers

`app/entities` - logical parts of application ( user, project, task...). 
Contain next parts:
- `model` - model itself and ngrx (reducer, actions, effects...)
- `api` - base CRUD operations
- `ui` - clear representative componentns, without logic ( all logic in `features` )

`app/features` - forms, actions... Everything where user can interact and change state. ( add task, change list, create project, change settings...)
It's not only forms. It's filters, searches, action buttons etc.
Features can contain multiple entities inside ( create task can contain entities/user/ui/badge if you tag someone, or comments, or permission form, or attachment form etc...)

`app/pages` - it's just pages. Simple components used in routes, have route specific logic ( parse url params for example ).
Pages can have multiple features. There will be next main parts:
- `app/pages/auth/*` - auth pages
- `app/pages/app/*` - main app pages (dashboard, settings etc)
- `app/pages/system/*` - not found, permission-denied etc.
Pages also define all routes.

`app/shared` - it's everything that used everywhere else, that logically does not part of some entity or feature ( simple ui components, generics, types )

`app/store` - NGRX store, store types etc.


Notes:
- there are no such thing as `entities/auth` ( from previous implementation ). Since we have `entities/user` - there are no sence to have separate entity for auth, since it's 
all about user.
- do not define models and interfaces that are not fully implemented ( api+ui ). I'm tolking about comments. Same thing for fields ( model files, table columns ... ).
it's just noise, that does not have any meaning.
- every component should have a folder. In other case it's just a mess. Add unit tests near component. 
- Do not span nested folders! Better think twice before defining another one. If it's using only once in parent cmp - define inside parent *.component.ts
- always use Entity approach in NGRX