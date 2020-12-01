# Development

This doc explains the development workflow so you can get started
[contributing](CONTRIBUTING.md) to Grafeas!

## Getting started

First, you will need to setup your GitHub account and create a fork:

1. Create [a GitHub account](https://github.com/join)
1. Setup [GitHub access via
   SSH](https://help.github.com/articles/connecting-to-github-with-ssh/)
1. [Create and checkout a repo fork](#checkout-your-fork)

Once you have those, you can iterate on Grafeas:

1. [Run your instance of Grafeas](docs/running_grafeas.md)
1. [Run Grafeas tests](#testing-grafeas)

When you're ready, you can [create a PR](#creating-a-pr)!

## Checkout your fork

The Go tools require that you clone the repository to the `src/github.com/grafeas/grafeas` directory
in your [`GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH).

To check out this repository:

1. Create your own [fork of this
  repo](https://help.github.com/articles/fork-a-repo/)
2. Clone it to your machine:

  ```shell
  mkdir -p ${GOPATH}/src/github.com/grafeas
  cd ${GOPATH}/src/github.com/grafeas
  git clone git@github.com:${YOUR_GITHUB_USERNAME}/grafeas.git
  cd grafeas
  git remote add upstream git@github.com:grafeas/grafeas.git
  git remote set-url --push upstream no_push
  ```

_Adding the `upstream` remote sets you up nicely for regularly [syncing your
fork](https://help.github.com/articles/syncing-a-fork/)._

## Testing Grafeas

Grafeas has unit tests, which can be run with:

```shell
make test
```

:warning: These tests will not run correctly unless you have [checked out your fork into your `$GOPATH`](#checkout-your-fork).
:warning: Ensure that [`$GOPATH/bin` is on your system `$PATH`](https://github.com/golang/protobuf/issues/795#issuecomment-564523540).
:warning: These test require either a running Postgres instance (with appropriate credentials) or having the Postgres bin directory on your system `PATH`.

## Creating a PR

When you have changes you would like to propose to Grafeas, you will need to:

1. Ensure the commit message(s) describe what issue you are fixing and how you are fixing it
   (include references to [issue numbers](https://help.github.com/articles/closing-issues-using-keywords/)
   if appropriate)
1. [Create a pull request](https://help.github.com/articles/creating-a-pull-request-from-a-fork/)

### Reviews

Each PR must be reviewed by a maintainer. You may be asked to meet with a member
of the core Grafeas team to discuss the PR at high level, before they start the
detailed review.
