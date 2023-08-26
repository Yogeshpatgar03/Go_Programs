package main
import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
)
type Post struct {
UserID int `json:"userId"`
ID int `json:"id"`
Title string `json:"title"`
Body string `json:"body"`
}
func main() {
var postID int
fmt.Print("Enter the Post ID: ")
fmt.Scan(&postID)
url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID)
response, err := http.Get(url)
if err != nil {
fmt.Printf("HTTP GET request failed: %s\n", err)
return
}
defer response.Body.Close()
body, err := ioutil.ReadAll(response.Body)
if err != nil {
fmt.Printf("Failed to read response body: %s\n", err)
return
}
var post Post
err = json.Unmarshal(body, &post)
if err != nil {
fmt.Printf("Failed to parse JSON: %s\n", err)
return
}
fmt.Printf("UserID: %d\n", post.UserID)
fmt.Printf("ID: %d\n", post.ID)
fmt.Printf("Title: %s\n", post.Title)
fmt.Printf("Body: %s\n", post.Body)
}