package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//生成 Json
func GetJson(m interface{}) []byte {

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Marshal 出错", err)
	}
	return b
}

//解析 Json
func ParseMsg(data []byte, m interface{}) error {

	err := json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("json 解析出错", err)
		return err
	}
	return nil
}

func GetBackendApi(sid string, method string, url string, b []byte) []byte {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Sid", sid)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
