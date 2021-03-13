
run:
	go build
	go run main.go

test:
	go test ./test/... -v

.PHONY: test