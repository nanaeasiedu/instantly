FROM alpine:3.3
MAINTAINER Eugene Asiedu <ngene84@gmail.com>

ENV INSTANTLY_VERSION 0.0.1-rc

RUN \
  apk update && \
  apk add ca-certificates && \
  update-ca-certificates && \
  cd /tmp && \
  wget https://github.com/ngenerio/instantly/releases/download/v$instantly_VERSION/instantly_linux_amd64.zip && \
  unzip instantly_linux_amd64.zip -d /usr/bin && \
  mv /usr/bin/instantly_linux_amd64 /usr/bin/instantly && \
  rm -f instantly_linux_amd64.zip

EXPOSE 5000
CMD ["/usr/bin/instantly"]
