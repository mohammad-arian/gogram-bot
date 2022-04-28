package gogram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

func multipartSetter(s any, w *multipart.Writer, tag string) error {
	switch j := s.(type) {
	case string:
		if err := w.WriteField(tag, s.(string)); err != nil {
			return err
		}
	case int:
		if err := w.WriteField(tag, strconv.Itoa(s.(int))); err != nil {
			return err
		}
	case float64:
		if err := w.WriteField(tag, fmt.Sprintf("%v", s.(float64))); err != nil {
			return err
		}
	case bool:
		if err := w.WriteField(tag, strconv.FormatBool(s.(bool))); err != nil {
			return err
		}
	case nil:
		return nil
	// use *os.File is for methods like SendVideo() and SendPhoto() that
	// the fieldName of CreateFormFile() can't be the name of the file, instead it must be json tag.
	case *os.File:
		if j != nil {
			name := tag
			if name == "" {
				name = j.Name()
			}
			file, _ := w.CreateFormFile(name, j.Name())
			_, _ = io.Copy(file, j)
			_, _ = j.Seek(0, io.SeekStart)
		}
	case []*os.File:
		for _, f := range j {
			if err := multipartSetter(f, w, ""); err != nil {
				return err
			}
		}
	default:
		Type := reflect.TypeOf(s).Kind()
		if Type == reflect.Slice || Type == reflect.Struct || Type == reflect.Ptr {
			// for Keyboard and InlineKeyboard, ReplyKeyboard
			if Type == reflect.Struct && tag == "" {
				return structMultipartParser(j, w)
			}
			a, err := json.Marshal(j)
			if err != nil {
				return err
			}
			if err = w.WriteField(tag, string(a)); err != nil {
				return err
			}
		} else {
			return errors.New("incompatible type: " + Type.String())
		}
	}
	return nil
}

func structMultipartParser(s any, w *multipart.Writer) error {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		return errors.New("value is a pointer")
	}
	for i := 0; i < v.NumField(); i++ {
		tag := reflect.TypeOf(s).Field(i).Tag.Get("json")
		value := v.Field(i).Interface()
		if err := multipartSetter(value, w, tag); err != nil {
			return err
		}
	}
	return nil
}

func Request(method string, bot Bot, data Method, response Response) (Response, error) {
	if err := data.Check(); err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.Token,
		method), nil)
	var body = &bytes.Buffer{}
	w := multipart.NewWriter(body)
	if data != nil {
		if err := structMultipartParser(data, w); err != nil {
			return response, err
		}
	}
	w.Close()
	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	if res, err := http.DefaultClient.Do(req); err != nil {
		return response, err
	} else {
		defer res.Body.Close()
		return response.set(res)
	}
}

// globalEmptyFieldChecker is for general cases that we want to Check if a field is empty or not.
// it accepts a map in which keys are field names and values are checked against switch cases.
// if a field is empty, an error in "<field name> is empty" format will be returned.
// globalEmptyFieldChecker treats ParseMode fields differently. if a field name is ParseMode, and its type is
// string, then its value must be either "MarkdownV2", "HTML" or "Markdown"
func globalEmptyFieldChecker(a map[string]any) error {
	for i, j := range a {
		switch v := j.(type) {
		case string:
			if v == "" && i != "ParseMode" {
				return errors.New(i + " is empty")
			} else if i == "ParseMode" && v != "" {
				if v != "MarkdownV2" && v != "HTML" && v != "Markdown" {
					return errors.New(i + " must be MarkdownV2, HTML or Markdown")
				}
			}
		case int:
			if v == 0 {
				return errors.New(i + " is empty")
			}
		case float64:
			if v == 0 {
				return errors.New(i + " is empty")
			}
		case bool:
			if v == false {
				return errors.New(i + " is false")
			}
		case nil:
			return errors.New(i + " is empty")
		case *os.File:
			if j == nil {
				return errors.New(i + " is empty")
			}
		case []*os.File:
			if j == nil {
				return errors.New(i + " is empty")
			}
		default:
			Type := reflect.ValueOf(v)
			if Type.Kind() == reflect.Slice {
				if Type.Len() == 0 {
					return errors.New(i + " is empty")
				}
			}
		}
	}
	return nil
}
