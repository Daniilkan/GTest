package http

import (
	"bytes"
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
func ResponseContains(response *http.Response, seq []byte) (bool, error) {
	buff := make([]byte, len(seq))

	n, err := response.Body.Read(buff)
	if err != nil && err != io.EOF {
		return false, err
	}
	if n != len(seq) {
		return false, nil
	}

	for {
		if bytes.Equal(seq, buff) {
			return true, nil
		}
		buff = append(buff[1:], 0)
		_, err := response.Body.Read(buff[len(buff)-1:])
		if err != nil {
			return false, nil
		}
	}
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
func WebPageContains(address string, seq []byte) (bool, error) {
	response, err := http.Get(address)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	buff := make([]byte, len(seq))

	n, err := response.Body.Read(buff)
	if err != nil && err != io.EOF {
		return false, err
	}
	if n != len(seq) {
		return false, nil
	}

	for {
		if bytes.Equal(seq, buff) {
			return true, nil
		}
		buff = append(buff[1:], 0)
		_, err := response.Body.Read(buff[len(buff)-1:])
		if err != nil {
			return false, nil
		}
	}
}
