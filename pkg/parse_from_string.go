package pkg

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func ParseFromString(f any, str string) error {
	if str == "" {
		return nil
	}

	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("target must be non-nil pointer")
	}

	parser, ok := map[reflect.Kind]func(string) (any, error){
		reflect.Int64:  parseInt64,
		reflect.Bool:   parseBool,
		reflect.Struct: parseTime,
	}[v.Elem().Kind()]

	if !ok {
		return fmt.Errorf("unsupported type: %T", f)
	}

	parsed, err := parser(str)
	if err != nil {
		return err
	}

	v.Elem().Set(reflect.ValueOf(parsed))
	return nil
}

func parseInt64(s string) (any, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err
}

func parseBool(s string) (any, error) {
	return strconv.ParseBool(s)
}

func parseTime(s string) (any, error) {
	return time.Parse(time.RFC3339, s)
}
