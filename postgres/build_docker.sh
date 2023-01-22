#!/bin/bash
docker build -t my_postgres:1.0.0 .
docker tag my_postgres:1.0.0 europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/my_postgres:1.0.0
docker push europe-central2-docker.pkg.dev/mim-navigation/aaaaaa/my_postgres:1.0.0