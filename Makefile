NAME := para

.PHONY: build
build:
	go build -o bin/$(NAME)
