### RabbitMQ

Use queue providers so we can use in memory channel based queue and event system 
OR external driver - RabbitMQ

### Connect different types of database

Update schema definitions and make different database providers - SQLite3, Postgres,
MySQL. *(its even possible single json file instead of db, since data schema is 
not very complicated)* 🤔

### Db backup

Export database as an json file. Since we use adapters for API, we can use them for this.

### Webhooks

Register webhooks and send request.

### Pages, notes, markdown editor, documentation

### Emails and custom email templates

### Multiple boards per project

Technically it's working now but UI is not support that feature for now

### add background image like in trello (?)

Background image, color or gradient

### timer, pomodoro timer working mode

In that mode - timer will stop and start via pomodoro timer logic and send browser notifications about it's current status
Also may be OS notifications? There should be std utilities for windows and linux? 🤔

### Cli mode

full cli mode without any web related things - sockets and http routers. Should looks like LazyGit

### Calendar

Calendar for fast task/reminder assignments

### documentation

add `man-like` documentation page. Big documentation text with all commands. Should gather that information from
different packages and embed ( or not ) into binary.


### different migrations for different database

Well the Kratos approach is to use sufixes like `*.sqlite.sql` or `*.postgres.sql` for different databases.
This solve the issue with specific sql syntax and logic for different databases.