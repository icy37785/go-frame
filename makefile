SHELL := /bin/bash

# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 0.0.1

test:
	docker build \
    		-f deploy/docker/dockerfile.test-api \
    		-t test-api-amd64:$(VERSION) \
    		--build-arg BUILD_REF=$(VERSION) \
    		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
    		.