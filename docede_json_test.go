package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		LastName:   "Smith",
		MiddleName: "John Smith",
		Age:        18,
		IsMarried:  false,
	}

	result, _ := json.Marshal(customer)

	decodeJson := &Customer{}
	err := json.Unmarshal(result, decodeJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(decodeJson)
	fmt.Println(decodeJson.FirstName)
	fmt.Println(decodeJson.LastName)
	fmt.Println(decodeJson.MiddleName)
	fmt.Println(decodeJson.Age)
	fmt.Println(decodeJson.IsMarried)
}
