# auth-api

run test: `go test -tags testing ./..`

run test coverage: `go test ./.. -coverprofile=cover.out`

open coverage: `go tool cover -html=cover.out`

run app: `go run ./cmd/main.go`

dependency download: `go mod tidy`