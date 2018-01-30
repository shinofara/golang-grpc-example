proto: clean
	@protoc --proto_path=proto/src --go_out=plugins=grpc:./proto proto/src/example.proto

clean:
	@rm -f proto/*.pb.go || :
