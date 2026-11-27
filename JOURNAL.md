# Journal

A short entry for each merged pull request: the requirement and the key decisions.
The merge workflow adds each entry before it merges the pull request.

## 2026-11-27 — issue #1: Package skeleton: environment config and error types

**Requirement:** Create the initial `wsaa` package. Add the Environment type with the testing and production endpoint URLs. Add the credentials, SOAP fault, and response error types.

**Decisions:** The package name is `wsaa`, at the module root. `Environment.URL()` returns the full LoginCms URL from the wsaa-node README. An unknown Environment value gives an empty URL. The error types are struct types with pointer receivers, so callers can use `errors.As`. The CHANGELOG entries start with `(none)`, because the issue Release field is `none`.
