# RETIRED Dockerfiles — superseded by `docker buildx`

These per-arch Dockerfiles are **retired**, kept only for discoverability. Do not
resurrect them — the single multi-stage `build/Dockerfile` + `make docker-buildx`
(multi-arch manifest list) is the sole go-forward image build path.

- `Dockerfile.{dev,amd64,arm32v7,arm64v8}` — `FROM scratch`, each `ADD`ed a binary that
  the Makefile pre-compiled per arch. The arm images actually shipped an amd64 binary
  (the old `docker-build-arm*` targets built for the wrong arch in places), which is
  exactly the bug the buildx cross-compile (`GOARCH=${TARGETARCH}`) fixes.
- `Dockerfile.buildkit` — an earlier multi-arch attempt (golang:1.23, `cmd/manager/main.go`).
  Replaced by `build/Dockerfile` (golang:1.26-alpine, `./cmd/...`, `-tags=v1 -trimpath`).

See the retired Makefile targets (commented, marked `# RETIRED`) for how these were invoked.
