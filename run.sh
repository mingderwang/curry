#!/bin/bash
go generate
go build -o onion
./onion migratedb
./onion serve
