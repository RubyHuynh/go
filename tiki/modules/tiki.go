package tiki

import (
	"hthngoc/logging"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"
)

func init() {
}
type extractedData struct {
	Url string `json:"url_key"`
	Price int `json:"price"`
}

func init() {
}

func Access() string {
	return "hmm"
}

func Get(sku int) extractedData {
	_log := log.GetInstance()
	resp, err := http.Get("https://tiki.vn/api/v2/products/2076813")
	if err != nil {
		_log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_log.Printf("%v\n", err)
		//log.Fatalln(err)
	}
	data := extractedData{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		//log.Fatal(jsonErr)
	}
	/*Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	fmt.Printf("%v", sb)
	*/
	fmt.Println(data)
	_log.Printf("xxx %v\n", data)
	return data
}
