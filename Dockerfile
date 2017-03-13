FROM golang:latest

RUN apt-get update \
  && apt-get -y install software-properties-common python-software-properties \
  && go get bitbucket.org/liamstask/goose/cmd/goose \
  && apt-get update

ENV SERVER_ENV production
RUN mkdir -p /go/src/github.com/ngenerio/instantly
ADD . /go/src/github.com/ngenerio/instantly
WORKDIR  /go/src/github.com/ngenerio/instantly
RUN go build

ENTRYPOINT ["/go/src/github.com/ngenerio/instantly/instantly"]

EXPOSE 3000
