package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

var IsCertReceptor = tx.Transaction{
	Tag:         "isCertReceptor",
	Label:       "Is Car Certificate Receptor",
	Description: "Return car transfer intentions towards this person",
	Method:      "GET",
	Callers:     []string{"org2MSP", "org3MSP", "orgMSP"},

	Args: []tx.Argument{
		{
			Tag:      "receptor",
			Label:    "Receptor",
			DataType: "->person",
			Required: true,
		},
		{
			Tag:         "limit",
			Label:       "Limit",
			Description: "Limit the number of results",
			DataType:    "number",
		},
	},

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		receptor, _ := req["receptor"]
		limit, hasLimit := req["limit"].(float64)

		if hasLimit && limit <= 0 {
			return nil, errors.NewCCError("Limit must be greater than 0", 400)
		}

		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType": "car",
				"transferTo": receptor,
			},
		}

		if hasLimit {
			query["limit"] = limit
		}

		var err error
		response, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error searching for car's receptor", 500)
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error marshalling response", 500)
		}

		return responseJSON, nil
	},
}
