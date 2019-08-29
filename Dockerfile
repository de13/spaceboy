FROM golang:alpine AS builder
ADD goapp/ /go/src/goapp
WORKDIR /go/src/goapp
RUN go build -o goapp

FROM alpine
RUN apk update && apk add curl
COPY --from=builder /go/src/goapp/ .
ENTRYPOINT ["./goapp"]
CMD [""]
