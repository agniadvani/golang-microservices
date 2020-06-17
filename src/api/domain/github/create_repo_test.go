package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoAsJson(t *testing.T) {
	request := CreteRepoRequest{
		Name:        "Postman-Repo",
		Description: "This is your first Api repository",
		Homepage:    "https://github.com",
		Private:     false,
		HasIssues:   true,
		HasProject:  true,
		HasWiki:     true,
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreteRepoRequest

	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.EqualValues(t, request.Name, target.Name)
	assert.EqualValues(t, request.HasWiki, target.HasWiki)
	assert.EqualValues(t, request.HasProject, target.HasProject)
}
