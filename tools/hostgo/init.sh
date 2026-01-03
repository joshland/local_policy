#!/bin/bash

go mod init host-tool
go mod tidy
go build -o host main.go
