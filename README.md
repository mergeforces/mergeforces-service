# Mergeforces Service

[![Build Status](https://travis-ci.org/mergeforces/mergeforces-service.svg?branch=master)](https://travis-ci.org/mergeforces/mergeforces-service.svg?branch=master)[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/mergeforces/mergeforces-service/blob/master/.github/LICENSE.md)


Mergeforces service is a hands-on example of a complete web application written in Go for the purposes of applying knowledge gained from self learning and creating an open source project that provides free user group listings.

### Quickstart

To build run the web service on `http://localhost:8080`:

```bash
docker-compose up --build
```

Once the database has started, you will need to use `psql` or a client like `postico` to connect to the running postgres instance and create the `mergeforces_app` database:

```bash
psql postgres://postgres:postgres@0.0.0.0:5432
CREATE DATABASE mergeforces_app;
```

### Running tests

To run all tests within the repository:

```go
go test ./...
```

### Running migrations

Migrations will be run automatically upon restarting the application; however to run migraitons manually:

```bash
goose postgres "user=postgres dbname=mergeforces_app sslmode=disable" status
goose postgres "postgres://postgres:postgres@0.0.0.0:5432/mergeforces_app?sslmode=disable" up
```

### Running linting

Install `golangci-lint` via `brew install golangci/tap/golangci-lint` and run linting:

```bash
golangci-lint run
```
