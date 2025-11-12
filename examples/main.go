package examples

import (
	"fmt"
	"net/http"
	"reflect"

	gth "github.com/Daniilkan/GTest/http"
	gtu "github.com/Daniilkan/GTest/unit"
)

func unitExample() {

	obj1 := 9.45
	obj2 := 69264.586

	// Compare() returns GTest/unit type of comparison two objects
	switch gtu.Compare(obj1, obj2, reflect.Float32) {
	case gtu.CompareGreater:
		fmt.Println("Greater")
		break
	case gtu.CompareLess:
		fmt.Println("Less")
		break
	case gtu.CompareEqual:
		fmt.Println("Equal")
		break
	default:
		fmt.Println("Error")
		break
	}

	fmt.Println(gtu.IsEmpty(obj1)) // Returns true/false if object is empty

	fmt.Println(gtu.Nil(obj1)) // Returns true/false if object is nil
}

func httpExample() {
	url := "https://go.dev"

	response, _ := http.Get(url)

	fmt.Println(gth.ResponseCheck(response)) // prints "Success"

	fmt.Println(gth.ResponseContains(response, []byte("Go"))) // Returns true/false if response contains string

	fmt.Println(gth.ResponseNotEmpty(response)) // Returns true if response not empty, false if it is

	// This block returns type of response
	switch gth.GetResponseType(response) {
	case gth.TypeHtml:
		fmt.Println("HTML")
		break
	case gth.TypeJson:
		fmt.Println("JSON")
		break
	case gth.TypeError:
		fmt.Println("Error")
		break
	}
	// Returns HTML as go.dev is a webpage

	fmt.Println(gth.WebPageWorking(url)) // Returns true/false if webpage is working

	fmt.Println(gth.WebPageContains(url, []byte("Go"))) // Return true/false if webpage contains string
}
