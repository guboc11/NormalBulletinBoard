package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Post struct represents a single post in the message board
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user queried")
	// // Query the database for posts
	// rows, err := db.Query("SELECT id, title, body FROM posts")
	// if err != nil {
	// 	http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
	// 	return
	// }
	// defer rows.Close()

	// Create a slice to hold retrieved posts
	posts := []Post{}
	posts = append(posts, Post{1, "title1", "body1"})
	posts = append(posts, Post{2, "title2", "body2"})

	// // Iterate through the rows and append posts to the slice
	// for rows.Next() {
	// 	var post Post
	// 	if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
	// 		http.Error(w, "Failed to scan posts", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	posts = append(posts, post)
	// }

	// // Check for errors during row iteration
	// if err := rows.Err(); err != nil {
	// 	http.Error(w, "Failed to iterate over posts", http.StatusInternalServerError)
	// 	return
	// }

	// Marshal posts slice to JSON
	postJSON, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Failed to marshal posts to JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(postJSON)
}

func main() {
	fmt.Println("hello backend")

	// // Database connection string
	// db, err := sql.Open("postgres", "host=localhost port=5432 user=yourusername password=yourpassword dbname=yourdbname sslmode=disable")
	// if err != nil {
	// 	log.Fatal("Could not connect to the database:", err)
	// }
	// defer db.Close()
	// // API endpoint to retrieve posts
	http.HandleFunc("/posts", getPostsHandler)

	// Start the server
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
