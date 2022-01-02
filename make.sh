#!/usr/bin/env bash

# Author: René Kray <rene@kray.info>
# Date:   2022-01-02

GOARCH=amd64

case $1 in
    build)
        #GOOS=windows go build
        GOOS=linux go build
        #GOOS=darwin go build

        BASENAME=$(basename $(awk '/^module/{print $2}' go.mod))
        strip $BASENAME
        ;;
    test)
        ;;
    clean)
        go clean
        ;;
    *)
        echo "$0 build - complie the current project"
        echo "$0 clean - cleanup the current project"
        ;;
esac
