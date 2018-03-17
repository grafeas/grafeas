# How to Contribute

We'd love to accept your patches and contributions to this project. There are
just a few small guidelines you need to follow.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution,
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Proposals and PRs

If you would like to make a large change, please start with a proposal issue that includes:
* What you would like to achieve
* Why you'd like to make this change
* A design overview

## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Getting started

To build the codebase, you're going to need a few things:
- [Go](https://golang.org/doc/install)
- [Dep](https://golang.github.io/dep/docs/installation.html)

The first time you clone the repository, make sure to run the prepare task:
```
$ make prepare
```

Now you should be able to build.
```
$ make build
```

Before submitting your PR, make sure to run the tests!
```
$ make test
```

We're using [dep](https://golang.github.io/dep/) for dependency management.
If you want to add a new dependency, make sure to use the [`dep ensure --add path/to/dep`
command](https://golang.github.io/dep/docs/daily-dep.html#using-dep-ensure).

If you have any questions feel free to [file an issue!](https://github.com/grafeas/grafeas/issues)
