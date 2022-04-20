#!/bin/sh

NAME_LINUX_AMD64="server.linux.amd64"
NAME_LINUX_ARM64="server.linux.arm64"

NAME_DARWIN_AMD64="server.darwin.amd64"
NAME_DARWIN_ARM64="server.darwin.arm64"

NAME_WINDOWS_AMD64="server.windows.amd64"
NAME_WINDOWS_ARM64="server.windows.arm64"

export GO111MODULE=auto

rm -f $NAME_LINUX_AMD64
export GOOS=linux
export GOARCH=amd64
go build -o $NAME_LINUX_AMD64

rm -f $NAME_LINUX_ARM64
export GOOS=linux
export GOARCH=arm64
go build -o $NAME_LINUX_ARM64

rm -f $NAME_DARWIN_AMD64
export GOOS=darwin
export GOARCH=amd64
go build -o $NAME_DARWIN_AMD64

rm -f $NAME_DARWIN_ARM64
export GOOS=darwin
export GOARCH=arm64
go build -o $NAME_DARWIN_ARM64

rm -f $NAME_WINDOWS_AMD64
export GOOS=windows
export GOARCH=amd64
go build -o $NAME_WINDOWS_AMD64

rm -f $NAME_WINDOWS_ARM64
export GOOS=windows
export GOARCH=arm64
go build -o $NAME_WINDOWS_ARM64