package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
)

type (
	Action struct {
		Action string `json:"action"`
		ObjName string `json:"object"`
	}
	Teacher struct {
		ID string  `json:"id"`
		Salary float64 `json:"salary"`
		Subject string `json:"subject"`
		Classroom []string `json:"classroom"`
		Person struct {
			Name string `json:"name"`
			Surname string `json:"surname"`
			PersonalCode string `json:"personalCode"`
		} `json:"person"`
	}
	Teachermutex struct {
		Ch chan int
	}
	UpdateTeacher struct {
		T Teacher `json:"data"`
	}
	CreateTeacher struct {
		T Teacher `json:"data"`
	}
	DeleteTeacher struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	ReadTeacher struct {
		T Teacher
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.Write()
	} 
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else {
		w.WriteHeader(405)
	}
	
}

func Handler2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are on page1")
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/page1", Handler2)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
