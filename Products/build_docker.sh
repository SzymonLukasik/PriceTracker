#!/bin/bash
docker build -t products:1.0.0 .
docker tag products:1.0.0 europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/products:1.0.0
docker push europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/products:1.0.0