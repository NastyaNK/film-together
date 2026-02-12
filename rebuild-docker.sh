#!/bin/bash
set -e

./rebuild.sh

echo " - Building Docker image..."
docker build -t film-together:local .