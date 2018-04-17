package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var t test_struct
	for key, _ := range req.Form {
		log.Println(key)
		//LOG: {"test": "that"}
		err := json.Unmarshal([]byte(key), &t)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
