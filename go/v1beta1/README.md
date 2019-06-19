# Grafeas API Reference Implementation

This is an implementation of the [Grafeas v1beta1 API spec](https://github.com/grafeas/grafeas/tree/master/proto/v1beta1).

## Running the server
To run the server, do the following from the root directory:

```shell
cd go/v1beta1
go run main/*.go
```

This will start the Grafeas gRPC and REST APIs on `localhost:8080`.

To start grafeas with a custom configuration use the `--config` flag (e.g. `--config config.yaml`). See [`config.yaml.sample`](config.yaml.sample) that can be used as a starting point when creating your own config file.

### Access REST API with curl

Grafeas provides both a REST API and a gRPC API. Here is an example of using the REST API to list projects in Grafeas.

`curl http://localhost:8080/v1beta1/projects`

### Access gRPC API with a go client

[`example/client.go`](example/client.go) contains a small example of a go client that connects to Grafeas and outputs the notes in `myproject`.

```shell
go run example/client.go
```
