package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Posts struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
type PostResult struct {
	Posts []Posts `json:"data"`
	Error error   `json:"error"`
}

type Users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserResult struct {
	Users []Users `json:"data"`
	Error error   `json:"error"`
}

type CombinedResult struct {
	UserResult UserResult `json:"users"`
	PostResult PostResult `json:"posts"`
}

// https://jsonplaceholder.typicode.com/posts?_start=0&_limit=3
func fetchPosts(wg *sync.WaitGroup, jwtToken string, ch chan<- PostResult) {
	defer wg.Done()
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?_start=0&_limit=3")
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts?_start=0&_limit=3", nil)
	if err != nil {
		fmt.Println(err)
		postRes := PostResult{
			Posts: nil,
			Error: err,
		}
		ch <- postRes
		return
	}
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		postRes := PostResult{
			Posts: nil,
			Error: err,
		}
		ch <- postRes
		return
	}
	var posts []Posts
	json.NewDecoder(resp.Body).Decode(&posts)
	postRes := PostResult{
		Posts: posts,
		Error: nil,
	}
	ch <- postRes

}

// https://jsonplaceholder.typicode.com/users?_start=0&_limit=3
func fetchUsers(wg *sync.WaitGroup, jwtToken string, ch chan<- UserResult) {
	defer wg.Done()
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/users?_start=0&_limit=3")
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users?_start=0&_limit=3", nil)
	if err != nil {
		fmt.Println(err)
		userRes := UserResult{
			Users: nil,
			Error: err,
		}
		ch <- userRes
		return
	}
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		userRes := UserResult{
			Users: nil,
			Error: err,
		}
		ch <- userRes
		return
	}
	var users []Users
	json.NewDecoder(resp.Body).Decode(&users)
	userRes := UserResult{
		Users: users,
		Error: nil,
	}
	ch <- userRes

}

func FetchData(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup
	wg.Add(2)

	chPosts := make(chan PostResult)
	chUsers := make(chan UserResult)
	jwtToken := "Token JWT"

	go fetchPosts(&wg, jwtToken, chPosts)
	go fetchUsers(&wg, jwtToken, chUsers)

	results := CombinedResult{
		UserResult: <-chUsers,
		PostResult: <-chPosts,
	}
	wg.Wait()

	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
