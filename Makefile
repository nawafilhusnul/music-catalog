mock:
	go generate -v ./...
run:
	go run cmd/main.go
test:
	go test -v ./...
test-cover:
	go test -v -cover ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
