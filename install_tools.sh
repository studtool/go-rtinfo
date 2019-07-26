#!/usr/bin/env bash

dir="$1"

if ! which "${dir}/golangci-lint" &> /dev/null ; then
    GOBIN="${dir}" go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1 && go mod tidy
fi
