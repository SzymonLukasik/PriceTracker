#!/bin/bash
docker build -t frontend:1.0.1 .
docker tag frontend:1.0.1 europe-central2-docker.pkg.dev/promising-node-365908/test/frontend:1.0.1
docker push europe-central2-docker.pkg.dev/promising-node-365908/test/frontend:1.0.1