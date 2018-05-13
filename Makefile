clean:
	rm -rf ./build

build: clean
	go build -o build/roomo-api *.go && \
        go build -o build/parse docker/parse.go

run:
	./build/roomo-api --config config.yaml

test:
	go test -v *.go

test-getimagelist:
	go test -v -run GetImageList *.go

