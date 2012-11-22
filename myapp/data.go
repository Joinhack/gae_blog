package myapp

import (
//	"fmt"
	"time"

	"encoding/json"
	app "appengine"
	ds "appengine/datastore"
)

type Comment struct {
	User string
	Value string
}


type Session struct {
	Id string
	Expires time.Time
	SessionData []byte
	datas map[string]interface{} `datastore:"-"`
}

type Tag struct {
	TagName string `datastore:",index"`
}

type User struct {
	Name string
	LoginId string `datastore:",index"`
	Password string
	Tags []*Tag
}

type Blog struct {
	Id int64
	Title string
	Tags []*Tag
	Time time.Time `datastore:",noindex"`
	Content string
}

func  NewSession()  (session *Session) {
	session = new(Session)
	session.datas = make(map[string]interface{})
	return session
}

func (session *Session) SetData(key string, data interface{}) {
	session.datas[key] = data
}

func (session *Session) GetData(key string)interface{} {
	return session.datas[key]
}

func (session *Session) RemovetData(key string) {
	delete(session.datas, key)
}

func (session *Session) Save(ctx app.Context) (err error) {
	key := ds.NewKey(ctx, "Session", session.Id, 0, nil)
	session.SessionData, err = json.Marshal(session.datas)
	if err != nil {
		return nil
	}
	key, err = ds.Put(ctx, key, session)
	return err
}

func (session *Session) GetById(ctx app.Context) error {
	key := ds.NewKey(ctx, "Session", session.Id, 0, nil)
	var err error = nil
	err = ds.Get(ctx, key, session)
	if err != nil {
		return err
	}
	err = json.Unmarshal(session.SessionData, &session.datas)
	return err
}

func (u *User) Add(ctx app.Context) (err error) {
	if err != nil {
		return
	}
	key := ds.NewKey(ctx, "User", u.LoginId, 0, nil)
	_, err = ds.Put(ctx, key, u)

	return
}

func GetUserByLoginId(ctx app.Context, loginId string) (user *User, err error) {
	key := ds.NewKey(ctx, "User", loginId, 0, nil)
	user = new(User)
	err = ds.Get(ctx, key, user)
	if err == ds.ErrNoSuchEntity {
		return nil, nil
	}
	return user, nil
}

func GetTags(ctx app.Context) []*Tag {
	var tags = make([]*Tag, 0, 20)
	q := ds.NewQuery("Tag")
	for iter := q.Run(ctx); ; {
		var tag Tag
		_, err := iter.Next(&tag)
		if err == ds.Done {
			break
		}
		tags = append(tags, &tag)
	}
	return tags
}

func GetBlogs(ctx app.Context, offset, limit int) []*Blog {
	var blogs = make([]*Blog, 0, 20)
	q := ds.NewQuery("Blog").Offset(offset).Limit(limit).Order("-time")
	for iter := q.Run(ctx); ; {
		var blog Blog
		_, err := iter.Next(&blog)		
		if err == ds.Done {
			break
		}
		blogs = append(blogs, &blog)
		
	}
	return blogs
}

