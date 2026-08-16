# Tasks — adopt-go-ci

- [x] From the meta session, run `/alemax:update-skills` so the class-M set (incl. `ci.yml`) is staged for this repo. → done meta-side; branch `meta-broadcast/deliver-class-m-templates-ciyml--hygiene` on origin + worktree `../kubedge-operator-ecds-claude-meta`.
- [x] In this repo's session, run `/alemax:complete-update` to apply the update branch. → **cherry-picked** the broadcast onto `chore/openspec-seed` (not `main` — the whole retrofit arc lives on the seed branch and merges via one PR). 10 files landed clean (all adds).
- [x] Confirm `.github/workflows/ci.yml` present and its jobs gate on `go.mod`. → present; go-build/go-vet/go-test/golangci-lint gated on `detect.outputs.go`.
- [~] Trial push the branch; confirm `go-build`/`go-vet`/`go-test`/`golangci-lint` are green. → **branch pushed**, but `ci.yml` triggers only on `push`/`pull_request` to `main`, so CI fires when the **PR** is opened (pending — see below). Predicted green locally: `go build/vet/test -race` ✅, `golangci-lint run ./...` → **0 issues**, `shellcheck bin/set-secret.sh` ✅.
- [x] Confirm the rest of class-M landed: `.editorconfig`, `.gitattributes`, `.github/*`, `dependabot.yml`, `.pre-commit-config.yaml`, `bin/set-secret.sh`. → all present.

## Adaptations for a green trial

- **Re-trimmed `ci.yml` to the go stack** (Phase 3c): removed the python-only `type-check` job; added the `gomod` Dependabot ecosystem. Committed separately (`fix(ci): re-trim …`).
- **golangci-lint pin bumped** v1.61.0 → **v2.12.2** (action v6 → v7), matching operator-base — v1 can't lint a `go 1.26` module.
- **lint-cleanup** to reach `0 issues` on v2 defaults (base runs configless + clean, so no `.golangci.yml`):
  - `//nolint:staticcheck` on `mgr.GetEventRecorderFor` (SA1019 — the new `GetEventRecorder` returns an incompatible `events.EventRecorder` vs base's `record.EventRecorder` field; base's convention).
  - dropped redundant `fmt.Sprintf("%s", string(blob))` in the renderer test (S1025).
  - removed dead `pkg/ecdscluster/log.go` (unused `var log`).

## Remaining

- Open PR `chore/openspec-seed` → `main` to fire the `pull_request` CI and confirm the 4 Go jobs green (also clears the 4 stale dependabot alerts on `main`, which are pre-realign).
