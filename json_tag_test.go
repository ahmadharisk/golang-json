package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Tahu",
		ImageUrl: "http://contoh.com/image.png",
	}

	productJson, _ := json.Marshal(product)
	fmt.Println(string(productJson))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Tahu","image_url":"http://contoh.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	_ = json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
	fmt.Println(product.ImageUrl)
}
