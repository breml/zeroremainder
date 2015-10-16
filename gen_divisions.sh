#!/bin/bash
#
# Generate division implementations based on skelleton in division_uint.go
# Also removes the build tag and the go:generate instructions
#
# Usage:
#   go generate *.go
#
cat $GOFILE | sed "s/\([uU]\)int8/\1int${1}/g" | grep -v "// +build go:generate" | grep -v "go:generate" | gofmt > ${GOFILE%\.go}${1}.go
