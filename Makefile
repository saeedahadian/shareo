.PHONY: build install share connect format commit

build:
	go build -o bin/shareo ./cmd/shareo

install:
	go install ./cmd/shareo

share:
	go run ./cmd/shareo share -p 3372

connect:
	go run ./cmd/shareo connect -p 3372

format:
	go fmt ./...

commit:
	make format && git add . && git commit -F -