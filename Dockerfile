FROM golang:1.20-alpine as builder
# Args & ENVs
ENV BUILD_PATH=/go/src/github.com/adilhusain-s/http-server

RUN apk update && apk add --no-cache curl gcc git libc-dev
# COPY local files
WORKDIR ${BUILD_PATH}
COPY ./ .

# Get go dependencies
RUN go mod download 
RUN  go build  -o  http-server main.go
RUN cp ./http-server /bin/http-server 

# --- Stage 2:
FROM alpine:3
# Install dependencies
RUN apk update && apk add --no-cache ca-certificates tzdata libc6-compat
# Copy binary from builder
COPY --from=builder /bin/http-server /http-server
COPY --from=builder /go/src/github.com/adilhusain-s/http-server/testdata /testdata
# Run the application on container startup.
CMD ["/http-server"]
EXPOSE 8080