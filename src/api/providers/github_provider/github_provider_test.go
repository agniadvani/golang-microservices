package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/agniadvani/golang-microservices/src/api/clients/restclient"
	"github.com/agniadvani/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockUp()
	os.Exit(m.Run())
}
func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)

}
func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.FlushMockUp()
	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, "invalid restclient response", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMockUp()
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)

}
func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
			StatusCode: http.StatusUnauthorized,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123"}`)),
			StatusCode: http.StatusCreated,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal github create repo response", err.Message)
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "Postman-Repo","full_name": "agniadvani/Postman-Repo"}`)),
			StatusCode: http.StatusCreated,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.ID)
	assert.EqualValues(t, "Postman-Repo", response.Name)
	assert.EqualValues(t, "agniadvani/Postman-Repo", response.FullName)

}
