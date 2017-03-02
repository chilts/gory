all:
	echo 'Provide a target: gory clean'

fmt:
	find . -name '*.go' -exec go fmt {} ';'

run: fmt
	go run gory.go

build: fmt
	go build gory.go

gory: build
	./gory

test:
	go test

clean:
	rm -rf gory

.PHONY: gory
