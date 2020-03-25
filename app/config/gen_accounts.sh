#!/usr/bin/env bash
for i in $(seq 2000); do geth --datadir . account new --password password ; done
