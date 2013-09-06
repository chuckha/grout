package main

import (
	"fmt"
	"github.com/ChuckHa/grout"
	"net/http"
)

type Blog struct {
}

func (b *Blog) PostById(id string) string {
	return "a blog post"
}

func BlogByName(name string) *Blog {
	return &Blog{}
}

func blogsHandler(w http.ResponseWriter, r *http.Request, data map[string]string) {
	// "name" and "othername" comes from the named Regex used in the URL
	theBlog := BlogByName(data["name"])
	post := theBlog.PostById(data["othername"])
	fmt.Fprintf(w, post)
}

func main() {
	mux := grout.NewRouteMux()
	mux.Route(`/blogs/(?P<name>[a-z][a-z_-]+[a-z])/(?P<othername>[0-9]+)`, blogsHandler)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
