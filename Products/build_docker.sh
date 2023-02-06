#!/bin/bash
docker build -t products:1.0.0 .
docker tag products:1.0.0 europe-central2-docker.pkg.dev/promising-node-365908/test/products:1.0.0
docker push europe-central2-docker.pkg.dev/promising-node-365908/test/products:1.0.0