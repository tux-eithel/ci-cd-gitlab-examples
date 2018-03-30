.PONY: help rsync build

.DEFAULT_GOAL := help
SHELL := /bin/bash

destination = 192.168.1.5:~

help:
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

build: ## Build the hello app
	go build -o hello main.go

rsync: ## Rsync hello exec to remote server
	rsync hello $(destination)


