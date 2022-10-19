## Run
Please check [Prerequisite](#prerequisite) before runing
```sh
nx run-many --target=serve
```

## TODO
- Buildscript
- DB setup
- Talk about grpc and api
- code linter / format
  - prettier / editorconfig
- Github workflows
- Autmated test

## prerequisite
- Golang
- Protoc
  - protoc-gen-go-grpc
- ~~cargo (rust)~~
- npm / pnpm
- Nx cli

### Golang
[Golang official website](https://go.dev/doc/install)

### Protoc-gen-go-grpc
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Nx cli
```sh
npm i -g nx
```