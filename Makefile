build:
	go build -o ./bin/bytex

run: build
	./bin/bytex

test:
	go test ./...