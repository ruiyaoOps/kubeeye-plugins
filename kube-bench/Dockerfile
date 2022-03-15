# Build the manager binary
FROM golang:1.17 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY vendor/ vendor/

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod tidy
RUN go mod vendor

# Build
ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -a -o ke-kubebench main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM aquasec/kube-bench:v0.6.6
COPY --from=builder /workspace/ke-kubebench /usr/local/bin/
USER 65532:65532

ENTRYPOINT ["/usr/local/bin/ke-kubebench"]
