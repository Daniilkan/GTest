package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

func ResponseCheck(response *http.Response) string {
	if response.StatusCode == 200 {
		return "Success"
	}
	return "Fail"
}

func ResponseNotEmpty(response *http.Response) bool {
	data, _ := io.ReadAll(response.Body)
	if data != nil || len(data) > 0 {
		return true
	}
	return false
}

type ResponseType = responseResult
type responseResult int

const (
	TypeError responseResult = iota - 1
	TypeHtml
	TypeJson
)

func GetResponseType(response *http.Response) responseResult {
	data, _ := io.ReadAll(response.Body)
	var dataBody struct{}
	err := json.Unmarshal(data, &dataBody)
	if err == nil {
		return TypeJson
	}
	d := xml.NewDecoder(response.Body)

	// Configure the decoder for HTML; leave off strict and autoclose for XHTML
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		_, err := d.Token()
		switch err {
		case io.EOF:
			return TypeHtml // We're done, it's valid!
		}
	}
	return TypeError
}

func ResponseContains(response *http.Response, subSlice string) bool {
	data, _ := io.ReadAll(response.Body)
	return strings.Contains(string(data), string(subSlice))
}

func WebPageWorking(address string) bool {
	response, err := http.Get(address)
	if err != nil {
		return false
	}
	if response.StatusCode == 200 {
		return true
	}
	return false
}

func WebPageContains(address string, subSlice string) bool {
	response, err := http.Get(address)
	if err != nil {
		return false
	}
	data, _ := io.ReadAll(response.Body)
	return strings.Contains(string(data), string(subSlice))
}
