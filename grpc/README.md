

build Protocol Buffers files
```
protoc \
--go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
./protoc/echo.proto
```
go_out -> data code
go-grpc_out -> service code

run
```
$ go run server/server.go
$ go run client/client.go -name yoshiki
```
