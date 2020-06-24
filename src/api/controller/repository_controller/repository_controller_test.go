package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/agniadvani/golang-microservices/src/api/domain/repositories"

	"github.com/stretchr/testify/assert"

	"github.com/agniadvani/golang-microservices/src/api/clients/restclient"
	"github.com/agniadvani/golang-microservices/src/api/utilis/errors"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	restclient.StartMockUp()
	os.Exit(m.Run())
}
func TestCreateRepoJsonErr(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	c.Request = request
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestCreateRepoAuthorizationErr(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	c.Request = (request)
	restclient.FlushMockUp()

	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
			StatusCode: http.StatusUnauthorized,
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)
	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, "Requires authentication", apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	c.Request = (request)

	restclient.FlushMockUp()

	restclient.AddMock(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name":"testing","owner":{"login":"testname"}}`)),
			StatusCode: http.StatusCreated,
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)
	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, result.ID, 123)
	assert.EqualValues(t, result.Name, "testing")
	assert.EqualValues(t, result.Owner, "testname")

}
