package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"a":"apple","b":"ball","c":"cat","d":"dog","e":"elephant"}`
	x := map[string]string{}

	json.Unmarshal([]byte(jsonStr), &x)
	fmt.Println(x)
}
