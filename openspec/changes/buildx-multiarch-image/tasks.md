# Tasks — buildx-multiarch-image (recipe proven on operator-base)

- [ ] Ensure a live builder: `docker buildx ls`; on Apple-Silicon **`colima start`** first (base hit "no builder live" here — buildx CLI present but daemon down).
- [ ] Makefile: `VERSION_V1`→`VERSION`; add `PLATFORMS ?= linux/arm64,linux/amd64`; keep `IMG ?= ${DHUBREPO}:v${VERSION}`.
- [ ] Makefile: replace the four `docker-build-{dev,amd64,arm32v7,arm64v8}` targets with:
      ```
      docker-buildx: vet-v1
      	docker buildx build --platform ${PLATFORMS} -f build/Dockerfile -t ${IMG} -t ${DHUBREPO}:latest --push .
      docker-build: vet-v1
      	docker buildx build --load -f build/Dockerfile -t ${IMG} .
      docker-push: docker-buildx
      ```
- [ ] Delete `DHUBREPO_{DEV,AMD64,ARM32V7,ARM64V8}` and `IMG_{DEV,AMD64,ARM32V7,ARM64V8}`.
- [ ] Rewrite `build/Dockerfile` multi-stage (copy base's, rename binary to `kubedge-ecds-operator`, entrypoint `./cmd/...` = cmd/manager):
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
- [ ] `install`/`purge` → Helm v3: `install: docker-buildx` → `helm install kubedge-ecds-operator chart --set images.tags.operator=${IMG}`; `purge:` → `helm uninstall kubedge-ecds-operator`.
- [ ] Verify: `make -n docker-buildx docker-build` expands, no residual `-v1` image refs, Go build/vet/test stay green (Makefile/Dockerfile don't touch Go).
- [ ] With colima up: `make docker-buildx` then `docker buildx imagetools inspect ${IMG}` → manifest list incl. linux/arm64.
