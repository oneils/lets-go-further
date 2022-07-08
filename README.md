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