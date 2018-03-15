FROM golang:1.8.4-jessie as builder
ENV buildpath=/go/src/github.com/mad01/kubectl-client
RUN mkdir -p $buildpath
WORKDIR $buildpath

COPY . .

RUN make build/release

FROM debian:8
COPY --from=builder /go/src/github.com/mad01/kubectl-client/_release/kubectl-client /kubectl-client

ENTRYPOINT ["/kubectl-client"]
