package apiclient

import (
	"bytes"
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

func Post(url string, body interface{}) (interface{}, error) {
	client := getClient()
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return handleResponse(client.Post(url, "application/json", bytes.NewReader(b)))
}
