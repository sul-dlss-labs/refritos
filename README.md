# Refritos

The next generation processing framework that will support [TACO](https://github.com/sul-dlss-labs/taco)

## Client Work

[Ruby KCL Branch](https://github.com/sul-dlss-labs/refritos/tree/ruby_kcl)
[Go KCL Router](https://github.com/sul-dlss-labs/refritos/tree/go_kcl_routerl)

## Running GO KCL Router

_NOTE:_ This is a investigative prototype, take all work with a grain of salt (or squeeze of lime).

### Export necessary ENV VARS

```shell
export AWS_REGION=localstack
export AWS_KINESIS_ENDPOINT=http://localhost:4568
export AWS_SECRET_KEY=1231
export AWS_ACCESS_KEY_ID=999999
export AWS_KINESIS_STREAM=deposit
```

### Run the router in go

```shell
go run main.go
```