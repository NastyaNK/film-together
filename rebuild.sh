#!/bin/bash
set -e

echo " - Building Go binary..."
rm -f ./dist/myapp
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/myapp ./cmd/film

echo " - Build complete: ./myapp"
