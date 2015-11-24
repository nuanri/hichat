package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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

type BackenApi struct {
	c   *gin.Context
	Sid string
}

func (b *BackenApi) Get(obj interface{}, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Sid", b.Sid)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//fmt.Printf("%#v\n", body)
	//fmt.Printf("[BACKEND API] GET %s\n%s\n", url, string(body))

	return json.Unmarshal(body, obj)
}

func GetBackendApi2(c *gin.Context) *BackenApi {
	sid, err := c.Get("Sid")

	if err {
		return &BackenApi{
			c:   c,
			Sid: sid.(string),
		}
	} else {
		return &BackenApi{
			c: c,
		}
	}
}
