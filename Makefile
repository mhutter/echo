SOURCES := $(wildcard *.go)

server: $(SOURCES)
	go build -race -o ./server ./cmd/server

dev:
	gin --build cmd/server

image: $(SOURCES) Dockerfile .dockerignore
	docker build -t mhutter/echo .

test:
	go test -v -race -cover ./...

clean:
	rm -f server gin-bin

.PHONY: dev image test clean
