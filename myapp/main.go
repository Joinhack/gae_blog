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
	http.HandleFunc("/login_do", login_do)
	http.HandleFunc("/user/add", userAdd)
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

func userAdd(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	user := new(User)
	user.Name = r.FormValue("userName")
	user.Password = r.FormValue("passwd")
	user.LoginId = r.FormValue("email")
	err := user.Add(ctx)
	if err != nil {
		OutputJson(w, &map[string]interface{}{"code":-1, "msg": err.Error()})
		return
	}
	OutputJson(w, &map[string]interface{}{"code":0, "msg": "sucess", "loginId": user.LoginId})
}

func login_do(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	loginId := r.FormValue("loginId")
	password := r.FormValue("password")
	user, err := GetUserByLoginId(ctx, loginId)
	if err != nil {
		OutputJson(w, &map[string]interface{}{"code":-1, "msg": err.Error()})
		return
	}
	if user == nil || password != user.Password {
		OutputJson(w, &map[string]interface{}{"code":-1, "msg": "user or password error"})
		return
	}
	OutputJson(w, &map[string]interface{}{"code":0, "msg": "sucess", "loginId": user.LoginId})
}

func login(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("templates/login.html")
	if e != nil {
		fmt.Fprintf(w, "%s\n", e)
		return
	}
	t.Execute(w, nil)
}
