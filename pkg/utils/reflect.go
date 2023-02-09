package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	ErrNoFieldFound = errors.New("no field found")
)

func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func ReflectSetField(to interface{}, fieldName string, value interface{}) error {
	target := reflect.ValueOf(to).Elem()
	kind := reflect.TypeOf(value).Kind()
	s := fmt.Sprintf("%v", value)

	if target.FieldByName(fieldName) == (reflect.Value{}) {
		return ErrNoFieldFound
	}

	switch kind {
	case reflect.Int, reflect.Int64, reflect.Int32:
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		target.FieldByName(fieldName).SetInt(v)
	case reflect.Float64:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		target.FieldByName(fieldName).SetFloat(v)
	case reflect.String:
		target.FieldByName(fieldName).SetString(s)

	default:
		panic(fmt.Errorf("unrecognized filed type: %s when refect set target field: %s value", kind, fieldName))
	}
	return nil
}
