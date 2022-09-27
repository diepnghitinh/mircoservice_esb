FROM golang:1.12-alpine as builder

WORKDIR $GOPATH/src/esb-worker
COPY . .

RUN apk add --no-cache make curl git build-base

ENV GO111MODULE=on
ENV GIN_MODE=release

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /api api/api.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /consumer cmd/consumer.go

FROM golang:1.12-alpine

EXPOSE 8080

COPY --from=builder /consumer .
COPY --from=builder /api .

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN ["chmod", "+x", "/docker-entrypoint.sh"]