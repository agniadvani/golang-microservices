package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMock = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	URL        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func FlushMockUp() {
	mocks = make(map[string]*Mock)
}
func getMockId(HttpMethod string, Url string) string {
	return fmt.Sprintf("%s_%s", HttpMethod, Url)
}

func StartMockUp() {
	enabledMock = true
}

func StopMockUp() {
	enabledMock = false
}

func AddMock(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.URL)] = &mock
}

func Post(url string, body interface{}, header http.Header) (*http.Response, error) {
	if enabledMock {
		mock := mocks[getMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = header
	client := http.Client{}
	return client.Do(request)
}
