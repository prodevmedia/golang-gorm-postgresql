# Prodev Media Golang REST API

- app -> all script used
- app\controllers -> For Controller Logic
- app\middleware -> for middleware routing
- app\models -> for model database
- app\utils -> for utilities project
- cmd -> for go run command
- config -> for configuration project
- database -> for database script
- routes -> for route api
- templates -> html format for email style or print

## Technology

- Golang
- Docker
- Postgres
- Air (Fast Reload Development Golang)

## Prerequisites

- Golang
- Docker

#

- Familiarity with Golang, SQL, and PostgreSQL queries will be highly beneficial
- Have PostgreSQL installed. This is optional since we will be using Docker to run the Postgres server.
- VSCode as the IDE for Developing Go. I recommend VS Code because it has tools, extensions, and an integrated terminal to make your development process a breeze.

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

## Install Air Package

> go install github.com/cosmtrek/air@latest

Or check other

https://github.com/cosmtrek/air

## How To Run

> go get

> go run cmd/migrate.go

> air

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

## SOURCE TUTORIAL

Golang GORM RESTful API with Gin Gonic and PostgreSQL
https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres/
