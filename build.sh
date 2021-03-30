#!/bin/bash
set -xe
LDFLAGS="-s -w"
rm -rf output
mkdir -p output
TMP='/tmp'
VERSION='1.0.0'

function buildProduct() {
    echo "build go-ops beging"
    git pull
    go build -ldflags "$LDFLAGS" -o go-ops cmd/server/main.go
    mv kjcloud output
    systemctl restart go-ops
}

function _help() {
  echo "Welcome to go-ops build system"
}
function main() {

    _help
    select ch in "go-ops";
    do
        case $ch in
        "go-ops" ) buildProduct
              break;;

        *)  echo "please choose a number"
              break;;
        esac
    done
}
main