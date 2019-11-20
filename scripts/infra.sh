#!/usr/bin/env bash

set -e

function infra_up() {
    cd $GOPATH/src/fizzbuzz
    docker-compose up -d
}

function infra_down() {
    cd $GOPATH/src/fizzbuzz 
    docker-compose stop
}

case $1 in
    up) infra_up ;;
    down) infra_down ;;
esac