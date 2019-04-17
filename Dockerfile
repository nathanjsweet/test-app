FROM golang:1.12.4-alpine as builder
WORKDIR /go/src/github.com/nathanjsweet/test-app
COPY . .
RUN go build main.go

FROM alpine:latest 
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/nathanjsweet/test-app/main /root/app
CMD ["/root/app"]