package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	file, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(file)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
}
