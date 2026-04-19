package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"completed"`
}

func fetchPosts(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("posts error:", err)
		return
	}
	defer resp.Body.Close()

	var post Post
	json.NewDecoder(resp.Body).Decode(&post)

	fmt.Println("Post:", post.Title)
}

func fetchUser(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("user error:", err)
		return
	}
	defer resp.Body.Close()

	var user User
	json.NewDecoder(resp.Body).Decode(&user)

	fmt.Println("User:", user.Name)
}

func fetchTodo(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("todo error:", err)
		return
	}
	defer resp.Body.Close()

	var todo Todo
	json.NewDecoder(resp.Body).Decode(&todo)

	fmt.Println("Todo:", todo.Title, "Done:", todo.Done)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go fetchPosts(&wg)
	go fetchUser(&wg)
	go fetchTodo(&wg)

	wg.Wait()
	fmt.Println("All API calls finished")
}
