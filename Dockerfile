FROM golang:alpine AS builder
ADD goapp/ src/goapp
WORKDIR src/goapp
RUN go build -o goapp

FROM alpine
COPY --from=builder /go/src/goapp/ .
ENTRYPOINT ["./goapp"]
CMD [""]
