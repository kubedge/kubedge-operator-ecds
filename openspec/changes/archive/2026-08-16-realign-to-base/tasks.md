# Tasks — realign-to-base

- [x] `go get github.com/kubedge/kubedge-operator-base@v0.1.36-kubedge.20260815`
- [x] `go mod tidy` (confirmed k8s.io/* → v0.36.3, controller-runtime → v0.24.1, go → 1.26)
- [x] `go build ./...` → fixed the controller-runtime Watch API breakage (`c.Watch(source.Kind(cache, client.Object(&ECDSCluster{}), handler))`). The printf-vet family did NOT fire (`go vet` clean).
- [x] `go vet ./... && go test ./... -race` green.
- [x] `golangci-lint run … ./...` recorded — 3 issues at the time; cleared later under adopt-go-ci's lint-cleanup (now 0).
- [x] Confirm the operator still builds its binary — `go build ./cmd/...` green; the container image (`build/Dockerfile`) compiles + runs it live (smoke gate).
