# Tasks — container-run-smoke

> Consumer note: the proposal is base-templated ("base is a library, no standalone
> image"). ecds IS a standalone operator with its own image, so the smoke gate is
> direct: build ecds's image → deploy → reconcile a sample CR → teardown.

- [x] DECISION: local cluster for K8S01 = **kind on colima's docker**. Rationale: colima already provides the docker daemon buildx needs, and `kind load docker-image` moves the image straight into the node's containerd (k3s/containerd can't see colima-docker images otherwise). minikube / colima+k3s rejected for extra image-plumbing.
- [x] `make generate` → CRD applied to the cluster; confirmed it registers. → CRD ships in `chart/crds/`; `helm install` registers it (proven step 4 + smoke).
- [x] Build the operator image (buildx arm64) and deploy it. → `bin/smoke.sh deploy`: `docker buildx build --load` + `kind load docker-image` + `helm upgrade --install … --set images.pull_policy=Never`.
- [x] Apply a sample CR; confirm the operator pod starts and status reports satisfied. → asserts `.status.satisfied == true` (polled) and `kubectl rollout status` on all 5 component StatefulSets.
- [x] `make undeploy`; confirm clean teardown. → `bin/smoke.sh down` / `make smoke-down`: delete CR (finalizer GCs children) + `helm uninstall`.

## Formalized as a repeatable gate

- **`bin/smoke.sh`** (`up` / `down` / `nuke` / `cycle`) — idempotent; reuses the cluster
  if present; shellcheck + shfmt clean (CI's `lint` job scans `bin/*.sh`).
- **`make smoke` / `smoke-up` / `smoke-down`** — thin wrappers.
- **Verified live:** `make smoke` (full cycle) passes in ~30s — build → load → deploy →
  apply CR → `SATISFIED=true` + 5/5 StatefulSets Ready → teardown, exit 0.
- Documented the `make install` caveat (it pushes; use `make smoke` locally).

## Not automated in GitHub Actions (by design)

- The smoke gate needs docker+kind+a cluster; GitHub-hosted runners don't have colima and
  spinning kind per-PR is heavy. This stays a **local pre-merge gate** (run `make smoke`);
  CI keeps the fast `go build/vet/test/lint` jobs. Revisit if a self-hosted runner appears.
