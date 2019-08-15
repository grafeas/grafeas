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

You can find the releases in Github, e.g. [v0.1.0](https://github.com/grafeas/grafeas/releases/tag/v0.1.0).

## Release the Docker image

Ensure you have the right access to push images. Set up the `gcloud` project:

```shell
PROJECT=grafeas
gcloud config set project $PROJECT
gcloud auth configure-docker
... gcloud credential helpers already registered correctly.
```

Now, build the Grafeas server image inside the fork, `go/src/github.com/grafeas/grafeas`:

```shell
docker build --tag=grafeas .
docker tag grafeas us.gcr.io/grafeas/grafeas-server:vX.Y.Z
docker push us.gcr.io/grafeas/grafeas-server:vX.Y.Z
```

## Release the Grafeas Helm chart

You'll need to create service account key for `grafeas-helm-manual-release` and
download it in JSON format. See [instructions](https://cloud.google.com/docs/authentication/production) in the *Obtaining and providing service account credentials manually* section. Once you downloaded the key, set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable:

```shell
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/json-key.json"
```

Inside the fork, `go/src/github.com/grafeas/grafeas`, run the commands below:

```shell
cd grafeas-charts
helm init --client-only
helm plugin install https://github.com/hayorov/helm-gcs
helm repo add grafeas-charts-repository gs://grafeas-charts/repository
helm package .
helm gcs push grafeas-charts-*.tgz grafeas-charts-repository
```

[Grafeas Helm chart](https://storage.cloud.google.com/grafeas-charts/repository/grafeas-charts-0.1.0.tgz?organizationId=433637338589).
