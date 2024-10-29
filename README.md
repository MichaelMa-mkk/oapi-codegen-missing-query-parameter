# oapi-codegen-missing-query-parameter

## Start

```bash
go generate ./...
go run ./cmd/api
```

### Problem

- `GET http://127.0.0.1:8080/ping?option=`, which doesn't have the required query parameter, will receive a 500 response instead of 400
- `GET http://127.0.0.1:8080/ping?option=error`, which returns an error directly in strict handler, will receive a 400 response instead of 500