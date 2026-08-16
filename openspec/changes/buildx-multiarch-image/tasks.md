# Tasks — buildx-multiarch-image (recipe proven on operator-base)

> Policy: the legacy per-arch `docker-build` machinery is painful and retired. **Comment it
> out (preserve, don't delete)** so it stays discoverable and can't silently return.
> `docker buildx` is the sole go-forward path.

- [ ] Ensure a live builder: `docker buildx ls`; on Apple-Silicon **`colima start`** first (base hit "no builder live").
- [ ] Makefile: `VERSION_V1`→`VERSION`; add `PLATFORMS ?= linux/arm64,linux/amd64`; keep `IMG ?= ${DHUBREPO}:v${VERSION}`.
- [ ] Makefile: add the buildx targets:
      ```
      docker-buildx: vet-v1
      	docker buildx build --platform ${PLATFORMS} -f build/Dockerfile -t ${IMG} -t ${DHUBREPO}:latest --push .
      docker-build: vet-v1        # single-arch --load, for local dev iteration only
      	docker buildx build --load -f build/Dockerfile -t ${IMG} .
      docker-push: docker-buildx
      ```
- [ ] **Comment out (retire, don't delete)** the legacy `docker-build-{dev,amd64,arm32v7,arm64v8}`
      targets and the `DHUBREPO_{DEV,AMD64,ARM32V7,ARM64V8}` / `IMG_{DEV,AMD64,ARM32V7,ARM64V8}`
      vars — with a `# RETIRED: superseded by docker-buildx (multi-arch)` header.
- [ ] Add the single multi-stage `build/Dockerfile` (base's recipe; binary `kubedge-ecds-operator`, entrypoint `./cmd/...`):
      ```
      # syntax=docker/dockerfile:1
      FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS builder
      ARG TARGETOS
      ARG TARGETARCH
      WORKDIR /workspace
      COPY go.mod go.sum ./
      RUN go mod download
      COPY . .
      RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
          go build -trimpath -tags=v1 -o /out/kubedge-ecds-operator ./cmd/...
      # runtime stage: FROM alpine (or distroless), COPY --from=builder /out/kubedge-ecds-operator ...
      ```
- [ ] **Comment out / retire the per-arch `build/Dockerfile.{dev,amd64,arm32v7,arm64v8}`** (leave a one-line `# RETIRED` note or move to a `legacy/` dir); the multi-stage `build/Dockerfile` is the only live one.
- [ ] **Comment out the per-arch image entries in the Helm `chart/`** (values + any templated arch-specific image refs); the chart references only `${IMG}` / `images.tags.operator`.
- [ ] `install`/`purge` → Helm v3: `install: docker-buildx` → `helm install kubedge-ecds-operator chart --set images.tags.operator=${IMG}`; `purge:` → `helm uninstall kubedge-ecds-operator`.
- [ ] Verify: `make -n docker-buildx docker-build` expands with **no live `-dev/-amd64/-arm32v7/-arm64v8` refs**; a chart render references only the single image; Go build/vet/test stay green.
- [ ] With colima up: `make docker-buildx` then `docker buildx imagetools inspect ${IMG}` → manifest list incl. linux/arm64.
