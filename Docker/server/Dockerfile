FROM golang:1.10.2-alpine3.7
RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep 

WORKDIR /go/src/github.com/nagisa-intern/ebutuoy/server
ENV GOPATH=/go
ADD . .
RUN dep ensure -v -vendor-only=true
RUN go build -o app
ENTRYPOINT ["./app"]
