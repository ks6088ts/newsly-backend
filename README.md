![test](https://github.com/ks6088ts/newsly-backend/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ks6088ts/newsly-backend)](https://goreportcard.com/report/github.com/ks6088ts/newsly-backend)
[![GoDoc](https://godoc.org/github.com/ks6088ts/newsly-backend?status.svg)](https://godoc.org/github.com/ks6088ts/newsly-backend)

# newsly-backend

## Environment variables

Specify following environment variables for configurations

- DB_TYPE = `mock` or `postgresql`
- DB_DATA_SOURCE_NAME = `"host=db port=5432 user=user password=password dbname=db sslmode=disable"`
- ALLOWED_ORIGIN = `http://localhost:3000`


```bash
# help
./newsly-backend --help

# Start GraphQL server
./newsly-backend server
```
