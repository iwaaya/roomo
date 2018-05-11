clean:
	rm -rf ./build

build: clean
	go build -o build/roomo-api *.go

run:
	./build/roomo-api --config config.yaml

test:
	go test -v *.go
