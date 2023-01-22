#!/bin/bash
docker build -t frontend:1.0.1 .
docker tag frontend:1.0.1 europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/frontend:1.0.1
docker push europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/frontend:1.0.1