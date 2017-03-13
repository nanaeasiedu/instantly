FROM golang:latest

RUN apt-get update \
  && apt-get install software-properties-common python-software-properties \
  && go get bitbucket.org/liamstask/goose/cmd/goose \
  && add-apt-repository ppa:masterminds/glide \
  && apt-get update \
  && apt-get install glide

ENV SERVER_ENV production
RUN mkdir -p /go/src/github.com/ngenerio/instantly
ADD . /go/src/github.com/ngenerio/instantly
WORKDIR  /go/src/github.com/ngenerio/instantly
RUN glide install \
  && go build

ENTRYPOINT ["instantly"]

EXPOSE 3000
g
