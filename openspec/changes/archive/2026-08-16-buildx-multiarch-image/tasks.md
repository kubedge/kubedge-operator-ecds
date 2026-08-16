# Tasks — buildx-multiarch-image (recipe proven on operator-base)

> Policy: the legacy per-arch `docker-build` machinery is painful and retired. **Comment it
> out (preserve, don't delete)** so it stays discoverable and can't silently return.
> `docker buildx` is the sole go-forward path.

- [x] Ensure a live builder: `docker buildx ls`; on Apple-Silicon **`colima start`** first (base hit "no builder live"). → colima up, builder `colima` running (v0.30.0).
- [x] Makefile: `VERSION_V1`→`VERSION`; add `PLATFORMS ?= linux/arm64,linux/amd64`; keep `IMG ?= ${DHUBREPO}:v${VERSION}`.
- [x] Makefile: add the buildx targets `docker-buildx` (multi-arch `--push`), `docker-build` (single-arch `--load`), `docker-push` (alias → docker-buildx).
- [x] **Comment out (retire, don't delete)** the legacy `docker-build-{dev,amd64,arm32v7,arm64v8}` targets and `DHUBREPO_*`/`IMG_*` vars, marked `# RETIRED: superseded by docker-buildx (multi-arch)`.
- [x] Add the single multi-stage `build/Dockerfile` (base's recipe; binary `kubedge-ecds-operator`, entrypoint `./cmd/...`). **Runtime = `FROM scratch`** (ecds's proven contract), copying CA certs + `build/ecds-templates` → `/opt/kubedge-operators/ecds-templates`; runs the binary directly (no user_setup/entrypoint chain, unlike base's alpine image).
- [x] Retire the per-arch `build/Dockerfile.{dev,amd64,arm32v7,arm64v8}` + `Dockerfile.buildkit` → moved to `build/legacy/` with a `README.md` explaining the retirement.
- [x] Per-arch image entries in the Helm `chart/`: **none existed** — `values.yaml` already had a single `images.tags.operator`. Fixed a related bug instead: `operator.yaml` hardcoded `kubedge1/kubedge-ecds-operator:v0.2.0`; now `{{ .Values.images.tags.operator }}` + `{{ .Values.images.pull_policy }}` so `install --set` works.
- [x] `install`/`purge` → Helm v3: `install: docker-buildx` → `helm install … --set images.tags.operator=${IMG}`; `purge:` → `helm uninstall …`.
- [x] Verify: `make -n docker-buildx docker-build install purge` expands with no live `-dev/-amd64/-arm32v7/-arm64v8` target/image refs; Go build/vet/test stay green.
- [x] With colima up: multi-arch build proven. `make docker-build` (single-arch `--load`) builds+loads an arm64 image. A direct `docker buildx build --platform linux/arm64,linux/amd64 … --output type=oci` produces a **manifest list containing linux/arm64 + linux/amd64** (verified via the OCI index). `make docker-buildx` itself (`--push`) is unrun here — needs docker.io creds; the OCI build exercises the identical cross-compile path.

## Deviations / follow-ups surfaced for step 4 (DEPLOY)

- **`helm` is not installed** on this machine — `make install`/`purge` (Helm v3) cannot run until it is (`brew install helm`). Blocks step 4.
- **Operator scheduling:** `chart/templates/operator.yaml` pins `nodeSelector: czZone: enabled`, and the example CRs pin `ezZone/czZone` nodeSelectors on the component pods. On a plain colima/kind node these won't schedule — step 4 will need node labels or nodeSelector removal.
- **Registry push unverified:** `make docker-buildx` (`--push`) needs `docker login`; only the local OCI multi-arch build was exercised.
