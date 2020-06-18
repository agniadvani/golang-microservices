package github_provider

import (
	"errors"
	"testing"

	"github.com/agniadvani/golang-microservices/src/api/clients/restclient"
	"github.com/agniadvani/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepo(t *testing.T) {
	restclient.AddMock(restclient.Mock{
		URL: "https://api.github.com/user/repos",
		Err: errors.New("no mockup found for given request"),
	})
	restclient.StartMockUp()
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}
