package myapp

import (
	"bytes"
	"fmt"
	"time"
	"io"

	"crypto/md5"
	"net/http"
	"encoding/json"
	"text/template"
)

func OutputJson(w http.ResponseWriter, data interface{}) (err error) {
	w.Header().Set("content-type", "text/plain; charset=UTF-8")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(data)
	return
}

func Template2String(t *template.Template,name *string, data interface{}) (str *string, err error) {
	str = new(string)
	buf := bytes.NewBuffer(nil)
	if name == nil {
		err = t.Execute(buf, data)
	} else {
		err = t.ExecuteTemplate(buf, *name, data)
	}
	if err != nil {
		return nil, err
	}
	*str = buf.String()
	return str, nil
}

func UserInfoHash(user *User, expired *time.Time) string {
	md5 := md5.New()
	s := fmt.Sprintf("%s%s%d",user.LoginId, user.Password, expired.Unix())
	io.WriteString(md5, s)
	return fmt.Sprintf("%x", md5.Sum(nil))
}
