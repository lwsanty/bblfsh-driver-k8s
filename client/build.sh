#!/usr/bin/env bash

go test -run TestK8s -o test
docker build -t client .
