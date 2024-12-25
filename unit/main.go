package unit

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

type CompareType = compareResult
type compareResult int

const (
	compareLess compareResult = iota - 1
	compareEqual
	compareGreater
	compareError
)

func Compare(a interface{}, b interface{}, kind reflect.Kind) (compareResult, error) {
	switch kind {
	case reflect.Int:
		intObjA, ok := a.(int)
		if !ok {
			return compareError, errors.New("Not needed type of A arg")
		}
		intObjB, ok := b.(int)
		if !ok {
			return compareError, errors.New("Not needed type of B arg")
		}
		if intObjA > intObjB {
			return compareGreater, nil
		} else if intObjA < intObjB {
			return compareLess, nil
		} else {
			return compareEqual, nil
		}
	case reflect.Float32, reflect.Float64:
		floatObjA, ok := a.(float64)
		if !ok {
			return compareError, errors.New("Not needed type of A arg")
		}
		floatObjB, ok := b.(float64)
		if !ok {
			return compareError, errors.New("Not needed type of B arg")
		}
		if floatObjA > floatObjB {
			return compareGreater, nil
		} else if floatObjA < floatObjB {
			return compareLess, nil
		} else {
			return compareEqual, nil
		}
	case reflect.Slice:
		if !(reflect.ValueOf(a).CanConvert(reflect.TypeOf([]byte{}))) {
			break
		}
		bytesObjA, ok := a.([]byte)
		if !ok {
			return compareError, errors.New("Not needed type of A arg")
		}
		bytesObjB, ok := b.([]byte)
		if !ok {
			return compareError, errors.New("Not needed type of B arg")
		}
		return compareResult(bytes.Compare(bytesObjA, bytesObjB)), nil
	}
	return compareEqual, nil
}

func Nil(t *testing.T, param interface{}) (interface{}, bool) {
	if param != nil {
		t.Logf("Expected nil, got not nil")
		return "Expected nil, got not nil", false
	}
	return nil, true
}

func IsEmpty(object interface{}) bool {
	if object == nil {
		return true
	}
	objValue := reflect.ValueOf(object)
	if objValue.Kind() == reflect.Chan || objValue.Kind() == reflect.Map || objValue.Kind() == reflect.Slice {
		return objValue.Len() == 0
	}
	return true
}
