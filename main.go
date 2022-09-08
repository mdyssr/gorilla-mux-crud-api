package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Author struct {
	FullName string `json:"fullName"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author Author `json:"author"`
}

var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error: invalid id format"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("id not found"))
		return
	}

	post := posts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	// get item value from the request body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	w.Header().Set("Content-Type", "application/json")
	posts = append(posts, newPost)

	json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error: invalid id"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("id not found"))
		return
	}

	// get the value from the json body
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)
	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

func patchPost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error: invalid id"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("id not found"))
		return
	}

	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)
	// posts[id] = post

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error: invalid id"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("id not found"))
		return
	}
	posts = append(posts[:id], posts[id+1:]...)
}
