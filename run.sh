#!/usr/bin/env bash

# scraping target on localhost:5000, using https://github.com/prometheus/client_ruby/tree/master/examples/rack
pushd exporter
bundle exec unicorn -c ./unicorn.conf
popd

# prometheus on localhost:9090 (build using `promu build`)
./prometheus/prometheus \
    --config.file ./prometheus.yml \
    --storage.tsdb.path=./data \
    --web.enable-lifecycle

# alertmanager on localhost:9093 (build using `promu build`)
./alertmanager/alertmanager \
    --config.file /home/mtayer/dev/repo/scripts/data/alertmanager.yml \
    --storage.path ./data

# SMTP on localhost:25
sudo java -jar FakeSMTP/target/fakeSMTP-2.1-SNAPSHOT.jar

