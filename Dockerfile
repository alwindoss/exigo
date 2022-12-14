FROM golang:1.18 AS builder
WORKDIR /go/src/app
COPY . .
RUN make setup
RUN make docker

FROM busybox:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/bin/exigo .
CMD [ "./exigo" ]