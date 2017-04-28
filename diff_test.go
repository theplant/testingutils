package testingutils

import (
	"fmt"
	"testing"
)

func ExamplePrettyJsonDiff() {
	type Company struct {
		Name string
	}
	type People struct {
		Name    string
		Age     int
		Company Company
	}

	p1 := People{
		Name: "Felix",
		Age:  20,
		Company: Company{
			Name: "The Plant",
		},
	}
	p2 := People{
		Name: "Tom",
		Age:  21,
		Company: Company{
			Name: "Microsoft",
		},
	}

	fmt.Println(PrettyJsonDiff(p1, p2))
	//Output:
	// --- Expected
	// +++ Actual
	// @@ -1,7 +1,7 @@
	//  {
	// -	"Name": "Felix",
	// -	"Age": 20,
	// +	"Name": "Tom",
	// +	"Age": 21,
	//  	"Company": {
	// -		"Name": "The Plant"
	// +		"Name": "Microsoft"
	//  	}
	//  }

}

func ExamplePrettyJsonDiff_JSONRawMessage(t *testing.T) {
	str1 := `{
"Active": true,
	"Addresses": [],	"Age": 0,
		"Avatar": "",
"Company": "",
		"CreditCard": {
			"Number": "411111111111",
			"Issuer": "VISA"
		},
		"ID": 1,
	"Languages": null,
		"Name": "jinzhu",
		"Profile": {
			"ID": 0,
			"Name": "jinzhu",
			"Phone": {
				"ID": 0,
				"Num": "110"
			},
			"Sex": "male"
		},
		"RegisteredAt": "2017-01-01 00:00",
		"Role": ""
	}`

	str2 := `{
		"Active": true,
		"Addresses": [],
		"Age": 0,
		"Avatar": "",
		"Company": "",
		"ID": 1,
		"Languages": null,
		"Name": "jinzhu",
		"Profile": {
			"ID": 0,
			"Name": "jinzhu",
			"Phone": {
				"ID": 0,
				"Num": "110"
			},
			"Sex": "male"
		},
		"RegisteredAt": "2017-01-01 00:00",
		"Role": ""
	}`

	fmt.Println(PrettyJsonDiff(str1, str2))
	// Output
	//  	"Age": 0,
	//  	"Avatar": "",
	//  	"Company": "",
	// -	"CreditCard": {
	// -		"Number": "411111111111",
	// -		"Issuer": "VISA"
	// -	},
	//  	"ID": 1,
	//  	"Languages": null,
	//  	"Name": "jinzhu",
}
