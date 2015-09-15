#!/bin/bash
set -o xtrace
go generate
go build -o onion
./onion migratedb
./onion serve
