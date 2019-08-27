# Grafeas Server Image

Grafeas server images are available in `us.gcr.io/grafeas/grafeas-server`.
If you've never pulled/pushed images from GCR, you can find
[instructions on how to do this](https://cloud.google.com/container-registry/docs/pushing-and-pulling).

The Grafeas server image can be built via its [Dockerfile](../Dockerfile):

```bash
docker build --tag=grafeas .
```

## Pushing new Grafeas server image

Make sure you have access to push new Grafeas server images.
At this point, only official maintainers of the Grafeas project have such access.

To push a new Grafeas server image, you'll need to build it first:

```bash
... inside go/src/github.com/grafeas/grafeas ...
docker build --tag=grafeas .
```

Set the `gcloud` project:

```bash
PROJECT=grafeas
gcloud config set project $PROJECT
... if needed, update...
gcloud components update
```

Ensure docker auth is configured:

```
gcloud auth configure-docker

gcloud credential helpers already registered correctly.
```

Then, tag and push it to GCR:

```bash
docker tag grafeas us.gcr.io/grafeas/grafeas-server
docker push us.gcr.io/grafeas/grafeas-server

The push refers to repository [us.gcr.io/grafeas/grafeas-server]
a4d748b5fa54: Pushed
a464c54f93a9: Layer already exists
latest: digest: sha256:be16348fd0444426af195b33935fbf8b1479c971eda5ef5ebeb265129cfa6538 size: 739
```
