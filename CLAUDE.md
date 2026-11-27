# CLAUDE.md

## Project

- **Name:** wsaa-go
- **Purpose:** Go client for the AFIP WSAA authentication service. It gets the TOKEN and SIGN
  credentials that other AFIP web services need. A Go port of `wsaa-ruby` and `wsaa-node`.
- **Type:** library
- **Stack:** Go
- **Deploy target:** none. A Go module has no registry; `proxy.golang.org` discovers tagged
  versions on its own.

## Commands

- **Install:** `go mod download`
- **Test:** `go test ./...`
- **Lint:** `go vet ./...`

## Version

- **Version file:** `VERSION`
- **Doc format:** Go doc comments

## Specs

- Specs are in `specs/`. Name each spec `NNN-feature-name.md`. Use `specs/TEMPLATE.md`.
- Use the `/spec` skill. It pushes the spec on a `build/NNN-name` branch, and the push starts the build agent.
- Shared agent rules are in `prompts/common.md` in `leanucci/workflows`.

### Stack

- Go modules. Commit `go.sum`.
- Standard library `testing` package. Run tests with `go test ./...`.
- `go vet ./...` for lint. No separate linter.

### Code

- Use Go doc comments: a sentence starting with the identifier's name, directly above each
  exported function, type, and constant.
- Keep secrets, certificates, and private keys out of the repo. Read them from environment
  variables or from files the caller provides.

### Release Policy

- Releases happen by hand, with the `/release` skill. Specs accumulate on `main` until then.
- The build agent never changes `VERSION`. Each entry it adds under `## [Unreleased]` in
  `CHANGELOG.md` starts with its severity in parentheses, from the spec's "Release" field:
  `(major)`, `(minor)`, or `(patch)`.
- `/release` computes the new version from the highest severity under `## [Unreleased]`,
  moves those entries under that version with today's date, and sets `VERSION`.

### Publish

- There is no publish step and no registry account. Go's module proxy
  (`proxy.golang.org`) discovers a version from its git tag the first time anything fetches
  it.
- `release.yml` creates and pushes the `v<version>` tag, and the GitHub release, once
  `VERSION` changes. No secret, no token, no approval.

### First Release Setup

Nothing to set up. The first `v<version>` tag `release.yml` creates is the first thing the
module proxy indexes.
