#!/bin/bash
docker build -t go-base:1.0.0 .
docker tag go-base:1.0.0 europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/go-base:1.0.0
docker push europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/go-base:1.0.0