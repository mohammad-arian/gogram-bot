package gogram

import (
	"net/url"
	"reflect"
	"strconv"
)

func parameterSetter(s interface{}, q *url.Values) {
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
