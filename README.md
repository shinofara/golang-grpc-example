# golang-grpc-example

## 1. Example 1

Mac上で、gRPC ServerとgRPC Client x 2を実行して、SIGTERM

gRPC Server

```
$ cd server
$ ./server
```

gRPC Client

```
$ cd client
$ go run main.go --id 1
$ go run main.go --id 2
```

Kill -SIGTERM

```
$ pkill -SIGTERM server
```

![GIF1](image/grpc-graceful-direct-connect.gif)

## 2. ServerだけDocker上に移動

### 手順

Server起動

```
$ cd server
$ docker build -t a .
$ docker run -p 19003:19003 --rm a
```

Client1

```
$ cd client
$ go run main.go
```

GracefulStop

```
pkill -fl -SIGTERM bin/docker
```

Client2

Graceful中に新しいリクエストを受け付けない事を確認

```
$ cd client
$ go run main.go --id 1
```

![GIF2](image/grpc-graceful-server-on-docker.gif)
