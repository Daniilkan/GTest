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
	if result := gtu.Compare(1, 2, reflect.Int); result != gtu.CompareLess {
		t.Errorf("Expected CompareLess, got %v", result)
	}
	if result := gtu.Compare(2.0, 1.0, reflect.Float64); result != gtu.CompareGreater {
		t.Errorf("Expected CompareGreater, got %v", result)
	}
	if result := gtu.Compare([]byte{1, 2}, []byte{1, 2}, reflect.Slice); result != gtu.CompareEqual {
		t.Errorf("Expected CompareEqual, got %v", result)
	}

	// Test Nil function
	if !gtu.Nil(nil) {
		t.Errorf("Expected true for nil input")
	}
	if gtu.Nil(1) {
		t.Errorf("Expected false for non-nil input")
	}

	// Test IsEmpty function
	if !gtu.IsEmpty(nil) {
		t.Errorf("Expected true for nil input")
	}
	if !gtu.IsEmpty([]int{}) {
		t.Errorf("Expected true for empty slice")
	}
	if gtu.IsEmpty([]int{1}) {
		t.Errorf("Expected false for non-empty slice")
	}

	// Test CheckFunctionResult function
	fn := func(a, b int) int { return a + b }
	if !gtu.CheckFunctionResult(fn, []interface{}{1, 2}, 3) {
		t.Errorf("Expected true for correct function result")
	}
	if gtu.CheckFunctionResult(fn, []interface{}{1, 2}, 4) {
		t.Errorf("Expected false for incorrect function result")
	}
}

func TestHTTPFunctions(t *testing.T) {
	// Mock response for testing
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(`{"key":"value"}`)),
	}

	// Test ResponseCheck function
	if result := gth.ResponseCheck(mockResponse); result != "Success" {
		t.Errorf("Expected Success, got %v", result)
	}

	// Test ResponseNotEmpty function
	if !gth.ResponseNotEmpty(mockResponse) {
		t.Errorf("Expected true for non-empty response body")
	}

	// Test GetResponseType function
	if result := gth.GetResponseType(mockResponse); result != gth.TypeJson {
		t.Errorf("Expected TypeJson, got %v", result)
	}

	// Test ResponseContains function
	if !gth.ResponseContains(mockResponse, "key") {
		t.Errorf("Expected true for response containing 'key'")
	}

	// Test WebPageWorking function
	if !gth.WebPageWorking("https://wikipedia.org") {
		t.Errorf("Expected true for working web page")
	}

	// Test WebPageContains function
	if !gth.WebPageContains("https://wikipedia.org", "English") {
		t.Errorf("Expected true for web page containing 'English'")
	}
}
