build:
	go build -o bin/tree src/*.go

run: build
	./bin/tree