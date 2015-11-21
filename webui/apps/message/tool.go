package message

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func GetBackendApi(method string, url string, b []byte) []byte {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Sid", sid)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
