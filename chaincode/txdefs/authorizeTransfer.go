package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

var AuthorizeTransfer = tx.Transaction{
	Tag:         "authorizeTransfer",
	Label:       "Authorize Transfer",
	Description: "Authorize a transfer of an asset",
	Method:      "PUT",
	Callers:     []string{`$org2MSP`, "orgMSP"},

	Args: []tx.Argument{
		{
			Tag:         "car",
			Label:       "Car",
			Description: "Car",
			DataType:    "->car",
			Required:    true,
		},
	},

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		carKey, ok := req["car"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter car must be an asset")
		}

		carAsset, err := carKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		carMap := (map[string]interface{})(*carAsset)

		// Check if the transfer is authorized
		if carMap["transfer"].(string) != "requested" {
			return nil, errors.WrapError(nil, "Transfer is not requested")
		}

		updatedOwnerKey := make(map[string]interface{})
		updatedOwnerKey["@assetType"] = "person"
		updatedOwnerKey["@key"] = carMap["transferTo"].(map[string]interface{})["@key"]

		carMap["owner"] = updatedOwnerKey
		carMap["transfer"] = "completed"
		carMap["transferTo"] = nil

		carMap, err = carAsset.Update(stub, carMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update asset")
		}

		carJSON, nerr := json.Marshal(carMap)
		if nerr != nil {
			return nil, errors.WrapError(nerr, "failed to marshal asset")
		}

		return carJSON, nil
	},
}
