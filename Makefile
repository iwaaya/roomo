clean:
	rm -rf ./build

build: clean
	go build -o build/roomo-api *.go

run:
	./build/roomo-api

test:
	go test -v *.go
