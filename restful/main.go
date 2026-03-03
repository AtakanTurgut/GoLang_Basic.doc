// PS ...\restfull> go run main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://jsonplaceholder.typicode.com/todos/

type Todo struct { // Alt Gr + , => ` `
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// GET //
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close() // defer => function bitince kapanır

	bodyBytes, _ := ioutil.ReadAll(response.Body) // response.Body => io.ReadCloser

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo) // Unmarshal => JSON verisini Go struct'ına dönüştürür

	fmt.Printf("Todo: %+v \n", todo) // {UserId:1 Id:1 Title:delectus aut autem Completed:false}

	// POST	//
	todoPost := Todo{
		UserId:    1,
		Id:        201,
		Title:     "dicta sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Completed: false,
	}

	jsonTodo, err := json.Marshal(todoPost) // Marshal => Go struct'ını JSON verisine dönüştürür

	responsePost, errPost := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonTodo))
	// bytes.NewBuffer(jsonTodo) => jsonTodo'yu io.Reader'a dönüştürür

	if errPost != nil {
		fmt.Println(errPost)
	}

	defer responsePost.Body.Close()

	/// POST response'unu okuyalım

	bodyBytesPost, _ := ioutil.ReadAll(responsePost.Body)

	bodyStringPost := string(bodyBytesPost)
	fmt.Println(bodyStringPost)

	var todoPostResponse Todo
	json.Unmarshal(bodyBytesPost, &todoPostResponse) // Unmarshal => JSON verisini Go struct'ına dönüştürür

	fmt.Printf("Todo: %+v \n", todoPostResponse)

	//TodoPost()
}

func TodoPost() {
	// POST-2 //
	todoPost := Todo{
		UserId:    1,
		Id:        202,
		Title:     "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Completed: false,
	}

	jsonTodo, _ := json.Marshal(todoPost) // Marshal => Go struct'ını JSON verisine dönüştürür

	responsePost, errPost := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonTodo))
	// bytes.NewBuffer(jsonTodo) => jsonTodo'yu io.Reader'a dönüştürür

	if errPost != nil {
		fmt.Println(errPost)
	}

	defer responsePost.Body.Close()

	/// POST response'unu okuyalım

	bodyBytesPost, _ := ioutil.ReadAll(responsePost.Body)

	bodyStringPost := string(bodyBytesPost)
	fmt.Println(bodyStringPost)

	var todoPostResponse Todo
	json.Unmarshal(bodyBytesPost, &todoPostResponse) // Unmarshal => JSON verisini Go struct'ına dönüştürür

	fmt.Printf("Todo: %+v \n", todoPostResponse)
}
