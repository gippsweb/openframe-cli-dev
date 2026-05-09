# flamingo-stack → gippsweb Migration

**Date:** 2026-05-09  
**Repo:** `openframe-cli-dev` (`github.com/gippsweb/openframe-cli-dev`)

## Goal

Migrate all Go module paths and select config references from `github.com/flamingo-stack/openframe-cli` to `github.com/gippsweb/openframe-cli-dev`.

## Steps

### 1. Pull from origin
```
git pull origin main
```
Restores the working tree in `/home/mark/openframe-cli-dev/` (currently shows all files deleted locally).

### 2. Go module rename (137 files + go.mod)
Replace across all `.go` files and `go.mod`:
```
github.com/flamingo-stack/openframe-cli → github.com/gippsweb/openframe-cli-dev
```

### 3. .goreleaser.yml (3 lines)
Same replacement for linker `-X` flags that embed version info.

### 4. helm-values-example.yaml (1 line)
```
github.com/flamingo-stack/openframe-oss-tenant → github.com/gippsweb/openframe-oss-tenant
```

### 5. Commit
Single commit on `main`.

### 6. Push
User will supply a PAT for the push to `origin`.

## Explicitly Out of Scope

- All `README.md`, `CONTRIBUTING.md`, `RELEASE_NOTES.md`, `SECURITY.md` — unchanged
- All `ghcr.io/flamingo-stack/...` container image references — unchanged
- `github.com/flamingo-stack/openframe-saas-tenant` in `helm-values-example.yaml` — unchanged
