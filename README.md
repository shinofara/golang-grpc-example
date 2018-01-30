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