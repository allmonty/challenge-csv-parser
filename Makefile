
run:
	go build
	go run main.go examples/roster1.csv examples/roster2.csv examples/roster3.csv examples/roster4.csv

test:
	go test ./test/integration_test.go
	go test ./test/... -v

.PHONY: test