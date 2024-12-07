# Steps to Release Grafeas

We're following [semantic versioning](https://semver.org/) approach to releases in Grafeas server. See [versioning](docs/versioning.md) document for the rationale of API, server, and client versioning.

## Create a Release PR
Modify the Grafeas version values for the containers in these locations:

* [Chart.yaml](grafeas-charts/Chart.yaml#L5)
* [values.yaml](grafeas-charts/values.yaml#L6)

Assemble all the meaningful changes since the last release into the [CHANGELOG.md](CHANGELOG.md) file.

## Tag the release

Make sure your fork of the repository is updated. Assuming `git remote` shows the `origin` (the fork) and `upstream` (the main repository), do:

```shell
git pull origin master
git pull upstream master
git tag -am "grafeas-vX.Y.Z release" vX.Y.Z
git push upstream --tags
```

NOTE: the last command will not work if you set `git remote set-url --push upstream no_push` as described in [DEVELOPMENT.md](DEVELOPMENT.md). You will need to re-enable the `push` for this to work, so proceed with caution.

You can find the releases in GitHub, e.g. [v0.1.0](https://github.com/grafeas/grafeas/releases/tag/v0.1.0).
