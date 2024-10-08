#!/bin/bash

# PREFIX="ALPHA"
TAGS="" # "server"


echo "Building amd64..."
if [ "$PREFIX" != "" ]; then
    env CGO_ENABLED=0 go build -tags "$TAGS" -ldflags="-s -w" -o "bin/$PREFIX-mavbake-linux-amd64" 
else
    env CGO_ENABLED=0 go build -tags "$TAGS" -ldflags="-s -w" -o "bin/mavbake-linux-amd64" 
fi
echo "Building arm64..."
if [ "$PREFIX" != "" ]; then
    env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags "$TAGS" -ldflags="-s -w" -o "bin/$PREFIX-mavbake-linux-arm64";
else
    env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags "$TAGS" -ldflags="-s -w" -o "bin/mavbake-linux-arm64";
fi
echo "Building docs..."
go run -tags "$TAGS" ./docs/build/
