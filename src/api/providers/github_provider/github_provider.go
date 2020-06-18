package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/agniadvani/golang-microservices/src/api/clients/restclient"
	"github.com/agniadvani/golang-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {

	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := restclient.Post(urlCreateRepo, request, headers)
	fmt.Println(response)
	fmt.Println(err)
	if err != nil {
		log.Printf("error when trying to create a new repository in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()
	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if json.Unmarshal(bytes, &result); err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json body response"}
	}
	return &result, nil
}
