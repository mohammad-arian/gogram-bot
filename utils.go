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
		err := w.WriteField(tag, s.(string))
		if err != nil {
			return err
		}
	case int:
		err := w.WriteField(tag, strconv.Itoa(s.(int)))
		if err != nil {
			return err
		}
	case float64:
		err := w.WriteField(tag, fmt.Sprintf("%v", s.(float64)))
		if err != nil {
			return err
		}
	case bool:
		err := w.WriteField(tag, strconv.FormatBool(s.(bool)))
		if err != nil {
			return err
		}
	// use *os.File for methods like SendVideo() and SendPhoto() that
	// the fieldName of CreateFormFile() can't be the name of the file.
	case *os.File:
		// some file fields are optional. below if statement makes sure program won't panic
		// even if a file field of data structure is empty.
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
			err := multipartSetter(f, w, "")
			if err != nil {
				return err
			}
		}
	default:
		Type := reflect.TypeOf(s).Kind()
		if Type == reflect.Slice || Type == reflect.Struct {
			a, err := json.Marshal(j)
			if err != nil {
				return err
			}
			err = w.WriteField(tag, string(a))
			if err != nil {
				return err
			}
		} else {
			return errors.New("incompatible type: " + Type.String())
		}
	}
	return nil
}

func structMultipartParser(s interface{}, w *multipart.Writer) error {
	for i := 0; i < reflect.ValueOf(s).Elem().NumField(); i++ {
		tag := reflect.TypeOf(s).Elem().Field(i).Tag.Get("json")
		value := reflect.ValueOf(s).Elem().Field(i).Interface()
		err := multipartSetter(value, w, tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func request(method string, bot Bot, data interface{},
	optionalParams interface{}, responseType interface{}) (response interface{}, error error) {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.Token, method),
		nil)
	var body = &bytes.Buffer{}
	var set bool
	w := multipart.NewWriter(body)
	if optionalParams != nil {
		if !reflect.ValueOf(optionalParams).IsNil() {
			err := structMultipartParser(optionalParams, w)
			if err != nil {
				return responseType, err
			}
			set = true
		}
	}
	if data != nil {
		if reflect.ValueOf(data).Kind() != reflect.Ptr {
			return responseType, errors.New("data parameter must be a pointer")
		}
		dataKind := reflect.Indirect(reflect.ValueOf(data)).Kind()
		if dataKind != reflect.Struct {
			return responseType, errors.New("data parameter must be a struct not " + dataKind.String())
		}
		err := structMultipartParser(data, w)
		if err != nil {
			return responseType, err
		}
		set = true
	}
	err := w.Close()
	if err != nil {
		return responseType, err
	}
	if set {
		req.Header.Add("Content-Type", w.FormDataContentType())
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	}
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	readRes, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(readRes, responseType)
	if err != nil {
		return responseType, err
	}
	return responseType, nil
}
