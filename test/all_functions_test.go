package test

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	gth "github.com/Daniilkan/GTest/http"
	gtu "github.com/Daniilkan/GTest/unit"
)

func TestUnitFunctions(t *testing.T) {
	// Test Compare function

	values := [][]interface{}{
		{1, 2, reflect.Int, gtu.CompareLess},
		{2, 1, reflect.Int, gtu.CompareGreater},
		{2.0, 1.0, reflect.Float64, gtu.CompareGreater},
		{[]byte{1, 2}, []byte{1, 2}, reflect.Slice, gtu.CompareEqual},
	}

	for _, v := range values {
		if result := gtu.Compare(v[0], v[1], v[2].(reflect.Kind)); result != v[3] {
			t.Errorf("Expected %v, got %v", v[3], result)
		}
	}

	values = [][]interface{}{{nil, true}, {1, false}}
	for _, v := range values {
		if result := gtu.Nil(v[0]); result != v[1] {
			t.Errorf("Expected %v, got %v", v[1], result)
		}
	}

	values = [][]interface{}{{nil, true}, {[]int{1, 5, 7}, false}, {[]int{}, true}}

	for _, v := range values {
		if result := gtu.IsEmpty(v[0]); result != v[1] {
			t.Errorf("Expected %v, got %v", v[1], result)
		}
	}
	// Test CheckFunctionResult function
	values = [][]interface{}{{func(a, b int) int { return a + b }, []interface{}{1, 2}, 3, true}, {func(a, b int) int { return a + b }, []interface{}{1, 2}, 4, false}, {func(a, b int) int { return a * b }, []interface{}{1, 2}, 2, true}}

	for _, v := range values {
		if result := gtu.CheckFunctionResult(v[0], v[1].([]interface{}), v[2]); result != v[3] {
			t.Errorf("Expected %v, got %v", v[3], result)
		}
	}
}

func TestHTTPFunctions(t *testing.T) {
	// Test ResponseCheck function
	values := [][]interface{}{
		{&http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"key":"value"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, "Success"}, {&http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(strings.NewReader(`{"key":"value"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, "Fail"},
	}
	for _, v := range values {
		if result := gth.ResponseCheck(v[0].(*http.Response)); result != v[1] {
			t.Errorf("Expected %v, got %v", v[1], result)
		}
	}

	// Test ResponseNotEmpty function
	values = [][]interface{}{
		{&http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"key":"value"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, true}, {&http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(``)),
		}, false},
	}
	for _, v := range values {
		if result := gth.ResponseNotEmpty(v[0].(*http.Response)); result != v[1] {
			t.Errorf("Expected %v, got %v", v[1], result)
		}
	}

	// Test GetResponseType function
	values = [][]interface{}{
		{&http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"key":"value"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, gth.TypeJson}, {&http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(strings.NewReader(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    
</body>
</html>`)),
			Header: http.Header{"Content-Type": []string{"text/html"}},
		}, gth.TypeHtml},
	}
	for _, v := range values {
		if result := gth.GetResponseType(v[0].(*http.Response)); result != v[1] {
			t.Errorf("Expected %v, got %v", v[1], result)
		}
	}

	// Test ResponseContains function
	values = [][]interface{}{
		{&http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"key":"value"}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, "key", true}, {&http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(strings.NewReader(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    
</body>
</html>`)),
			Header: http.Header{"Content-Type": []string{"text/html"}},
		}, "body", true},
	}
	for _, v := range values {
		if result := gth.ResponseContains(v[0].(*http.Response), v[1].(string)); result != v[2] {
			t.Errorf("Expected %v, got %v", v[2], result)
		}
	}

	// Test WebPageWorking function
	values = [][]interface{}{
		{"https://wikipedia.org", true},
		{"https://wikipedia.org/408", false},
	}
	for _, v := range values {
		if result := gth.WebPageWorking(v[0].(string)); result != v[1] {
			t.Errorf("Expected %v, got %v", v[1], result)
		}
	}

	// Test WebPageContains function
	values = [][]interface{}{
		{"https://wikipedia.org", "English", true},
		{"https://wikipedia.org", "Russian", false},
	}
	for _, v := range values {
		if result := gth.WebPageContains(v[0].(string), v[1].(string)); result != v[2] {
			t.Errorf("Expected %v, got %v", v[2], result)
		}
	}
}
