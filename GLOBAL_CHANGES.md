# Global Changes

Too many things I do not like in my application:

### Database

I thought not to use GORM because was trying to use std lib as much as possible. But
that can make too much code to maintain, subtle security and performance issues, longer
issue fixes and features delivery. Not sure about migrations. It's good for big project
with tons of data... But for this type application? Not sure. I'll have 15-20 tables...

With orm we do not need repositories any more since amount of code is much decreased.
Need only services.

### Web router

I make my own simple router, but it can be a little too grounded - no validations,
security, speed optimization and plugin ecosystem like in Gin.

### Auth

Now it's full of bugs, hardly maintainable mess of issues. I need something much safer,
simpler, need to easy implement new ways of authorization ( login with Google/Github )

### Services/Repositories

I think I'm doing it wrong. There are some architectural error I can't spot.

### Config

Store too little configuration variables. It should be single source of ANY configuration.

### UUID for all ids

I do not like number based id's

# Any good things so far?

### cmd

I like how it's done. Completely separate thing that init application and run different
commands

### Frontend

It's okay, in my opinion. Balance between FSD approach and my own way of organizing files.
Signals is good, kanban board is good, ui is stupidly simple ( no Material design or
any other ui frameworks. Simple tailwind. )
