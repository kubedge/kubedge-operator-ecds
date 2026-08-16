# container-packaging Specification

## Purpose
TBD - created by archiving change buildx-multiarch-image. Update Purpose after archive.
## Requirements
### Requirement: The operator image is a single multi-arch image via `docker buildx`

The operator SHALL be packaged as one multi-arch image `${DHUBREPO}:v${VERSION}` covering
at least `linux/arm64` (plus `linux/amd64` for x86 clusters) via `docker buildx build
--platform ${PLATFORMS} --push`, built from a single multi-stage `build/Dockerfile` that
compiles the binary per `TARGETOS/TARGETARCH`. `docker buildx` is the sole go-forward build
path.

#### Scenario: one tag serves all target arches
- **WHEN** `docker buildx imagetools inspect ${DHUBREPO}:v${VERSION}` runs after `make docker-buildx`
- **THEN** the tag resolves to a manifest list including `linux/arm64`

### Requirement: Legacy per-arch build paths are commented out, not maintained

The legacy build machinery SHALL be **commented out (retired-but-preserved), not deleted** —
so the history stays visible and nothing silently returns. This covers: the per-arch
`docker-build-{dev,amd64,arm32v7,arm64v8}` targets and their `DHUBREPO_*`/`IMG_*` variables
in the Makefile; the per-arch `build/Dockerfile.{dev,amd64,arm32v7,arm64v8}` files; and the
per-arch image entries in the Helm chart. `docker buildx` (multi-arch) replaces all of them;
a single `--load` dev build MAY remain for local iteration.

#### Scenario: legacy build entries are disabled but discoverable
- **WHEN** the Makefile, per-arch Dockerfiles, and chart are inspected after this change
- **THEN** the legacy per-arch build targets / Dockerfiles / chart image entries are present only as comments (or clearly-marked retired blocks), and the live build path is `docker buildx` only

#### Scenario: no live reference to the legacy paths
- **WHEN** `make -n docker-buildx` and a chart render are run
- **THEN** they reference only the single multi-arch image and never a `-dev/-amd64/-arm32v7/-arm64v8` image or a per-arch Dockerfile

