# Docker GoReleaser Sample

This is a sample Application, that shows how to create a small Go Application wrapped inside a Docker container by GoReleaser (https://goreleaser.com/).

* Features:
    * Directories inside container
    * Metrics in an exposed webserver with expvar package (https://pkg.go.dev/expvar)
    * Dockerfile with Delve and vscode launch settings.

0. Get GoReleaser:

    https://github.com/goreleaser/goreleaser

1. Let Goreleaser build your docker Go App:

    `./release.sh`

2. Run container:

    `./run_container.sh`
