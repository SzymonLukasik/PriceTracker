#!/bin/bash
docker build -t scraper:1.0.0 .
docker tag scraper:1.0.0 europe-central2-docker.pkg.dev/promising-node-365908/test/scraper:1.0.0
docker push europe-central2-docker.pkg.dev/promising-node-365908/test/scraper:1.0.0