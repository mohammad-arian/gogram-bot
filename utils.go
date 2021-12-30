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

func multipartSetter(s interface{}, w *multipart.Writer, tag string) error {
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
	// use *os.File for methods like SendVideo() and SendPhoto() that
	// the fieldName of CreateFormFile() can't be the name of the file.
	case *os.File:
		name := tag
		if name == "" {
			name = j.Name()
		}
		file, _ := w.CreateFormFile(name, j.Name())
		_, _ = io.Copy(file, j)
		_, _ = j.Seek(0, io.SeekStart)
	case []*os.File:
		for _, f := range j {
			if err := multipartSetter(f, w, ""); err != nil {
				return err
			}
		}
	default:
		Type := reflect.TypeOf(s).Kind()
		if Type == reflect.Slice || Type == reflect.Struct {
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

func structMultipartParser(s interface{}, w *multipart.Writer) error {
	for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
		tag := reflect.TypeOf(s).Field(i).Tag.Get("json")
		value := reflect.ValueOf(s).Field(i).Interface()
		if err := multipartSetter(value, w, tag); err != nil {
			return err
		}
	}
	return nil
}

func request(method string, bot Bot, data interface{}, responseType interface{}) (response interface{}, error error) {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.Token, method),
		nil)
	var body = &bytes.Buffer{}
	var set bool
	w := multipart.NewWriter(body)
	if data != nil {
		if err := structMultipartParser(data, w); err != nil {
			w.Close()
			return responseType, err
		}
		set = true
	}
	w.Close()
	if set {
		req.Header.Add("Content-Type", w.FormDataContentType())
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	}
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	readRes, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(readRes, responseType)
	if err != nil {
		return responseType, err
	}
	return responseType, nil
}
