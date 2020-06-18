package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	enabledMock = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	URL      string
	Response *http.Response
	Err      error
}

func StartMockUp() {
	enabledMock = true
}

func StopMockUp() {
	enabledMock = false
}

func AddMock(mock Mock) {
	mocks[mock.URL] = &mock
}

func Post(url string, body interface{}, header http.Header) (*http.Response, error) {
	if enabledMock {
		mock := mocks[url]
		if mocks[url] == nil {
			return nil, errors.New("invalid response type")
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
