package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		LastName:   "Doe",
		MiddleName: "Jane",
	}

	writer, _ := os.Create("customer2.json")
	encoder := json.NewEncoder(writer)

	err := encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
}
