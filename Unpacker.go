package main

import "encoding/json"

type Unpacker struct {

}

func (unpacker *Unpacker) UnpackJSON(productData []byte) map[string]interface{} {
	var rawData interface{}
	err := json.Unmarshal(productData, &rawData)
	if err != nil {
		panic(err)
	}

	return rawData.(map[string]interface{})
}
