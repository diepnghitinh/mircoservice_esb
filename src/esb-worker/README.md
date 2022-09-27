# Enterprise service bus

#### What support for Shared ?

1. Rabbitmq Broker: auto reconnect, Log error queue, retry message command.
2. Postgres: ORM.
3. Restful: Retry support.
4. Docker support.

#### How to start a Consumer ? 

Using file config
```
go run cmd/consumer.go -c cmd/identity.language/config.yml -service mail.handler worker
```

Or using enviroment variable for Docker
```
go run cmd/consumer.go -service mail.handler worker
```

####RUN TEST
```
go test ./test -v -cover
```

####GRPC GENERATE
```
cd deployments
./grpc.sh
```