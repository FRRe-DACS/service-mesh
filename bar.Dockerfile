# Build the manager binary
FROM golang:1.14.3-alpine as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.sum ./
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bar ./cmd/bar/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /workspace/bar /bar

EXPOSE 8080
EXPOSE 9090

ENTRYPOINT ["/bar"]