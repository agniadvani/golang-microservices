package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/agniadvani/golang-microservices/src/api/domain/repositories"

	"github.com/agniadvani/golang-microservices/src/api/clients/restclient"
)

func TestMain(m *testing.M) {

	restclient.StartMockUp()
	os.Exit(m.Run())
}

func TestCreateRepoNameError(t *testing.T) {

	testStruct := repositories.CreateRepoRequest{
		Name:        "",
		Description: "",
	}

	response, err := RepoService.CreateRepo(testStruct)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoAuthenticationError(t *testing.T) {

	restclient.FlushMockUp()
	testStruct := repositories.CreateRepoRequest{
		Name:        "test",
		Description: "",
	}
	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
			StatusCode: http.StatusUnauthorized,
		},
	})

	response, err := RepoService.CreateRepo(testStruct)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockUp()
	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "Postman-Repo","owner":{"login":"agniadvani"}}`)),
			StatusCode: http.StatusCreated,
		},
	})
	testStruct := repositories.CreateRepoRequest{
		Name:        "test",
		Description: "",
	}
	response, err := RepoService.CreateRepo(testStruct)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, response.ID)
	assert.EqualValues(t, "Postman-Repo", response.Name)
	assert.EqualValues(t, "agniadvani", response.Owner)
}
