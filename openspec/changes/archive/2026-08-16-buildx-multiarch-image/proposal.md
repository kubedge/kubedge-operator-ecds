# Build a single multi-arch image with buildx (proven on operator-base)

## Why

ecds's Makefile builds four arch-suffixed images (`kubedge1/kubedge-ecds-operator-{dev,amd64,arm32v7,arm64v8}`)
by pre-compiling a binary per arch then copying it in — the same legacy shape
`kubedge-operator-base` had. Base already converted this to one multi-arch image via
`docker buildx` + a multi-stage Dockerfile (verified: `make -n` expands correctly, Go stays
green). **Copy base's pattern verbatim, adapted to ecds's names.**

## What Changes (copy base's Makefile + Dockerfile pattern)

**Policy: `docker buildx` is the sole go-forward path. The legacy per-arch `docker-build`
machinery was painful to maintain — comment it out (retire, preserve), do NOT delete it, so
it stays discoverable and can't silently return.**

- `VERSION_V1` → `VERSION`; keep `IMG ?= ${DHUBREPO}:v${VERSION}` (DHUBREPO = `kubedge1/kubedge-ecds-operator`).
- Add `PLATFORMS ?= linux/arm64,linux/amd64` (arm64 primary — Apple-Silicon + Pi armv8).
- Add the buildx targets: `docker-buildx` (multi-arch `--push`), a `--load` single-arch
  `docker-build` for local dev only, `docker-push` alias.
- **Comment out (retire):** the `docker-build-{dev,amd64,arm32v7,arm64v8}` targets, the
  `DHUBREPO_{DEV,AMD64,ARM32V7,ARM64V8}` + `IMG_*` variants, the per-arch
  `build/Dockerfile.{dev,amd64,arm32v7,arm64v8}` files, and the per-arch image entries in the
  Helm `chart/`. Mark each with a `# RETIRED: superseded by docker-buildx` note.
- Add a single multi-stage `build/Dockerfile` (see tasks) that compiles per
  `TARGETOS/TARGETARCH` — the old per-arch Dockerfiles copying a prebuilt amd64 binary are
  **broken for arm64**.
- `install`/`purge` → Helm **v3** (`helm install/uninstall`), depending on `docker-buildx`.

## Non-goals

- The paired sim (`kubedge-sim-ecds`) image build — that's its own repo's change.

## Capabilities

### New Capabilities
- container-packaging: how the ecds operator image is built (single multi-arch).
