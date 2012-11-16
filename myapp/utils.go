package myapp

import (
	"bytes"
	"net/http"
	"encoding/json"

	"html/template"
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
