#!/usr/bin/env bash

set -e

cd $GOPATH

go build -o bin/fizzbuzz lbc.dev/fizzbuzz