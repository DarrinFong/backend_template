# backend_template
Learning Go with a generic Go backend for an online store.

*Under construction 👷*

### Run server
`go run main.go` To start server (using hardcoded mocks to test routes)

Access via http://localhost:8000/

Route | method | result
--- | --- | ---
`/inventory/` | GET | {"*id*": {*models.item*}...}
`/inventory/:id` | GET | {*models.item*}

### Run unit tests
`go test ./...` To run unit tests (using sqlmock for sql db)
