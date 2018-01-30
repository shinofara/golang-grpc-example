FROM golang:alpine AS build-env

COPY . /go/src/github.com/shinofara/golang-grpc-example/client
WORKDIR /go/src/github.com/shinofara/golang-grpc-example/client
RUN go build -ldflags '-s -w' -o /bin/client ./main.go

FROM alpine

ENV TZ Asia/Tokyo
ENV LC_ALL ja_JP.UTF-8
ENV PATH /bin:$PATH

RUN apk --update add tzdata ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && rm -rf /var/cache/apk/*

COPY --from=build-env /bin/client /bin/client

EXPOSE 19003
ENTRYPOINT ["/bin/client"]
