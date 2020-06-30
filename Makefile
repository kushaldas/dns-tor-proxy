DEFAULT_GOAL: help
SHELL := /bin/bash


.PHONY: build
build:  ## Builds the executable
		go build github.com/kushaldas/dns-tor-proxy/cmd/dns-tor-proxy
		sudo setcap cap_net_bind_service=+ep ./dns-tor-proxy

.PHONY: clean
clean:  ## Cleans the executable
		rm ./dns-tor-proxy

.PHONY: help
help: ## Print this message and exit.
		@printf "Makefile for developing and building dns-tor-proxy\n"
		@printf "Subcommands:\n\n"
		@awk 'BEGIN {FS = ":.*?## "} /^[0-9a-zA-Z_-]+:.*?## / {printf "\033[36m%s\033[0m : %s\n", $$1, $$2}' $(MAKEFILE_LIST) \
        | sort \
        | column -s ':' -t

