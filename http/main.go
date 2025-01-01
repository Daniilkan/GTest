package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

// ResponseCheck checks if response status code is 200
func ResponseCheck(response *http.Response) string {
	if response.StatusCode == http.StatusOK {
		return "Success"
	}
	return "Fail"
}

// ResponseNotEmpty checks if response body is not empty
func ResponseNotEmpty(response *http.Response) bool {
	data, _ := io.ReadAll(response.Body)
	response.Body = io.NopCloser(strings.NewReader(string(data))) // Reset Body
	return len(data) > 0
}

// GetResponseType returns type of response
type ResponseType = responseResult
type responseResult int

const (
	TypeError responseResult = iota - 1
	TypeHtml
	TypeJson
)

func GetResponseType(response *http.Response) responseResult {
	data, _ := io.ReadAll(response.Body)
	response.Body = io.NopCloser(strings.NewReader(string(data))) // Reset Body
	if isJSON(data) {
		return TypeJson
	}
	if isHTML(io.NopCloser(strings.NewReader(string(data)))) {
		return TypeHtml
	}
	return TypeError
}

func isJSON(data []byte) bool {
	var js struct{}
	return json.Unmarshal(data, &js) == nil
}

func isHTML(body io.Reader) bool {
	d := xml.NewDecoder(body)
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		_, err := d.Token()
		if err == io.EOF {
			return true
		}
		if err != nil {
			return false
		}
	}
}

// ResponseContains checks if response contains subSlice
func ResponseContains(response *http.Response, subSlice string) bool {
	data, _ := io.ReadAll(response.Body)
	response.Body = io.NopCloser(strings.NewReader(string(data))) // Reset Body
	return strings.Contains(string(data), subSlice)
}

// WebPageWorking checks if webpage is working
func WebPageWorking(address string) bool {
	response, err := http.Get(address)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == http.StatusOK
}

// WebPageContains checks if webpage contains subSlice
func WebPageContains(address string, subSlice string) bool {
	response, err := http.Get(address)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)
	return strings.Contains(string(data), subSlice)
}
