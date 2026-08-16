# Build a single multi-arch image with buildx (proven on operator-base)

## Why

ecds's Makefile builds four arch-suffixed images (`kubedge1/kubedge-ecds-operator-{dev,amd64,arm32v7,arm64v8}`)
by pre-compiling a binary per arch then copying it in — the same legacy shape
`kubedge-operator-base` had. Base already converted this to one multi-arch image via
`docker buildx` + a multi-stage Dockerfile (verified: `make -n` expands correctly, Go stays
green). **Copy base's pattern verbatim, adapted to ecds's names.**

## What Changes (copy base's Makefile + Dockerfile pattern)

- `VERSION_V1` → `VERSION`; keep `IMG ?= ${DHUBREPO}:v${VERSION}` (DHUBREPO = `kubedge1/kubedge-ecds-operator`).
- Add `PLATFORMS ?= linux/arm64,linux/amd64` (arm64 primary — Apple-Silicon + Pi armv8).
- Replace `docker-build-{dev,amd64,arm32v7,arm64v8}` with:
  - `docker-buildx: vet-v1` → `docker buildx build --platform ${PLATFORMS} -f build/Dockerfile -t ${IMG} -t ${DHUBREPO}:latest --push .`
  - `docker-build: vet-v1` → `docker buildx build --load -f build/Dockerfile -t ${IMG} .` (single-arch dev, `--load`)
  - `docker-push: docker-buildx` (alias)
- Delete the `DHUBREPO_{DEV,AMD64,ARM32V7,ARM64V8}` + `IMG_*` variants.
- Rewrite `build/Dockerfile` **multi-stage** (see tasks) so the binary compiles per
  `TARGETOS/TARGETARCH` — the old Dockerfile copying a prebuilt amd64 binary is **broken for arm64**.
- `install`/`purge` → Helm **v3** (`helm install/uninstall`), depending on `docker-buildx`.

## Non-goals

- The paired sim (`kubedge-sim-ecds`) image build — that's its own repo's change.

## Capabilities

### New Capabilities
- container-packaging: how the ecds operator image is built (single multi-arch).
