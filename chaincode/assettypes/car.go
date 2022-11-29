package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a car
var Car = assets.AssetType{
	Tag:         "car",
	Label:       "Car",
	Description: "Car",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "plate",
			Label:    "Car Plate",
			DataType: "string",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			Tag:      "owner",
			Label:    "Owner",
			DataType: "->person",
		},
		{
			Tag:      "transfer",
			Label:    "Transfer Request",
			DataType: "string",
		},
		{
			Tag:      "transferTo",
			Label:    "Transfer Recipient",
			DataType: "->person",
		},
	},
}
