package main

import (
	"bytes"
	"fmt"
	// "html"
	"net/http"
	"io"
	"io/ioutil"
	//"time"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/", fs)

	r := mux.NewRouter()

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/api{url:.+}", ApiHandler)

	// r.HandleFunc("/{path:.+}", DefaultHandler)
	r.HandleFunc("/{path:.+}", IndexHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8001", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)
	vars := mux.Vars(r)

	/*
	fmt.Println("time sleep 3s for test ...")
	time.Sleep(3 * time.Second)
	*/

	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
			fmt.Println("read request body failed: ", err)
			return
	}
	fmt.Println("req_body = ", string(req_body))

	api_url := r.URL
	api_url.Path = "http://127.0.0.1:8080" + vars["url"]

	req, err := http.NewRequest(r.Method, api_url.String(), bytes.NewBuffer(req_body))
	req.Header.Set("Sid", r.Header.Get("Sid"))
	// req.Header.Set("Content-Type", "application/json")
	fmt.Println("==>",r.Header.Get("Content-Type"))
	req.Header.Set("Content-Type", r.Header.Get("Content-Type"))
	fmt.Println("req.Header:", req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("{\"error\", %s}", err.Error()))
		return
	}
	defer resp.Body.Close()

	// resp.StatusCode
	// resp.Status
	//这句话必须放在 w.WriteHeader 之前！
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode)


	io.Copy(w, resp.Body)
}
