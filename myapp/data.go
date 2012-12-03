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
	Tags []string
	Time time.Time
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

func GetTags(ctx app.Context) ([]*Tag, error) {
	var tags = make([]*Tag, 0, 20)
	q := ds.NewQuery("Tag")
	_, err := q.GetAll(ctx, &tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func NewBlog() *Blog {
	return new(Blog);
}

func (blog *Blog) GetById(ctx app.Context) error {
	key := ds.NewKey(ctx, "Blog", "", blog.Id, nil)
	err := ds.Get(ctx, key, blog)
	if err != nil {
		return err
	}
	return nil
}

func (blog *Blog) Save(ctx app.Context) error {
	id, _, err := ds.AllocateIDs(ctx, "Blog", nil, 1)
	if err != nil {
		return nil
	}
	var tagKeys = make([]*ds.Key, 0, len(blog.Tags))
	var tags = make([]*Tag, 0, len(blog.Tags))
	for _, tag := range blog.Tags {
		tagKeys = append(tagKeys, ds.NewKey(ctx, "Tag", tag, 0, nil))
		tags = append(tags, &Tag{tag})
	}
	_, err = ds.PutMulti(ctx, tagKeys, tags)
	if err != nil {
		return err;
	}
	key := ds.NewKey(ctx, "Blog", "", id, nil)
	blog.Id = id
	_, err = ds.Put(ctx, key, blog)
	if err != nil {
		return err
	}
	return nil
}


func GetBlogsByTag(ctx app.Context, tag string, offset, limit int) ([]*Blog, error) {
	var blogs []*Blog
	q := ds.NewQuery("Blog").Filter("Tags ~=", tag)
	q.Offset(0).Limit(10).Order("-Time")
	_, err := q.GetAll(ctx, &blogs)
	if err != nil {
		return nil, err
	}	
	return blogs, nil
}



func GetBlogs(ctx app.Context, offset, limit int) ([]*Blog, error) {
	var blogs []*Blog
	q := ds.NewQuery("Blog").Offset(0).Limit(10).Order("-Time")
	_, err := q.GetAll(ctx, &blogs)
	if err != nil {
		return nil, err
	}	
	return blogs, nil
}

