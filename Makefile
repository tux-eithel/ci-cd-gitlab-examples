.PONY: help rsync build-amd build-arm

.DEFAULT_GOAL := help
SHELL := /bin/bash

destination = 192.168.1.5:~


help:
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

build-amd: ## Build the hello app for AMD
	GOOS=linux GOARCH=amd64 go build -o hello main.go

build-arm: ## Build the hello app for ARM
	GOOS=linux GOARCH=arm64 go build -o hello main.go

rsync: ## Rsync hello exec to remote server
	@printf "Rsync file to remote server"
	rsync hello $(destination)


