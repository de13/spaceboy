FROM golang:alpine AS builder
ADD goapp/ src/goapp
WORKDIR src/goapp
RUN go build -o goapp

FROM alpine
COPY --from=builder /go/src/goapp/goapp .
COPY --from=builder /go/src/goapp/templates/ templates/
ENTRYPOINT ["./goapp"]
CMD [""]
