#!/usr/bin/env zsh

#!/usr/bin/bash

CGO_ENABLED=0 GOOS=linux go build -v -gcflags "all=-N -l" -o ./go_sample cmd/helloworld.go

docker buildx build --no-cache --progress=plain --tag goreleaser_sample --file Dockerfile.goreleaser_sample .
