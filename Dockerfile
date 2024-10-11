
ARG ALPINE_VERSION=3.20
ARG GO_VERSION=1.23
ARG GOLANGCI_LINT_VERSION=v1.61.0

FROM qmcgaw/binpot:golangci-lint-${GOLANGCI_LINT_VERSION} AS golangci-lint

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
RUN apk --update add git g++
ENV CGO_ENABLED=0
COPY --from=golangci-lint /bin /go/bin/golangci-lint
WORKDIR /tmp/gobuild
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM base AS test
# Note on the go race detector:
# - we set CGO_ENABLED=1 to have it enabled
# - we installed g++ to support the race detector
ENV CGO_ENABLED=1
ENTRYPOINT go test -race -coverprofile=coverage.txt ./...

FROM base AS lint
RUN golangci-lint run --timeout=10m