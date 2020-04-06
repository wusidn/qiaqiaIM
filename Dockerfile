FROM golang:1.14-alpine3.11 


ENV GO111MODULE on
ENV GOPROXY https://goproxy.io,direct

WORKDIR /go/src/github.com/wusidn/qiaqia/

RUN apk add git vim \
    && go get -u github.com/kardianos/govendor

CMD ["/bin/sh"]