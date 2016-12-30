package testingutils

import (
	"fmt"
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
