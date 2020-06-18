package github_provider

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/agniadvani/golang-microservices/src/api/clients/restclient"
	"github.com/agniadvani/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockUp()
	os.Exit(m.Run())
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

}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
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

}
