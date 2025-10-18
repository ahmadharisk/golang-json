package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("peppo")
	logJson(100)
	logJson(true)
	logJson([]string{"peppo", "silpiana"})
	logJson(map[string]interface{}{
		"name": "silpiana",
		"age":  7,
	})
}
