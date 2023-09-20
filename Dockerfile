# --- Stage 1: Build the binary
FROM --platform=$BUILDPLATFORM golang:1.20 AS builder

ARG TARGETARCH
ENV BUILD_PATH=/go/src/github.com/adilhusain-s/http-server

RUN apt-get update && apt-get install -y curl gcc git libc-dev

WORKDIR ${BUILD_PATH}
COPY ./ .

RUN go mod download

# Statically compile the Go binary
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build -o http-server main.go

# --- Stage 2: Create the final minimal image
FROM --platform=${TARGETARCH} scratch

# Copy the statically compiled binary
COPY --from=builder /go/src/github.com/adilhusain-s/http-server/http-server /http-server

# Set the entrypoint to run the application
ENTRYPOINT ["/http-server"]
EXPOSE 8080
