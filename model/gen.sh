#!/bin/bash

set -e

# generate the protobufs
protoc --go_out=plugins=grpc:./ \
        -I./ ./kv.proto
        
protoc  --cuckoo_out=. kv.proto