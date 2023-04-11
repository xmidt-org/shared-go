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

## Workflow Usage

For your go project you are going to want at least 2 workflows:

- A ci workflow.
- A dependabot approving workflow. (See: [shared-dependabot-approver](https://github.com/comcast-cl/shared-dependabot-approver) for details.)
- (optional) A project tracking workflow. (See: [shared-project](https://github.com/comcast-cl/shared-project) for details.)

This workflow does the following things automatically:

- generates, then builds the program (may be disabled)
- runs all the unit tests (may be disabled)
- runs gofmt style checking (may be disabled)
- runs golangci-lint checking (may be disabled)
- checks the copyright headers using the [shared-copyright](https://github.com/comcast-cl/shared-copyright) workflow (may be disabled)
- dependency license checking (may be disabled)
- report results to sonarcloud (may be disabled)
- releasing via [gorelease](https://github.com/gorelease/gorelease) (may be disabled)

<!-- @overwrite-anchor=start -->

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
