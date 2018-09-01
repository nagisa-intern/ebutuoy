FROM golang:1.8

WORKDIR /go/src/github.com/nagisa-intern/ebutuoy/server
ENV GOPATH=/go
ADD . .
RUN dep ensure -v -vendor-only=true
RUN go build -o app
ENTRYPOINT ["./app"]
