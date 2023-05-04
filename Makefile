.DEFAULT_GOAL := help
SHELL         := /bin/bash
RED     := \033[31m
CYAN    := \033[36m
MAGENTA := \033[35m
RESET   := \033[0m

.PHONY:	setup	## Setup .envrc, etc.
setup:
	cp -i .envrc.sample .envrc
	@echo "$(CYAN) Please edit .envrc"


.PHONY:	app/run	## Run app with automatically building
app/run:
	go run .


.PHONY:	lint	## Lint files
lint:
	golangci-lint run ./...

.PHONY:	fmt	## Reformat files (overwrite)
fmt:
	go fmt ./...


.PHONY:	help	## Show Makefile tasks
help:
	@grep -E '^.PHONY:\s*\S+\s+#' Makefile | \
		sed -E 's/.PHONY:\t*//' | \
		awk 'BEGIN {FS = "(\\t*##\\s*)?"}; {printf "$(CYAN)%-22s$(RESET) %s\n", $$1, $$2}'
