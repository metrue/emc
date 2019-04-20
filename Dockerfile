FROM golang:1.12.0-alpine3.9 AS builder

LABEL maintainer="Minghe <h.minghe@gmail.com>"

ENV GO111MODULE on
ENV GOPROXY https://athens.azurefd.net

COPY . /go/src/github.com/metrue/fibonacci

WORKDIR /go/src/github.com/metrue/fibonacci

RUN go install -ldflags "-w -s"
RUN go build -o /go/bin/fibonacci

FROM alpine:3.9

RUN apk upgrade --no-cache && apk add ca-certificates
COPY --from=builder /go/bin/fibonacci /usr/bin/fibonacci

EXPOSE 8080

CMD ["fibonacci"]
