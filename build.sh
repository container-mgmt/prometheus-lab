#!/usr/bin/env bash

# TBD build, currently includes all needed binaries

yum install git-all golang ruby -y
gem install bundler

# Build metric exporter
pushd exporter
bundle intsall
popd