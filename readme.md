# Prerequisite
The following tools should be installed and added to `$PATH`

- [Mirate](https://github.com/golang-migrate/migrate)
- [PostgREST](https://docs.postgrest.org/en/v12/)

# Usage

```
NAME:
   gpm - CLI to make the workflow smoother when working with PostgREST

USAGE:
   gpm [global options] command [command options]

COMMANDS:
   migrate-up       Run all `.up.sql` migration files in the migration directory
   migrate-down     Run all `.down.sql` migration files in the migration directory. Effectively revert all the changes from the `up` files
   migrate-to       Migrate to a specific version. Revert all changes from the migration after the specified version
   make-migrations  Create new migration files
   check-db         Check the database config
   migrate-force    Force the migration to a specific version when it's dirty
   run-pg-rest      Run PostgREST API server
   help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```
