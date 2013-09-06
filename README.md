# Grout

## Simple routing for Golang

This routing is very similar to Django style routes. You just write some named
regular expressions and then the data shows up in your handler.

This Router looks for handler functions with this signature:

`type RouteHandler func(http.ResponseWriter, *http.Request, map[string]string)`

It's the same as http.Handler except it has an extra map[string]string associated with it.

## Installation

`$ go get github.com/chuckha/grout`

## Usage

(see example.go)
```go
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
```

## Things not working:

* Unnamed regex
* Combining routes (django's include)
* Anything more complex than this example
* Probably lots of bugs, no edge cases were tested.

Contributions are welcomed!
