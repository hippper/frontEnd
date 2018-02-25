#!/bin/bash

rm -rf output/bin
rm -rf output/conf
rm -rf output/log
rm -rf frontEnd.zip

mkdir -p output/bin
mkdir -p output/conf
mkdir -p output/log
mkdir -p output/status/frontEnd
make

cp load.sh output/load.sh

cp bin/supervise.frontEnd output/bin/supervise.frontEnd
cp bin/frontEnd output/bin/frontEnd || { echo error ; exit 1 ; }
cp conf/frontEnd.toml output/conf/frontEnd.toml
cp conf/log.json output/conf/log.json
cp -r static output/
cp -r templates output/