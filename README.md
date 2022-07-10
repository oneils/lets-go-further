# lets-go-further
Repository for exercise from the book https://lets-go-further.alexedwards.net

Run the app with command line flags:

```bash
go run ./cmd/api -port=3030 -env=production
```

## Create a local Postgres DB, user and extension

```sql
create database greenlight;

create role greenlight with login password 'greenlight';

-- adds a case-insensitive character string type to PostgreSQL
create extension if not exists citext;
```

## Add posrtgres dsn to env

```bash
 GREENLIGHT_DB_DSN='postgres://greenlight:greenlight@localhost/greenlight?sslmode=disable'
```

## Database Migration tool

Install tool on Mac OS:

```bash
brew install golang-migrate
```

Install tool on Linux:

```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 $GOPATH/bin/migrate
```

Create  migration scripts

```bash
migrate create -seq -ext=.sql -dir=./migrations create_movies_table
```

- The `-seq` flag indicates that we want to use sequential numbering like 0001, 0002, ... for the migration files (instead of a Unix timestamp, which is the default).
- The `-ext` flag indicates that we want to give the migration files the extension `.sql`. 
- The `-dir` flag indicates that we want to store the migration files in the `./migrations` directory (which will be created automatically if it doesn’t already exist).
- The name `create_movies_table` is a descriptive label that we give the migration files to signify their contents.

Execute migration scripts

```bash
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```

Verify migration version:

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN version
```

Migrate up or down to a specific version:

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN goto 1
```

Run the down migration:

```bash
migrate -path=./migrations -database =$EXAMPLE_DSN down 1
```

Force the DB version to a specific version:

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN force 1
```

Send a bunch concurrent request to update a movie to verify there is not data races:

```bash
xargs -I % -P8 curl -X PATCH -d '{"runtime": "97 mins"}' "localhost:4000/v1/movies/4" < <(printf '%s\n' {1..8})
```