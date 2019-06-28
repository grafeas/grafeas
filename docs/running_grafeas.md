# Running Grafeas Server

## Pre-requisites

* [Docker](https://www.docker.com/get-started), if planning to use Grafeas
  Docker image or build one
* [openssl](https://www.openssl.org/), if planning to use certificates

## Start Grafeas

The following options will start the Grafeas gRPC and REST APIs on `localhost:8080`.

### Using published Docker image

To start the Grafeas server from the publicly published Docker image, do:

```bash
docker pull us.gcr.io/grafeas/grafeas-server:0.1.0
docker run -p 8080:8080 --name grafeas \
  us.gcr.io/grafeas/grafeas-server:0.1.0
```

### Using Dockerfile

To start the Grafeas server from the [Dockerfile](../Dockerfile), do:

```bash
<inside the repository folder>
docker build --tag=grafeas .
docker run -p 8080:8080 --name grafeas grafeas
```

### Using Docker Compose with PostgreSQL

[grafeas-pgsql](https://github.com/grafeas/grafeas-pgsql) provides a way to run
the Grafeas server with PostgreSQL. Please refer to the instructions in the
repository to bring up the stack in your local environment.

### Using `go run`

```shell
<inside the repository folder>
cd go/v1beta1
go run main/main.go
```

### Use Grafeas with self-signed certificate

_NOTE: The steps described in this section is meant for development environments._

1. Generate CA:

    ```bash
    openssl genrsa -out ca.key 2048
    openssl req -new -x509 -days 365 -key ca.key -out ca.crt
    ```

1. Create the server key and CSR. Make sure to set `Common Name` to your domain, e.g. `localhost` (without port).

    ```bash
    openssl genrsa -out server.key 2048
    openssl req -new -key server.key -out server.csr
    ```

1. Create self-signed server certificate:

    ```bash
    openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
    ```
1. Update `config.yaml` by adding the following:

    ```
    cafile: ca.crt
    keyfile: server.key
    certfile: server.crt
    ```

## Access Grafeas API endpoints

### REST API with curl

When using curl with a self signed certificate you need to add `-k/--insecure` and specify the client certificate. To generate the combined certificate, do:

```bash
openssl pkcs12 -export -clcerts -in server.crt -inkey server.key -out server.p12
openssl pkcs12 -in server.p12 -out server.pem -clcerts
```

Now, `curl` the endpoint:

````
curl -k --cert server.pem https://localhost:8080/v1beta1/projects`
```

### gRPC with a go client

[client.go](../go/v1beta1/example/client.go) contains a small example of a go
client that connects to Grafeas and outputs the notes in `myproject`:

```bash
go run go/v1beta1/example/client.go
```

When using a go client to access Grafeas with a self signed certificate you need to specify the server certificate, server key and the CA certificate. See [cert\_client\.go](../go/v1beta1/example/cert_client/cert_client.go) for an example.

### Enable [CORS](https://enable-cors.org/) on the server

Add the following to your `config.yaml` file below the `api` key:

```
cors\_allowed\_origins:
   - "https://some.example.tld"
   - "https://*.example.net"
```
