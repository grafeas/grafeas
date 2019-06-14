# Grafeas HELM chart
This folder contains a sample helm chart for running grafeas using helm.
The setup will run a Greafeas instance backed by memstore by default, or embedded [boltdb](https://github.com/boltdb/bolt) data store with mutual tls authentication.

# Precondition
- Self signed certificates for mutual TLS authentication;
- If using embedded boltdb, [persistent volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/).

# Running the chart locally

TODO: update the instructions once the image is published

Build the Grafeas image locally
```sh
docker build -t grafeas:latest .
```

Generate self-signed certificates by following [instructions](../docs/running_grafeas.md#use-grafeas-with-self-signed-certificate.

If using in-memory store, do:

```
helm install --name grafeas ./helm/ --set storageType="memstore" --set certificates.ca="$(cat ca.crt)" --set certificates.cert="$(cat ca.crt)" --set "certificates.key=$(cat ca.key)"
```

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
helm install --name grafeas ./helm/ --set storageType="embedded" --set certificates.ca="$(cat ca.crt)" --set certificates.cert="$(cat ca.crt)" --set "certificates.key=$(cat ca.key)"
```

Check local services and verify Grafeas is running on port 443
```sh
kubectl get svc

NAME         TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
grafeas      ClusterIP   10.245.68.7   <none>        443/TCP   79s

kubectl get pods
```

# Deleting the chart
helm delete --purge grafeas
