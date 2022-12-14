## Install Dependency
```sh
nx run-many --target=install
```

## Run
Please check [Prerequisite](#prerequisite) before runing
```sh
nx run-many --target=serve
```

## Docker lazy script
```sh
git clone --sparse https://github.com/chula-overflow/chula-overflow-backend
cd chula-overflow-backend
cp .env.example .env
docker-compose --env-file .env up -d
```

## TODO
- code linter / format
  - prettier / editorconfig
- Automated test

## prerequisite
- [Golang](https://go.dev/doc/install)
- Protoc
  - protoc-gen-go-grpc
- [cargo (rust)](https://rustup.rs/)
- npm / pnpm
- Nx cli

### Protoc-gen-go-grpc
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Nx cli
```sh
npm i -g nx
```

## Docker
### Docker compose
```sh
docker compose --env-file .env up
```
### Building image
Since each service depends on `apps/proto` directory, images need to be build from project root
```sh
docker build <service>:<tag> -f apps/<service>/Dockerfile .
```

## Note
Please don't mind that auth take like 5 minutes compile time.

## Env
.env is supposed to use for containerization purpose. It will **not** automatically load into any service. please load it yourself.
