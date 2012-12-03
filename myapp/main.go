package myapp

import (
	"fmt"
	"time"
	"strings"
	"appengine"

	"net/http"
	"text/template"
)

func init() {
	var route = NewRoute()
	route.HandleFunc("/", index)
	/**
	http.HandleFunc("/login", login)
	http.HandleFunc("/login/do", login_do)
	http.HandleFunc("/login/json", login_json)
	http.HandleFunc("/user/add", userAdd)
	http.HandleFunc("/list/*", listByTag)
	http.HandleFunc("/new_topic", newTopic)
	http.HandleFunc("/new_topic/save", newTopic_save)*/
}

func isLogin(ctx *appengine.Context, w http.ResponseWriter, r *http.Request) ( *User, bool) {
	cookie, err := r.Cookie("UH")
	if err == http.ErrNoCookie {
		return nil, false
	}
	now := time.Now()
	var session *Session = NewSession()
	session.Id = cookie.Value
	session.GetById(*ctx)
	if session.Expires.After(now) {
		user, _ := GetUserByLoginId(*ctx, session.GetData("loginId").(string))
		return user,true
	}
	return nil, false
}

func newSession(user *User) *Session {
	now := time.Now()
	expires := now.Add(time.Duration(300) * time.Second)
	session := NewSession()
	session.Id = UserInfoHash(user, &expires)
	session.Expires = expires
	return session
}

func newCookie(r *http.Request, session *Session) *http.Cookie{
	cookie := new(http.Cookie)
	cookie.Name = "UH"
	cookie.Value = session.Id
	cookie.Expires = session.Expires
	cookie.Path = "/"
	cookie.Domain = r.URL.Host
	return cookie
}

func login_json(w http.ResponseWriter, r *http.Request) {
	templateName := "login_form"
	t, _ := template.ParseFiles("templates/login_form.html")
	content, _ := Template2String(t, &templateName, nil)
	OutputJson(w, &map[string]interface{}{"code":0, "msg": "sucess", "content": content})
}

func newTopic_save(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	blog := NewBlog()
	blog.Title = r.FormValue("title")
	blog.Content = r.FormValue("content")
	tagsStr := r.FormValue("tags")
	for _, tagStr := range strings.Split(tagsStr, " ") {
		tagStr = strings.Trim(tagStr, " ")
		if tagStr == "" {
			continue;
		}
		blog.Tags = append(blog.Tags, tagStr)
	}
	blog.Time = time.Now()
	err := blog.Save(ctx)
	if err != nil {
		OutputJson(w, &map[string]interface{}{"code":-1, "msg": err.Error()})
		return
	}
	OutputJson(w, &map[string]interface{}{"code":0, "msg": "sucess"})
}

func newTopic(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	templateName := "newTopic"
	if _, b := isLogin(&ctx, w, r); !b {
		http.Redirect(w,r, "/login/json", 302)
	}
	t, _ := template.ParseFiles("templates/newTopic.html")
	content, _ := Template2String(t, &templateName, nil)
	OutputJson(w, &map[string]interface{}{"code":0, "msg": "sucess", "content": content})
}

func listByTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	ctx := appengine.NewContext(r)
	t, e := template.ParseFiles("templates/index.html")
	if e != nil {
		fmt.Fprintf(w, "%s\n", e)
		return
	}
	models := make(map[string]interface{})
	blogs, err := GetBlogs(ctx, 0, 10)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	tags, err := GetTags(ctx)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	models[`blogs`] = blogs
	models["tags"] = tags
	err = t.Execute(w, models)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request, pathParam map[string]string) {
	ctx := appengine.NewContext(r)
	t, e := template.ParseFiles("templates/index.html")
	if e != nil {
		fmt.Fprintf(w, "%s\n", e)
		return
	}
	models := make(map[string]interface{})
	blogs, err := GetBlogs(ctx, 0, 10)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	tags, err := GetTags(ctx)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	models[`blogs`] = blogs
	models["tags"] = tags
	err = t.Execute(w, models)
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
	session := newSession(user)
	session.SetData("loginId", loginId)
	session.Save(ctx)
	http.SetCookie(w, newCookie(r, session))
	OutputJson(w, &map[string]interface{}{"code":0, "msg": "sucess", "loginId": user.LoginId})
}

func login(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("templates/login.html", "templates/login_form.html")
	if e != nil {
		fmt.Fprintf(w, "%s\n", e)
		return
	}
	t.Execute(w, nil)
}
