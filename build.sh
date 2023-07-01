#!/usr/bin/env bash
go build -a -ldflags '-s -w' \
    -gcflags="all=-trimpath=${PWD}" \
    -asmflags="all=-trimpath=${PWD}"
