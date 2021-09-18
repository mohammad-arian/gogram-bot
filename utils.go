package gogram

import (
	"io"
	"mime/multipart"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func urlValueSetter(s interface{}, q *url.Values) {
	for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
		tag := reflect.TypeOf(s).Field(i).Tag.Get("json")
		value := reflect.ValueOf(s).Field(i).Interface()
		switch j := value.(type) {
		case string:
			q.Set(tag, value.(string))
		case int:
			q.Set(tag, strconv.Itoa(value.(int)))
		case bool:
			q.Set(tag, strconv.FormatBool(value.(bool)))
		case KeyboardMarkup:
			q.Set(tag, j.toString())
		}
	}
}

func formFieldSetter(s interface{}, w *multipart.Writer) {
	for i := 0; i < reflect.ValueOf(w).NumField(); i++ {
		tag := reflect.TypeOf(w).Field(i).Tag.Get("json")
		value := reflect.ValueOf(s).Field(i).Interface()
		switch j := value.(type) {
		case string:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(value.(string)))
		case int:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(strconv.Itoa(value.(int))))
		case bool:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(strconv.FormatBool(value.(bool))))
		case KeyboardMarkup:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(j.toString()))
		}
	}
}
