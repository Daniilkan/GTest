package unit

import (
	"bytes"
	"reflect"
	"testing"
)

type CompareType = compareResult
type compareResult int

const (
	CompareLess compareResult = iota - 1
	CompareEqual
	CompareGreater
	CompareError
)

func Compare(a interface{}, b interface{}, kind reflect.Kind) compareResult {
	switch kind {
	case reflect.Int:
		intObjA, ok := a.(int)
		if !ok {
			return CompareError
		}
		intObjB, ok := b.(int)
		if !ok {
			return CompareError
		}
		if intObjA > intObjB {
			return CompareGreater
		} else if intObjA < intObjB {
			return CompareLess
		} else {
			return CompareEqual
		}
	case reflect.Float32, reflect.Float64:
		floatObjA, ok := a.(float64)
		if !ok {
			return CompareError
		}
		floatObjB, ok := b.(float64)
		if !ok {
			return CompareError
		}
		if floatObjA > floatObjB {
			return CompareGreater
		} else if floatObjA < floatObjB {
			return CompareLess
		} else {
			return CompareEqual
		}
	case reflect.Slice:
		if !(reflect.ValueOf(a).CanConvert(reflect.TypeOf([]byte{}))) {
			break
		}
		bytesObjA, ok := a.([]byte)
		if !ok {
			return CompareError
		}
		bytesObjB, ok := b.([]byte)
		if !ok {
			return CompareError
		}
		return compareResult(bytes.Compare(bytesObjA, bytesObjB))
	}
	return CompareEqual
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
