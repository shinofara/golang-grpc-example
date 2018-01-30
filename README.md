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

## 3. docker-composeを使った確認

結果killの仕方で思った結果とちがう結果が帰ってきた

```
$ docker-compose up serer
```

```
$ ps | grep docker-compose
kill -SIGTERM xxxxx
```

の場合は、docker-compose自体の停止となってしまい、コンテナ内部までSIGNALの通知が行っていない。
その為、docker-composeは自信で定めた時間は、内部プロセスが存在しても待つが、それをすぎると強制KILLを行っている

![GIF3](image/bad-pattern.gif)

docker-composeを使ってコンテナに対して正しく、SIGTERMを送るには

```
$ docker-compose kill -s SIGTERM server
```

が正しい、この場合 [test-2](https://github.com/shinofara/golang-grpc-example#2-server%E3%81%A0%E3%81%91docker%E4%B8%8A%E3%81%AB%E7%A7%BB%E5%8B%95) のときと同じ挙動を確認できた

![GIF4](image/good-pattern.gif)
