package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Peppo",
		LastName:   "Bond",
		MiddleName: "Jane",
		Hobbies:    []string{"game", "coding", "reading"},
	}

	jsonEncode, _ := json.Marshal(customer)
	fmt.Println(string(jsonEncode))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Peppo","LastName":"Bond","MiddleName":"Jane","Age":0,"IsMarried":false,"Hobbies":["game","coding","reading"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	_ = json.Unmarshal(jsonBytes, customer)
	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Hobbies[1])
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Peppo",
		LastName:   "Bond",
		MiddleName: "Jane",
		Hobbies:    []string{"game", "coding", "reading"},
		Addresses: []Address{
			{
				Street: "ssa",
				City:   "meli",
				County: "clube",
			},
			{
				Street: "for",
				City:   "thank",
				County: "led",
			},
		},
	}

	jsonEncode, _ := json.Marshal(customer)
	fmt.Println(string(jsonEncode))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Peppo","LastName":"Bond","MiddleName":"Jane","Age":0,"IsMarried":false,"Hobbies":["game","coding","reading"],"Addresses":[{"Street":"ssa","City":"meli","County":"clube"},{"Street":"for","City":"thank","County":"led"}]}`
	jsonBytes := []byte(jsonString)

	customer2 := &Customer{}
	_ = json.Unmarshal(jsonBytes, customer2)
	fmt.Println(customer2)
	fmt.Println(customer2.Hobbies)
	fmt.Println(customer2.Hobbies[1])
	fmt.Println(customer2.Addresses[0].City)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"ssa","City":"meli","County":"clube"},{"Street":"for","City":"thank","County":"led"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	_ = json.Unmarshal(jsonBytes, addresses)
	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	Addresses := []Address{
		{
			Street: "ssa",
			City:   "meli",
			County: "clube",
		},
		{
			Street: "for",
			City:   "thank",
			County: "led",
		},
	}

	jsonEncode, _ := json.Marshal(Addresses)
	fmt.Println(string(jsonEncode))
}
