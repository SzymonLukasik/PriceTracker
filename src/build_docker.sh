#!/bin/bash
docker build -t scraper:1.0.0 .
docker tag scraper:1.0.0 europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/scraper:1.0.0
docker push europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/scraper:1.0.0