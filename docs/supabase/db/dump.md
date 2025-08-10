## tealbase-db-dump

Dumps contents from a remote database.

Requires your local project to be linked to a remote database by running `tealbase link`. For self-hosted databases, you can pass in the connection parameters using `--db-url` flag.

Runs `pg_dump` in a container with additional flags to exclude Tealbase managed schemas. The ignored schemas include auth, storage, and those created by extensions.

The default dump does not contain any data or custom roles. To dump those contents explicitly, specify either the `--data-only` and `--role-only` flag.
