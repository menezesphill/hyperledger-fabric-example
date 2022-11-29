package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Request a transfer of an asset
// POST Method
var RequestTransfer = tx.Transaction{
	Tag:         "requestTransfer",
	Label:       "Request Transfer",
	Description: "Request a transfer of an asset",
	Method:      "PUT",
	Callers:     []string{`$org1MSP`, "orgMSP"},

	Args: []tx.Argument{
		{
			Tag:         "car",
			Label:       "Car",
			Description: "Car",
			DataType:    "->car",
			Required:    true,
		},
		{
			Tag:         "owner",
			Label:       "Owner",
			Description: "New owner of the car",
			DataType:    "->person",
			Required:    true,
		},
	},

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		carKey, ok := req["car"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter car must be an asset")
		}
		ownerKey, ok := req["owner"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter owner must be an asset")
		}

		carAsset, err := carKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		carMap := (map[string]interface{})(*carAsset)

		ownerAsset, err := ownerKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		ownerMap := (map[string]interface{})(*ownerAsset)

		updatedOwnerKey := make(map[string]interface{})
		updatedOwnerKey["@assetType"] = "person"
		updatedOwnerKey["@key"] = ownerMap["@key"]

		carMap["transferTo"] = updatedOwnerKey
		carMap["transfer"] = "requested"

		carMap, err = carAsset.Update(stub, carMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update asset")
		}

		carJSON, nerr := json.Marshal(carMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return carJSON, nil
	},
}
