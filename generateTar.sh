#!/usr/bin/env bash

# Make sure go mod is up to date
cd chaincode && go mod vendor && cd ..

# Compress file without rest-server (GoFabric will use the standard CC API)
tar -czf hyperledger-chaincode-demo.tar.gz chaincode

# Compress file with rest-server (GoFabric will use the one provided)
# tar -c --exclude=node_modules -zf hyperledger-chaincode-demo.tar.gz chaincode rest-server