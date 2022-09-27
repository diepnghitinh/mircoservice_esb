#!/usr/bin/env bash

protoc --proto_path=../pb/feed --go_out=plugins=grpc:../ feed_notification.proto
protoc --proto_path=../pb/feed --go_out=plugins=grpc:../ feed_test.proto
protoc --proto_path=../pb/feed --go_out=plugins=grpc:../ feed.proto