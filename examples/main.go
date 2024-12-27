package main

import (
	"fmt"
	"net/http"

	gt "github.com/Daniilkan/GTest/http"
)

func main() {
	url := "https://wikipedia.org"

	response, _ := http.Get(url)

	fmt.Println(gt.ResponseCheck(response)) // prints "Success"
}
