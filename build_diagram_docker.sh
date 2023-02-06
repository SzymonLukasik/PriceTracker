#!/bin/bash
docker build -t diags:1.0.0 .
docker tag diags:1.0.0 europe-central2-docker.pkg.dev/promising-node-365908/test/diags:1.0.0
docker push europe-central2-docker.pkg.dev/promising-node-365908/test/diags:1.0.0