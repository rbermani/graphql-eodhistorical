package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"json_query/server_type"
	"log"
	"net/http"
	"os"
	"time"
)

func GetFundamentals(ticker string) (*server_type.EquityFundamentals, error) {
	const eod_fundamental_api string = "https://eodhistoricaldata.com/api/fundamentals/AAPL.US?api_token="
	const eod_api_key string = "OeAFFmMliFG5orCUuwAKQ8l4WWFQ67YX"
	var http_client = &http.Client{Timeout: 10 * time.Second}
	result, err := http_client.Get(eod_fundamental_api + eod_api_key)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer result.Body.Close()
	fundamental_data := new(server_type.EquityFundamentals)

	err = json.NewDecoder(result.Body).Decode(&fundamental_data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return fundamental_data, nil
}

func main() {
	//fundptr, err := GetFundamentals("AAPL")
	//if err != nil {
	//	return
	//}
	jsonFile, err := os.Open("sample.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var fundamental_data server_type.EquityFundamentals
	if err = json.Unmarshal([]byte(byteValue), &fundamental_data); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", fundamental_data)
}
