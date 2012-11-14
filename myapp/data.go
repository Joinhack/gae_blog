package myapp

import (
	"time"

	app "appengine"
	ds "appengine/datastore"
)

const (
	MaxInt = (1<<31 - 1)
)

type Comment struct {
	User string
	Value string
}

type Tag struct {
	TagName string `datastore:",index"`
}

type User struct {
	Id int64
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

func (u *User) Add(ctx app.Context) (key *ds.Key, err error) {
	id, _, err := ds.AllocateIDs(ctx, "Users", nil, 1)
	u.Id = id
	if err != nil {
		return nil, err
	}
	key = ds.NewKey(ctx, "Users", "id", id, nil)
	return ds.Put(ctx, key, u)
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

