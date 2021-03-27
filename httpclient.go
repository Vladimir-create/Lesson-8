package main
import (
	"net/http"
	"io/ioutil"
	//"bytes"
	"strings"
	"io"
	"fmt"
)
func main() {
	client := &http.Client{}
	
	requestBody := strings.NewReader(`
		{
			"action":"create",
			"object":"Teacher",
			"data":
			{
				"id":"s001",
				"subject":"Math",
				"salary":2345,
				"classroom":
				[
					"CL-001",
					"CL-002",
					"CL-005"
				],
				"person":
				{
					"name":"Ivan",
					"surname":"Popov",
					"personalCode":"123422-43235"
				}
			}
		}
	`)
	//POST
	req_post, err := http.NewRequest("POST", "http://localhost:8080/", &requestBody)
	if err != nil {
        fmt.Println(err)
        return
    }
    resp_post, err := client.Do(req_post)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp_post.Body.Close()
    
    data_post, err := io.ReadAll(resp_post.Body)
    
	err = ioutil.WriteFile("p.html", data_post, 0666)
	if err != nil {panic(err) }
		
	//DELETE
	req_delete, err := http.NewRequest("DELETE", "http://localhost:8080/", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    resp_delete, err := client.Do(req_delete)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp_delete.Body.Close()
    
    data_delete, err := io.ReadAll(resp_delete.Body)
    
	err = ioutil.WriteFile("p.html", data_delete, 0666)
	if err != nil {panic(err) }
    
    	//READ
	req_read, err := http.NewRequest("GET", "http://localhost:8080/", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    resp_read, err := client.Do(req_read)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp_read.Body.Close()
    
    data_read, err := io.ReadAll(resp_read.Body)
    
	err = ioutil.WriteFile("p.html", data_read, 0666)
	if err != nil {panic(err) }
    
    //Update
    	update := strings.NewReader(`
		{
			"action":"update",
			"object":"Teacher",
			"data":
			{
				"id":"s001",
				"subject":"Math",
				"salary":2345,
				"classroom":
				[
					"CL-001",
					"CL-002",
					"CL-005"
				],
				"person":
				{
					"name":"Ivfan",
					"surname":"Podfpov",
					"personalCode":"12345422-43235"
				}
			}
		}
	`)
    req_update, err := http.NewRequest("PUT", "http://localhost:8080/", &update)
	if err != nil {
        fmt.Println(err)
        return
    }
    resp_update, err := client.Do(req_update)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp_update.Body.Close()
    
    data_update, err := ioutil.ReadAll(resp_update.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
	
	err = ioutil.WriteFile("p.html", data_update, 0666)
	if err != nil {panic(err) }	
}
