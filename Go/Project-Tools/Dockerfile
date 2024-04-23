# sytnax=docker/dockerfile:1
FROM golang:1.22.2 AS builder

WORKDIR /workdir
COPY . /workdir

ENV CGO_CPPFLAGS="-D_FORTIFY_SOURCE=2 -fstack-protector-all"
ENV GOFLAGS="-buildmode=pie"

RUN go build -ldflags "-s -w" -trimpath ./cmd/api

# Second Stage
FROM gcr.io/distroless/base-debian11:nonroot
COPY --from=builder /workdir/api /bin/api

USER 65534

ENTRYPOINT ["/bin/app"]