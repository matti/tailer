FROM golang:1.14-alpine3.12 as builder

WORKDIR /src
COPY src .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /tailer

FROM scratch
COPY --from=builder /tailer /
ENTRYPOINT [ "/tailer" ]