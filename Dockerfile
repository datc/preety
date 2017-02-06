FROM daocloud.io/library/golang:1.6.0
MAINTAINER toukii "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/
RUN go get github.com/datc/pretty
EXPOSE 8080
CMD ["pretty"]