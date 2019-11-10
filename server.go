package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	word "github.com/fachryansyah/test-golang-solusi-menghitung-huruf"
)

type Response struct {
	Status  int
	Message string
	Data    interface{}
}

func main() {

	http.HandleFunc("/solusi-satu", solutionOne)

	fmt.Println("Webservice running on port : 1337")
	http.ListenAndServe(":1337", nil)
}

func solutionOne(res http.ResponseWriter, req *http.Request) {

	var response Response
	// var data interface{}
	response.Status = 200
	response.Message = "OKE"

	if req.Method == "POST" {
		input := req.FormValue("word")
		vocal, consonant := word.CountWord(input)

		data := map[string]interface{}{
			"HurufHidup": vocal,
			"HurufMati":  consonant,
		}
		response.Data = data
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}
