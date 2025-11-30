## setup package

This is initialization scenario. On startup application checks users information:

- does database configuration correctly and we can access it?
- does user create admin account?
- does all working directories exists?
- does port which will be used as an server running port is not occupied?

CLI mode does not run setup. You will be able to set this information manually via cli comands.

But TUI mode (terminal user interface) - do use this setup script, since it's basically the same as a web, bun in terminal.
