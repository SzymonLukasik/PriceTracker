#!/bin/bash
docker build -t users:1.0.0 .
docker tag users:1.0.0 europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/users:1.0.0
docker push europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/users:1.0.0