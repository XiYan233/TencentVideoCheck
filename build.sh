#!/usr/bin/env bash

. /etc/profile

# Linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-s -w -buildid=" -o TencentVideoCheck-linux-arm64 main.go
# Linux amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w -buildid=" -o TencentVideoCheck-linux-amd64 main.go
# Windows amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-s -w -buildid=" -o TencentVideoCheck-windows-amd64 main.go
# macOS amd64
#CGO_ENABLED=0 GOOS=linux GOARCH=darwin go build -trimpath -ldflags="-s -w -buildid=" -o TencentVideoCheck-darwin-amd64 main.go