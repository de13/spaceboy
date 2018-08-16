FROM golang:alpine AS builder
ADD workshop/ src/goapp
WORKDIR src/goapp
RUN go build -o goapp

FROM alpine
COPY --from=builder /go/src/goapp/goapp .
ENTRYPOINT ./goapp
