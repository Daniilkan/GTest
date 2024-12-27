package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

func ResponseCheck(response *http.Response) string {
	if response.StatusCode == http.StatusOK {
		return "Success"
	}
	return "Fail"
}

func ResponseNotEmpty(response *http.Response) bool {
	data, _ := io.ReadAll(response.Body)
	return len(data) > 0
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
	if isJSON(data) {
		return TypeJson
	}
	if isHTML(response.Body) {
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

func ResponseContains(response *http.Response, subSlice string) bool {
	data, _ := io.ReadAll(response.Body)
	return strings.Contains(string(data), subSlice)
}

func WebPageWorking(address string) bool {
	response, err := http.Get(address)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == http.StatusOK
}

func WebPageContains(address string, subSlice string) bool {
	response, err := http.Get(address)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)
	return strings.Contains(string(data), subSlice)
}
