#!/bin/bash

export GO111MODULE=on
export GOCACHE=/go/cache
mkdir -p /go/cache
cd src
cp /tmp/go.mod .
go install
