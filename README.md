# Prodev Media Golang REST API

## Feature

- OAUTH 2.0 Support
- Example CRUD
- Auto Reload Editing
- Migrate use cmd
- Seed use cmd

## Technology

- Golang
- Docker
- Postgres
- Air (Fast Reload Development Golang)

## Structure Folder

- app -> all script used
- app\controllers -> For Controller Logic
- app\middleware -> for middleware routing
- app\models -> for model database
- app\utils -> for utilities project
- cmd -> for go run command
- config -> for configuration project
- database\fakers -> for models faker
- database\seeders -> for seeder models
- routes -> for route api
- templates -> html format for email style or print

## Prerequisites

- [Golang](#install-go)
- [Docker](#install-docker)

#

- Familiarity with Golang, SQL, and PostgreSQL queries will be highly beneficial
- Have PostgreSQL installed. This is optional since we will be using Docker to run the Postgres server.
- VSCode as the IDE for Developing Go. I recommend VS Code because it has tools, extensions, and an integrated terminal to make your development process a breeze.

## How To Install

### 1. Running a Docker Compose for Running Postgres Server

```bash
docker compose up -d
```

### 2. [Install the UUID OSSP Module for PostgreSQL](#install-the-uuid-ossp-module-for-postgresql)

### 3. [Install Air Package](#install-air-package)

### 4. Install Depedency Golang used in this project

```bash
go get
```

### 5. Migrate and Seed

```bash
go run main.go migrate
go run main.go seed
```

### 6. [Running Golang Server](#how-to-run)

## How To Run

### 1. Run Postgres from Docker

```bash
compose docker up -d
```

### 2. Run Golang Server

```bash
air
```

### Postman Collection

https://documenter.getpostman.com/view/29378887/2s9Ykt5yuV

## Install the UUID OSSP Module for PostgreSQL

By default, PostgreSQL natively supports the UUID (Universally Unique Identifier) data type but since we are using the uuid_generate_v4() function as a default value on the ID column, we need to manually install the UUID OSSP plugin for it to work.

The uuid_generate_v4() function will be evoked to generate a UUID value for each record we insert into the database. Most PostgreSQL distributions include the UUID OSSP Module by default but do not activate it.

To install the plugin, we need to run the CREATE EXTENSION command in the Postgres shell.

Now access the bash shell of the running Postgres container with this command

```bash
docker compose up -d
```

```bash
docker exec -it postgres bash
```

```bash
psql -U postgres golang-gorm
```

```bash
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

List all the available extensions

```bash
select * from pg_available_extensions;
```

## Migrating

```bash
go run cmd/migrate.go
```

## Seed

```bash
go run cmd/seed.go
```

## Install Air Package

### Via go install (Recommended)

```bash
go install github.com/cosmtrek/air@latest
```

### Via install.sh

```bash
# binary will be $(go env GOPATH)/bin/air

curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/

curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s


air -v
```

### Via goblin.run

```bash
# binary will be /usr/local/bin/air
curl -sSfL https://goblin.run/github.com/cosmtrek/air | sh

# to put to a custom path
curl -sSfL https://goblin.run/github.com/cosmtrek/air | PREFIX=/tmp sh
```

or check

https://github.com/cosmtrek/air

## Install Go

### How to Install: https://go.dev/doc/install

### PATH ON .profile or .zshrc

```bash
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOBIN
```

# Install Docker

https://docs.docker.com/desktop/

## SOURCE TUTORIAL

Golang GORM RESTful API with Gin Gonic and PostgreSQL
https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres/
