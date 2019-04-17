FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/nathanjsweet/test-app
RUN go get -d -v github.com/armon/go-proxyproto
COPY . .
RUN go build main.go

FROM alpine:latest 
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/nathanjsweet/test-app/main .
CMD ["./main"]