# reference
https://github.com/deepmap/oapi-codegen/tree/master

# build
```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
oapi-codegen -config api/type.yml openapi.yaml
oapi-codegen -config api/server.yml openapi.yaml
```

# Files
```
.
├── README.md // this file
├── api // API controllers
│   ├── api.gen.go // generated API handler + server interface
│   ├── api.go // server implementation
│   ├── server.yml // config to generate api.gen.go
│   ├── type.gen.go // generated types
│   └── type.yml // config to generate type.gen.go
├── go.mod
├── go.sum
├── main.go
├── openapi.yaml // source of api.gen.go & type.gen.go
└── test.http // test file for vscode-restclient
```
