.PHONY: build install share connect fmt commit debug

build:
	go build -o bin/shareo ./cmd/shareo

install:
	go install ./cmd/shareo

share:
	go run ./cmd/shareo share -p 3372

connect:
	go run ./cmd/shareo connect -p 3372

fmt:
	go fmt ./...

commit:
	make fmt && git add . && git commit -F -

debug:
	go run ./cmd/shareo debug