package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	LastName   string
	MiddleName string
	Age        int
	IsMarried  bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street string
	City   string
	County string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		LastName:   "Smith",
		MiddleName: "John Smith",
		Age:        18,
		IsMarried:  false,
	}

	result, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
