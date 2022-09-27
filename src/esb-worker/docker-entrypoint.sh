#!/bin/sh
set -e

if [ "$1" = 'api' ]; then
./api
fi


if [ "$1" != 'api' ]; then
./consumer -service "$1" worker
fi


exec "$@"
