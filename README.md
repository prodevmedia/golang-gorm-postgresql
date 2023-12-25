## SOURCE TUTORIAL

Golang GORM RESTful API with Gin Gonic and PostgreSQL
https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres/

## Install the UUID OSSP Module for PostgreSQL

By default, PostgreSQL natively supports the UUID (Universally Unique Identifier) data type but since we are using the uuid_generate_v4() function as a default value on the ID column, we need to manually install the UUID OSSP plugin for it to work.

The uuid_generate_v4() function will be evoked to generate a UUID value for each record we insert into the database. Most PostgreSQL distributions include the UUID OSSP Module by default but do not activate it.

To install the plugin, we need to run the CREATE EXTENSION command in the Postgres shell.

Now access the bash shell of the running Postgres container with this command

> docker exec -it postgres bash

> psql -U postges golang-gorm

> CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

## Migrating

> go run cmd/migrate.go

## TUTORIAL GO

### How to Install: https://go.dev/doc/install

### PATH ON .profile or .zshrc

```
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOROOT/bin
```
