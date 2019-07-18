
FROM golang:1.12.5 as base
COPY . /app
WORKDIR /app

FROM base as dev
CMD go run go/v1beta1/main/main.go

FROM base as builder
RUN CGO_ENABLED=0 go build -o grafeas-server go/v1beta1/main/main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/grafeas-server /grafeas-server
EXPOSE 8080
ENTRYPOINT ["/grafeas-server"]
