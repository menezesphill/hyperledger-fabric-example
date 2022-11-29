package main

import (
	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/hyperledger-chaincode-demo/chaincode/assettypes"
)

var assetTypeList = []assets.AssetType{
	assettypes.Person,
	assettypes.Secret,
	assettypes.Car,
}
