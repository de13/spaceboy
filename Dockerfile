FROM golang:1.12-alpine3.10 AS builder
ADD goapp/ /go/src/goapp
WORKDIR /go/src/goapp
RUN go build -o goapp

FROM alpine:3.10
RUN apk update && apk add curl
COPY --from=builder /go/src/goapp/ .
ENTRYPOINT ["./goapp"]
CMD [""]
