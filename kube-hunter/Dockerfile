# Build the manager binary
FROM golang:1.17 as builder

WORKDIR /workspace

# Copy the go source
COPY main.go main.go
COPY pkg/ pkg/
COPY vendor/ vendor/

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Build
ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -a -o ke-kubehunter main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM aquasec/kube-hunter:latest
COPY --from=builder /workspace/ke-kubehunter /usr/local/bin/
RUN addgroup -S kubeeye -g 1000 && adduser -S kubeeye -G kubeeye -u 1000
USER 1000:1000

ENTRYPOINT ["/usr/local/bin/ke-kubehunter"]
