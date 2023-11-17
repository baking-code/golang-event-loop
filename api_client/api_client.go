package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type jsonObject map[string]interface{}
type jsonArray []interface{}

var client *http.Client

func getClient() *http.Client {
	if client == nil {
		client = &http.Client{}
	}

	return client
}

func handleResponse(resp *http.Response, err error) (interface{}, error) {

	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body) // response body is []byte

		var result interface{}
		fmt.Println(string(body))
		if json.Unmarshal(body, &result) == nil {
			return result, err
		} else {
			return string(body), err
		}
	} else {
		return nil, err
	}
}

func Get(url string) (interface{}, error) {
	client := getClient()
	return handleResponse(client.Get(url))
}
