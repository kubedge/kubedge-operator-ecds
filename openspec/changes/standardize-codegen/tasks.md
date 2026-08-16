# Tasks — standardize-codegen

- [x] `go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.21.0`; confirm `controller-gen --version` (→ v0.21.0).
- [x] Edit the `generate` target: remove `crd:trivialVersions=true` (adopt `crd:generateEmbeddedObjectMeta=true`, matching base).
- [x] Fix the stale `/usr/local/kubebuilder/bin` PATH: it was a misleading `echo` in the `unittest` target (not `setup`) — removed it.
- [x] `make generate`; confirm the `ecdsclusters` CRD regenerates under `chart/templates/`.
- [x] `go build ./...` green after regeneration; reviewed the CRD yaml diff.

## Consumer deviation from the seeded (base-templated) plan

- **ecds has no `pkg/apis`** — the `ECDSCluster` Go type is defined in
  `kubedge-operator-base`; ecds is a pure *consumer* of that shared API.
- So there is **no deepcopy step** (`controller-gen object`) here — base owns the
  `zz_generated.deepcopy.go`. The proposal's "deepcopy under `pkg/apis/.../v1alpha1`"
  expectation does not apply to a consumer.
- The `generate` target now regenerates the CRD from the **pinned base module**
  (`go list -m -f '{{.Dir}}' github.com/kubedge/kubedge-operator-base`), so it tracks
  whatever base version `go.mod` resolves, and filters to the single kind ecds owns
  (`ecdsclusters`; base's apis emit all 5).
- Result: the regenerated CRD is **byte-identical** to base's checked-in
  `kubedgeoperators.kubedge.cloud_ecdsclusters.yaml` (v0.21.0, adds
  `x-kubernetes-list-type: atomic` etc.), replacing the stale v0.14.0 copy.
  `make generate` is idempotent; build/vet/test stay green.
