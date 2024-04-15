FROM go:latest as builder

ENV CGO_ENABLED=0

WORKDIR /build

COPY ./main.go .

RUN go build -o main ./main.go

FROM alpine

WORKDIR /dist

COPY --from=builder /builder/main .

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip

CMD ["/dist/main"]
