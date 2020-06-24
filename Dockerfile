FROM golang:1.14 as builder

WORKDIR /

ADD main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o dns-load main.go

FROM debian

COPY --from=builder /dns-load /dns-load

ENTRYPOINT [ "/dns-load" ]