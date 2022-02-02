SHELL := /bin/bash

service:
	docker build \
		-f Dockerfile \
		-t service:1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.
