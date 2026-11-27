# Journal

A short entry for each merged pull request: the requirement and the key decisions.
The merge workflow adds each entry before it merges the pull request.

## 2026-11-27 — issue #1: Package skeleton: environment config and error types

**Requirement:** Create the initial `wsaa` package. Add the Environment type with the testing and production endpoint URLs. Add the credentials, SOAP fault, and response error types.

**Decisions:** The package name is `wsaa`, at the module root. `Environment.URL()` returns the full LoginCms URL from the wsaa-node README. An unknown Environment value gives an empty URL. The error types are struct types with pointer receivers, so callers can use `errors.As`. The CHANGELOG entries start with `(none)`, because the issue Release field is `none`.

## 2026-11-27 — issue #3: Remove the portfolio mention from README.md

**Requirement:** Remove the mention of leanucci/portfolio from README.md. The README must not point readers at an internal process document.

**Decisions:** The build removes the full "agent cycle" paragraph, because its only purpose is the link. A Go test reads `README.md` and checks both acceptance criteria. The CHANGELOG entry starts with `(none)`, because the issue Release field is `none`.

## 2026-11-27 — 001: Richer README

**Requirement:** Expand the README to the depth of the wsaa-node and wsaa-ruby READMEs. Add install, usage, endpoints, errors, and development sections. Document only what the code does today.

**Decisions:** The intro keeps the client description, because an earlier test requires its words. The intro now calls it the goal of the client. A new Status section tells that the login, TRA generation, signing, and SOAP client do not exist yet. An `ExampleEnvironment` test runs the Usage code. Go tests parse the README and check each acceptance criterion. The CHANGELOG entry starts with `(none)`, because the spec Release field is `none`.
