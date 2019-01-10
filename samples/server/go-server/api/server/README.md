# Grafeas API Reference Implementation

This is a reference implementation of the [Grafeas API Spec](https://github.com/grafeas/grafeas/blob/master/README.md).

## Overview

This reference implementation comes with the following caveats, but not limited
to:
* No ACLs are used in this implementation;
* No authorization is in place [GH issue #28](https://github.com/grafeas/grafeas/issues/28);
* Filtering in list methods is not currently supported [GH issue #29](https://github.com/grafeas/grafeas/issues/29).


### Running the server
To run the server, follow these simple steps from the root directory:

```shell
cd samples/server/go-server/api/server
go run main/main.go
```

This will start the Grafeas gRPC and REST APIs on `localhost:8080`.

To start grafeas with a custom configuration use the `-config` flag (e.g. `-config config.yaml`). See [`config.yaml.sample`](config.yaml.sample) that can be used as a starting point when creating your own config file.

### Access REST API with curl

Grafeas provides both a REST API and a gRPC API. Here is an example of using the REST API to list projects in Grafeas.

`curl http://localhost:8080/v1beta1/projects`

### Access gRPC API with a go client

[`main/client.go`](main/client.go) contains a small example of a go client that connects to Grafeas and outputs any notes in `myproject`.

```shell
go run main/client.go
```
