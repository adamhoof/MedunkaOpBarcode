package main

import (
	"io/ioutil"
	"net/http"
)

type APIHandler struct {

}

func (apiHandler *APIHandler) RequestProductData(barcodeOutput string) []byte {

	resp, err := http.Get("https://medunka.cz/api/product?sku="+ barcodeOutput)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return data
}
