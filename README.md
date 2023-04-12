# shared-go

The shared golang workflow repository for the organization.

## Motivation

Releasing go libraries and programs is quite formulaic and hopefully boring.  Trying
to maintain lots of duplicate workflows that *should* be the same is tedious and
error prone.

The solution is to create a set of workflows that are designed to be tested and
reused throughout the organization and company.  This repository is designed to
test the workflows before they are put into production as well as provide a
way for dependabot to do much of the boring stuff for us (like keeping
dependencies up to date).  And to test everything.

## Maintaining a Repository Using `shared-go`

### Dependabot & Automation

To be efficient, dependabot really should be setup for each repo and allowed to
commit updates.  Protect robot's rights; never let a human do a robot's work.

Do ensure that branch protection rules requiring a `pull request` are in place
and require the `All checks passed check.` status to succeed to protect against
broken code from automatically being merged.

This does not prevent maintainers from pushing code as they deem reasonable, but
do think about using the workflow & a PR to verify code changes.  It is reasonable
to over-ride the requirement for an approval to merge your PR.  "You break it,
you fix it" rules are in effect.

### Clean Code

Strive to run with all the checking on in the workflow.  Despite being annoying
or tedious at times, the value is more consistent code across projects.  "Running
clean" means there are fewer exceptions to trip us up.

### Multiple Architectures

Expect to generate output for multiple architectures (x86 and arm64 today, others
in the future), so make code and containers work everywhere when possible.

### Changelog

No need to update CHANGELOG.md files anymore, as that information is automatically
gathered by gorelease for us based on commit messages.  The CHANGELOG.md files
will be removed so the source of truth is more clear soon.

### Commit Messages

The following comment "prefixes" will automatically categorize your work in the
release notes, so please use them.

- `feat(deps):` or `fix(deps):` will group the commit into **Dependency Updates**
- `feat:` or `feat([:word:]):` (where `[:word:]` is any word except `deps`) will group the commit into **New Features**
- `fix:` or `fix([:word:]):` (where `[:word:]` is any word except `deps`) will group the commit into **Bug Fixes**
- `doc:` or `doc([:word:]):` (where `[:word:]` is any word) will group the commit into **Documentation Updates**
- The following prefixes cause the commit to be omitted from the release notes.
   - `chore:`
   - `test:`
   - `merge conflict` (can happen anywhere in the commit message)
   - `Merge pull request` (can happen anywhere in the commit message)
   - `Merge remote-tracking branch` (can happen anywhere in the commit message)
   - `Merge branch` (can happen anywhere in the commit message)
   - `go mod tidy` (can happen anywhere in the commit message)
- Everything else will show up as **Other Work**

### Releasing

Releases are as simple as creating and pushing a semver tag in the following format:

```
v[0-9]+.[0-9]+.[0-9]+
```

The ci workflow will take care of the rest for you.

Example:
```bash
git pull --tags
git tag v1.0.1
git push --tags
```

Tags should be **forever** but releases can be removed if needed.

## Workflow Usage

For your go project you are going to want at least 2 workflows:

- A ci workflow.
- A dependabot approving workflow.
- (optional) A project tracking workflow.

This workflow does the following things automatically:

- generates, then builds the program (may be disabled)
- runs all the unit tests (may be disabled)
- runs gofmt style checking (may be disabled)
- runs golangci-lint checking (may be disabled)
- checks the copyright headers (may be disabled)
- dependency license checking (may be disabled)
- report results to sonarcloud (may be disabled)
- releasing via [gorelease](https://github.com/gorelease/gorelease) (may be disabled)

<!-- @overwrite-anchor=start -->

<!-- 🚀 Generated automatically by https://github.com/gizumon/github-actions-documenter 🚀 -->
<!-- Please do not edit the below manually since they are are generated automatically by this job. -->

# 🔰 Reusable Workflows 🔰


## 1: CI
## Golang CI Workflow Sample

```yaml
 SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
 SPDX-License-Identifier: Apache-2.0
---
name: 'CI'

on:
  push:
    tags:
      - 'v[0-9]+.[0-9]+.[0-9]+'
  pull_request:
  workflow_dispatch:

jobs:
  add-to-project:
    uses: xmidt-org/shared-go/.github/workflows/ci.yml@6dd1fab69f841fbea827a053e21fa83ea94774d9 # v3.0.0
    with:
      release-type: library
    secrets: inherit
```

### Inputs

| # | Required | Type | Name | Default | Description |
| :--- | :---: | :---: | :--- | :--- | :--- |
| 1 |  | string | go-version | ^1.20.x | The go version to use.  Example: '1.20.x' |
| 2 |  | boolean | go-version-latest | true | Will always use the latest version of go available. |
| 3 |  | boolean | go-generate-skip | false | Skip running go generate if needed. |
| 4 |  | string | go-generate-deps |  | The line sparated list of go generate dependencies to install via `go install` prior to running `go generate`. |
| 5 |  | string | working-directory | . | The working directory for this project. |
| 6 |  | boolean | build-skip | false | Skip building the program. |
| 7 |  | boolean | goreportcard-skip | false | Skip running the goreportcard update. |
| 8 |  | boolean | lint-skip | false | Skip running the lint check. |
| 9 |  | string | lint-timeout | 5m | The timeout to pass on to the linter. |
| 10 |  | string | lint-version | latest | The working directory for this project. |
| 11 |  | boolean | license-skip | false | Skip building the license check. |
| 12 |  | boolean | release-arch-amd64 | true | Set to enable amd64 binary and dockers to be built. |
| 13 |  | boolean | release-arch-arm64 | true | Set to enable arm64 binary and dockers to be built. |
| 14 |  | string | release-binary-name |  | If the project needs a custom name, use set it here. |
| 15 |  | boolean | release-custom-file | false | If the project needs a custom release file, use that instead. |
| 16 |  | boolean | release-docker | false | If set to true, release a container to gocr as well. |
| 17 |  | string | release-docker-extras |  | Provides a way to set the `extra_files` field with the list of files/dirs to make available. |
| 18 |  | string | release-docker-file | Dockerfile | Set to the docker file and path if you don't want the default of `Dockerfile` in the project root. |
| 19 |  | boolean | release-docker-latest | false | If set to true, release this container as the latest. |
| 20 |  | boolean | release-docker-major | false | If set to true, release this container as the latest for the major version. |
| 21 |  | boolean | release-docker-minor | false | If set to true, release this container as the latest for the minor version. |
| 22 |  | string | release-main-package | . | Path to main.go file or main package. |
| 23 |  | string | release-project-name |  | The project name / binary name to use if not the repo name. |
| 24 |  | string | release-skip-publish |  | Set to --skip-publish to skip publishing. |
| 25 | ✅ | string | release-type |  | The type of artifact to expect and release. [ library | program ]. |
| 26 |  | boolean | copyright-skip | false | Skip validating that all files have copyright and licensing information. |
| 27 |  | boolean | style-skip | false | Skip building the gofmt check. |
| 28 |  | boolean | tests-skip | false | Skip running the unit tests. |
| 29 |  | boolean | tests-race | true | If set to "true" (default), race condition checking will be performed during unit tests.  Otherwise no race condition checking will be done. |



<!-- @overwrite-anchor=end -->


## Workflow Development

There are three special directories:
- command
- library
- program

The `command` directory represents a go program in a cmd structured directory for
release testing purposes.

The `library` directory represents a go library for release purposes.

The `program` directory represents a go program for release purposes.

Releases are tested automatically now.  The repo is tagged, but the release is
not actually uploaded.  To verify the release is what you expect, comment out
the `--skip-publish` options and trigger a PR.

Delete the `1.0.x` releases when you're done.

Generally create a PR to test a change.

As it makes sense, add more `test*.yaml` workflows to implement more tests.
The idea is to generally cover the use cases as well as bugs that we find.

### Releasing

After code has been tested to prove it works, release it with a release version
starting with `v3.0.0` and follows [semantic versioning](http://semver.org/)
conventions.  Dependabot should take care of the rest for us.
