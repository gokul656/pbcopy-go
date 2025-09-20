CGO_ENABLED=1

run:
	go run cmd/main.go

install-linux:
	sudo cp bin/pbcopy-linux /usr/local/bin/pbcopy

build:
	GOOS=darwin go build -o bin/pbcopy-darwin cmd/main.go
	GOOS=linux go build -o bin/pbcopy-linux cmd/main.go
	@make install-linux

.PHONY: run build install-linux
