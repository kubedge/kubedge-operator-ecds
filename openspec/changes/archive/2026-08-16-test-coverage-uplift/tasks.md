# Tasks — test-coverage-uplift

> Consumer note: the proposal is base-templated (it names base's shared engine /
> conditions / common_types). ecds owns none of those — they live in
> `kubedge-operator-base`. The highest-value ecds-local logic is `pkg/ecdscluster`
> (renderer + manager factory). Tests target that; base's engine is base's to test.

- [x] Inventory packages with no test files. → `pkg/ecdscluster` was 43.6% (1 test file); `cmd/manager`, `pkg/controller`, `pkg/controller/ecdscluster` all 0.0% (need envtest — out of scope for a unit pass).
- [x] Add tests for the manager/renderer: owner-ref stamping + render-value mapping. → `manager_internal_test.go`: `NewECDSClusterManager` stamps a controller owner-ref from the CR + carries name/namespace; `initRenderValues` carries the stage; `initRenderFiles` empty; `ExecutionContextKind.String()`; sibling factory no-ops pinned.
- [~] Conditions package Get/Has/IsTrue/IsFalse — **N/A for ecds**: that package lives in base. Owed by base's own test-coverage-uplift.
- [~] Shared status/state helpers in common_types — **N/A for ecds**: base-owned.
- [x] `go test ./... -race` green; record coverage delta. → **`pkg/ecdscluster` 43.6% → 82.1%**; full suite green with `-race`; `golangci-lint run ./...` → 0 issues.

## Not done (needs envtest, deliberately deferred)

- `pkg/controller/ecdscluster` (the reconcile loop) stays at 0% — meaningful tests need
  `sigs.k8s.io/controller-runtime/pkg/envtest` (a real apiserver) or a fake client harness.
  The **live smoke deploy** (container-run-smoke / step 4) already exercises the full
  reconcile end-to-end, so this is lower-priority than a controller unit harness would be
  on a repo without a run gate.
