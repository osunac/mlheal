# Multi-stage Dockerfile to build Golang application

# ----------------------------------------------------------------------
# Build container
FROM golang:1.10 AS builder
WORKDIR $GOPATH/src/github.com/osunac/mlheal

# Dependencies
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

# Main program, static link
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

# ----------------------------------------------------------------------
# Run container
FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
