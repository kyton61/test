package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

var tmpl = template.Must(template.New("FortuneTeller").
	Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です！</body></html>"))

type Result struct {
	Name    string
	Omikuji string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func assert(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fortuneTeller() string {
	var result string
	i := rand.Intn(6)
	switch i {
	case 0:
		result = "大吉"
	case 1, 2:
		result = "吉"
	case 3, 4:
		result = "凶"
	default:
		result = "大凶"
	}
	return result
}
func handler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Name:    r.FormValue("p"),
		Omikuji: fortuneTeller(),
	}
	tmpl.Execute(w, result)
	/*
		w.Header().Set("Content-type", "application/json; charset=utf-8")
		v := struct {
			Msg string `json:"msg"`
		}{
			Msg: "Hello",
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Println("Error:", err)
		}
	*/
}

func main() {
	/*
		rand.Seed(time.Now().UnixNano())
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	*/
	var url string
	for {
		fmt.Print("url:")
		fmt.Scan(&url)
		resp, err := http.Get(url)
		assert(err)
		defer resp.Body.Close()
		fmt.Println("Request")
		fmt.Println(resp.Request)
		fmt.Println("header")
		fmt.Println(resp.Header)
		fmt.Println("body")
		fmt.Println(resp.Body)

		fmt.Println(resp)
	}
}
