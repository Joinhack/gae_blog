package myapp

import (
	"fmt"
	"appengine"
	
	"net/http"
	"html/template"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	t, e := template.ParseFiles("templates/index.html")
	if e != nil {
		fmt.Fprintf(w, "%s\n", e)
		return
	}
	models := make(map[string]interface{})
	blogs := GetBlogs(ctx, 0, 10)
	tags := GetTags(ctx)
	models[`blogs`] = blogs
	models["tags"] = tags
	err := t.Execute(w, models)
	if err != nil {
		panic(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("templates/login.html")
	if e != nil {
		fmt.Fprintf(w, "%s\n", e)
		return
	}
	t.Execute(w, nil)
}
