# Grafeas HELM chart

This folder contains a sample helm chart for running Grafeas using helm on k8s.
The setup will run a Grafeas instance backed by memstore by default, or embedded [boltdb](https://github.com/boltdb/bolt) data store, with mutual TLS authentication.

## Requirements

* [Kubernetes](https://kubernetes.io/)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* [Helm](https://helm.sh/)

## Running the chart

* Basic example, without certificate.
```
$ helm install  grafeas ./grafeas-charts/ --set container.port=4000 --set certificates.enabled=false
$ export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=grafeas-server,app.kubernetes.io/instance=grafeas" -o jsonpath="{.items[0].metadata.name}")
$ kubectl --namespace default port-forward $POD_NAME 4000
```
Note: this basically forwards your localhost:4000 to port 4000 of the pod.

Now open another terminal:
```
$ curl http://localhost:4000/v1beta1/projects
{"projects":[],"nextPageToken":""}%  
```

* Basic example, with certificate.

First, generate self-signed certificates by following [instructions](../docs/running_grafeas.md#use-grafeas-with-self-signed-certificate).

```
$ helm install  grafeas ./grafeas-charts/ --set container.port=5000 --set secret.enabled=true  --set certificates.enabled=true --set service.port=5000 --set certificates.name="foo" --set certificates.ca="$(cat ca.crt)" --set certificates.cert="$(cat server.crt)" --set "certificates.key=$(cat server.key)"
$ export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=grafeas-server,app.kubernetes.io/instance=grafeas" -o jsonpath="{.items[0].metadata.name}")
$ kubectl --namespace default port-forward $POD_NAME 5000
```     
Now open another terminal:
```
$ curl -k --cert server.pem https://localhost:5000/v1beta1/projects
Enter PEM pass phrase:
{"projects":[],"nextPageToken":""}%
```
Note that in the above basic examples, we used the in-memory store.


If using embedded boltdb, create a local persistent volume and a claim:

```shell
cat <<EOF | kubectl apply -f - \

kind: PersistentVolume
apiVersion: v1
metadata:
  name: task-pv-volume
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/var/grafeas"
EOF

cat <<EOF | kubectl apply -f - \

kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: "grafeas"
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
EOF
```

Now, install the helm chart:

```sh
helm install --name grafeas ./helm/ --set storageType="embedded" --set certificates.ca="$(cat ca.crt)" --set certificates.cert="$(cat server.crt)" --set "certificates.key=$(cat server.key)"
```

Check local services and verify Grafeas is running on port 443:

```sh
kubectl get svc

NAME         TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
grafeas-server      ClusterIP   10.245.68.7   <none>        443/TCP   79s

kubectl get pods

NAME                              READY   STATUS      RESTARTS   AGE
grafeas-server-4cf696-ncbk7   1/1     Running     0          17h
```

# Deleting the chart

```sh
helm delete --purge grafeas
```

## Configuration

The following table lists the configurable parameters of the Grafeas chart and their default values.

| Parameter                                   | Description                               | Default                                    |
| ------------------------------------------  | ----------------------------------------  | -------------------------------------------|
| `replicaCount`                                | Number of replicas deployed               | `1`                                          |
| `deploymentStrategy`                          | Deployment strategy                       | `{}`                                         |
| `image.repository`                            | Image repository                          | `us.gcr.io/grafeas`                          |
| `image.name`                                  | Image name                                | `grafeas-server`                             |
| `image.tag`                                   | Image tag                                 | `v0.1.0`                                     |
| `image.pullPolicy`                            | Image pull policy                         | `IfNotPresent`                               |
| `nameOverride`                                | App name                                  | `grafeas-server`                             |
| `fullnameOverride`                            | App name                                  | `grafeas-server`                             |
| `persistentVolumeClaimName`                   | The name of persistent volume             | `grafeas`                                    |
| `storageType`                                 | The type of storage used, supported options: memstore or embedded | `memstore`           |
| `service.type`                                | Kubernetes Service type                   | `ClusterIP`                                  |
| `service.port`                                | Kubernetes Service port                   | `8080`                                       |
| `container.port`                              | Grafeas container port                    | `8080`                                       |
| `certificates.enabled`                        | Whether to enable client certificates for auth | `false`                                 |
| `certificates.name`                           | Certificate name                          | `grafeas-ssl-certs`                          |
| `certificates.ca`                             | Certificate CA                            | `null`                                       |
| `certificates.cert`                           | Certificate body                          | `null`                                       |
| `certificates.key`                            | Certificate key                           | `null`                                       |
| `resources`                                   | CPU/Memory resource requests/limits       | `{}`                                         |
| `resources.limits.cpu`                        | CPU limit                                 | `100m`                                       |
| `resources.limits.memory`                     | Memory limit                              | `128Mi`                                      |
| `resources.requests.cpu`                      | CPU requests                              | `100m`                                       |
| `resources.requests.memory`                   | Memory requests                           | `128Mi`                                      |
| `livenessprobe.initialDelaySeconds`           | Liveness probe initial delay seconds      | `15`                                         |
| `livenessprobe.periodSeconds`                 | Liveness probe period seconds             | `10`                                         |
| `livenessprobe.failureThreshold`              | Liveness probe failure threshold          | `3`                                          |
| `readinessprobe.initialDelaySeconds`          | Readiness probe initial delay seconds     | `15`                                         |
| `readinessprobe.periodSeconds`                | Readiness probe period seconds            | `10`                                         |
| `readinessprobe.failureThreshold`             | Readiness probe failure threshold         | `3`                                          |
| `nodeSelector`                                | Node labels for pod assignment            | `{}`                                         |
| `tolerations`                                 | Toleration labels for pod assignment	  | `[]`                                         |
| `affinity`                                    | Affinity settings for pod assignment      | `{}`                                         |
