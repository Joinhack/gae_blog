package myapp

import (
	"time"

	"appengine"
	"appengine/datastore"
)

type Comment struct {
	User string
	Value string
}

type Tag struct {
	TagName string `datastore:",index"`
}

type User struct {
	Name string
	LoginId string
	Password string
	Tags []Tag
}

type Blog struct {
	Id uint32
	Title string
	Tags []Tag
	Time time.Time `datastore:",noindex"`
	Content string
}

func (u *User) Add(ctx appengine.Context) (err error) {
	ctx.Put(*u)
}

func GetTags(ctx appengine.Context) []Tag {
	var tags []Tag = make([]Tag, 0, 20)
	q := datastore.NewQuery("Tag")
	for iter := q.Run(ctx); ; {
		var tag Tag
		_, err := iter.Next(&tag)		
		if err == datastore.Done {
			break
		}
		tags = append(tags, tag)
	}
	return tags
}

func GetBlogs(ctx appengine.Context, offset, limit int) []Blog {
	var blogs []Blog = make([]Blog, 0, 20)
	q := datastore.NewQuery("Blog").Offset(offset).Limit(limit).Order("-time")
	for iter := q.Run(ctx); ; {
		var blog Blog
		_, err := iter.Next(&blog)		
		if err == datastore.Done {
			break
		}
		blogs = append(blogs, blog)
		
	}
	return blogs
}

