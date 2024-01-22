package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func main() {
	data :=
		`
{
  "type": "1",
  "info": {
    "a": 1,
    "b": 1,
    "info2": {
      "c": 1,
      "bb": "hello world"
    },
	"ll":[1,2,3],
	"ll2":["4","5","6"]
  }
}
    `
	var v map[string]interface{}
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		log.Fatal(err)
	}

	fmt.Println(GetVal(v, []string{"type"}))
	x := GetVal(v, []string{"type"}).(string)
	fmt.Println(x)
	ll := InterfaceSliceToIntSlice(GetVal(v, []string{"info", "ll"}).([]interface{}))
	ll2 := InterfaceSliceToFloatSlice(GetVal(v, []string{"info", "ll"}).([]interface{}))
	ll3 := InterfaceSliceToStringSlice(GetVal(v, []string{"info", "ll"}).([]interface{}))
	fmt.Println(ll)
	fmt.Println(ll2)
	fmt.Println(ll3)

	ll4 := InterfaceSliceToIntSlice(GetVal(v, []string{"info", "ll2"}).([]interface{}))
	fmt.Println(ll4)
	//x := GetVal(v, []string{"type"})
	fmt.Println(GetVal(v, []string{"info", "info2", "bb"}))
}

func GetVal(data map[string]interface{}, keys []string) interface{} {
	for _, key := range keys {
		val, ok := data[key]
		if !ok {
			return nil
		}

		switch v := val.(type) {
		case map[string]interface{}:
			data = v
		default:
			return v
		}
	}
	return nil
}

func InterfaceSliceToIntSlice(slice []interface{}) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		switch v.(type) {
		case string:
			x, _ := strconv.Atoi(v.(string))
			result[i] = x
		case float64:
			x, _ := v.(float64)
			result[i] = int(x)
		}
	}
	return result
}

func InterfaceSliceToFloatSlice(slice []interface{}) []float64 {
	result := make([]float64, len(slice))
	for i, v := range slice {
		result[i] = v.(float64)
	}
	return result
}

func InterfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}
