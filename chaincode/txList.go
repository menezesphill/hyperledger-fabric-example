package main

import (
	txdefs "github.com/goledgerdev/hyperledger-chaincode-demo/chaincode/txdefs"

	tx "github.com/goledgerdev/cc-tools/transactions"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.RequestTransfer,
	txdefs.AuthorizeTransfer,
	txdefs.GetCarByOwner,
	txdefs.IsCertReceptor,
	txdefs.AcceptTransfer,
}
