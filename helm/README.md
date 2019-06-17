# Grafeas HELM chart
This folder contains a sample helm chart for running grafeas using helm.
The default setup will run a Greafeas instance backed by memstore or embedded [boltdb](https://github.com/boltdb/bolt) data store  with mutual tls authentication.

# Precondition
- [Persistent volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) (if embedded storage is used)
- Self signed certificates for mutual TLS authentication

# Running the chart locally
- Build the Grafeas image locally
```sh
docker build -t grafeas:latest .
```
- Create a local persistent volume
```yaml
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
```
- Create a local persistent volume claim
```yaml
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
```
- Generate self signed certificates (see example [here](https://github.com/kabukky/httpscerts)).
- Install the helm chart (note that only the ca certs are being used); use `storageType` to control the type of storage used (supported options are `embedded` or `memstore`)
```sh
helm install --name grafeas ./helm/ --set storageType="embedded" --set certificates.ca="$(cat ca.crt)" --set certificates.cert="$(cat ca.crt)" --set "certificates.key=$(cat ca.key)"
```
- Check local services and verify Grafeas is running on port 443
```sh
kubectl get svc
NAME         TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
grafeas      ClusterIP   10.245.68.7   <none>        443/TCP   79s
```
# Deleting the chart
helm delete --purge grafeas