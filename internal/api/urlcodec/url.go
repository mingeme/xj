package urlcodec

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func StructToValues(data any) url.Values {
	values := url.Values{}
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		name := fieldType.Tag.Get("url")
		if name == "" {
			name = fieldType.Name
		}

		switch field.Kind() {
		case reflect.Int:
			values.Add(name, strconv.Itoa(int(field.Int())))
		case reflect.String:
			values.Add(name, field.String())
			// Add more cases as needed for other field types
		}
	}

	return values
}

func StructToStringReader(data any) *strings.Reader {
	return strings.NewReader(StructToValues(data).Encode())
}
