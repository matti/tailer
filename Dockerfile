FROM golang:1.17-alpine3.15 as builder

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -o /tailer

FROM scratch
COPY --from=builder /tailer /
ENTRYPOINT [ "/tailer" ]