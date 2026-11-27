# 001: Richer README

- **Release:** none
- **Human approval:** not required

## Summary

Expand the README to match the depth of `wsaa-node`'s and `wsaa-ruby`'s: install instructions,
a usage example for what exists today, the configuration and endpoints, the error types, and
how to develop the project. Document only what the code does today.

## Requirements

1. The README has an Install section: `go get github.com/leanucci/wsaa-go`, and the minimum Go
   version from `go.mod`.
2. The README has a Usage section with a working code example for the `Environment` type and
   its `URL()` and `String()` methods, since that is the only public API today.
3. The README has an Endpoints section listing the testing and production LoginCms URLs, from
   `environment.go`.
4. The README has an Errors section listing `CredentialsError`, `SoapFaultError`, and
   `ResponseError`, each with what it means, matching their doc comments in `errors.go`.
5. The README has a Development section with the commands from `CLAUDE.md`.
6. The README does not claim a login method, TRA generation, signing, or a SOAP client exist.
   None of them do yet.

## Acceptance Criteria

1. Given the README's Usage section, when its code example is read, then it only calls methods
   that exist in `environment.go` today.
2. Given the README's Errors section, when compared to `errors.go`, then every exported error
   type is listed, and none are invented.
3. Given the README's Endpoints section, when compared to `environment.go`, then the URLs
   match exactly.
4. Given the README's Development section, when compared to `CLAUDE.md`'s Commands, then every
   command listed there appears.

## Out of Scope

- A login usage example, a Login cooldown section, and a Credential persistence section, all
  present in `wsaa-node`'s README. They describe the client, which this project has not built
  yet. Add them in the spec that adds the client.
- Changing `environment.go` or `errors.go`.

## Notes

Reference READMEs:
- `wsaa-node`: https://github.com/leanucci/wsaa-node/blob/main/README.md
- `wsaa-ruby`: https://github.com/leanucci/wsaa-ruby/blob/main/README.md

`wsaa-node`'s is the more complete structure to follow; `wsaa-ruby`'s predates it and is
simpler. Match sections to what the Go code actually has, not a 1:1 copy of either.
