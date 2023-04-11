<!-- @overwrite-anchor=start -->

<!-- 🚀 Generated automatically by https://github.com/gizumon/github-actions-documenter 🚀 -->
<!-- Please do not edit the below manually since they are are generated automatically by this job. -->

# 🔰 Reusable Workflows 🔰

* [1: CI](#1-ci) ( [📄](.github/workflows/ci.yml) )

## 1: CI

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
| 17 |  | string | release-docker-file | Dockerfile | Set to the docker file and path if you don't want the default of `Dockerfile` in the project root. |
| 18 |  | boolean | release-docker-latest | false | If set to true, release this container as the latest. |
| 19 |  | boolean | release-docker-major | false | If set to true, release this container as the latest for the major version. |
| 20 |  | boolean | release-docker-minor | false | If set to true, release this container as the latest for the minor version. |
| 21 |  | string | release-main-package | . | Path to main.go file or main package. |
| 22 |  | string | release-project-name |  | The project name / binary name to use if not the repo name. |
| 23 |  | string | release-skip-publish |  | Set to --skip-publish to skip publishing. |
| 24 | ✅ | string | release-type |  | The type of artifact to expect and release. [ library | program ]. |
| 25 |  | boolean | copyright-skip | false | Skip validating that all files have copyright and licensing information. |
| 26 |  | boolean | style-skip | false | Skip building the gofmt check. |
| 27 |  | boolean | tests-skip | false | Skip running the unit tests. |
| 28 |  | boolean | tests-race | true | If set to "true" (default), race condition checking will be performed during unit tests.  Otherwise no race condition checking will be done. |



<!-- @overwrite-anchor=end -->

