# Running Grafeas

## Start Grafeas

To start the sample server, follow the instructions on [running the
server](https://github.com/grafeas/grafeas/tree/master/).

## Use Grafeas with self-signed certificate

### Generate CA, keys and certs

_NOTE: The steps described in this section is meant for development environments._

Make sure to set `Common Name` to your domain, e.g. localhost (without port).

```
# Create CA
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 365 -key ca.key -out ca.crt

# Create the Client Key and CSR
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr

# Create self-signed client cert
openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out client.crt

# Convert Client Key to PKCS
openssl pkcs12 -export -clcerts -in client.crt -inkey client.key -out client.p12

# Convert Client Key to (combined) PEM
openssl pkcs12 -in client.p12 -out client.pem -clcerts
```

This is basically following https://gist.github.com/mtigas/952344 with some tweaks

### Update config

Add the following to your `config.yaml` file:

```
    cafile: ca.crt
    keyfile: ca.key
    certfile: ca.crt
```

### Access REST API with curl

When using curl with a self signed certificate you need to add `-k/--insecure` and specify the client certificate.

`curl -k --cert path/to/client.pem https://localhost:8080/v1beta1/projects`

### Access gRPC with a go client

When using a go client to access Grafeas with a self signed certificate you need to specify the client certificate, client key and the CA certificate. See [main/client\_cert.go](main/client_cert.go) for an example.

```
package main

```

## Enable [CORS](https://enable-cors.org/) on the sample server

### Update config

Add the following to your config file below the `api` key:

```
    cors_allowed_origins:
       - "https://some.example.tld"
       - "https://*.example.net"
```
