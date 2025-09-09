## Frontend structure

Trying to follow FSD principles.

`/app` - where everything is bootstrapped

`/app/layouts` - different layout for pages ( auth-layout, main-layout, profile-layout etc )

`/entities` - where entities described

`/entities/project/api` - CRUD api calls, not related to any features

`/entities/project/model` - actual model + NGRX things

`/entities/project/ui` - representational only components

`/environments` - environments

`/features` - features, business logic related interfaces that DO something with the entities

`/features/project/create-form` - create form feature.

`/features/project/create-form/api` - api service for create api endpoints

`/features/project/create-form/model` - feature specific models + NGRX things

`/features/project/create-form/ui` - interface of the form. May or not may use entities/ui components

`/pages/project` - static pages that that may contain features or entities ui elements

`/shared` - just a bunch of project specific shared code. Note, most of the things can be moved in the `lib` library and used somewhere else, in admin panel for example