package main

import (
	"net/http"
	"io"
	//"strings"
	"io/ioutil"
	"fmt"
)


func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res, err := http.Get("http://localhost:8080/")
		if err != nil {
			fmt.Println(err)
		}
		data, _ := ioutil.ReadAll( res.Body )
		res.Body.Close()
		err = ioutil.WriteFile("p.html", data, 0666)
		if err != nil {panic(err) }	
		res.Body.Close()
	} 
	if req.Method == "POST" {
		io.WriteString(w, "successful post")
	}
	if req.Method == "PUT" {
		io.WriteString(w, "successful put")
	}
	if req.Method == "DELETE" {
		io.WriteString(w, "successful delete")
	}
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are on page1\n")
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/page1", Handler2)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
