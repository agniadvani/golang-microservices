package services

import (
	"net/http"
	"sync"

	"github.com/agniadvani/golang-microservices/src/api/config"
	"github.com/agniadvani/golang-microservices/src/api/domain/github"
	"github.com/agniadvani/golang-microservices/src/api/providers/github_provider"

	"github.com/agniadvani/golang-microservices/src/api/domain/repositories"
	"github.com/agniadvani/golang-microservices/src/api/utilis/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (r *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiErr(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}
	return &result, nil
}

func (r *repoService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)
	var wg sync.WaitGroup
	go r.handleRequest(&wg, input, output)
	for _, currentRepo := range request {
		wg.Add(1)
		go r.createRepos(currentRepo, input)
	}
	wg.Wait()
	close(input)
	result := <-output

	successCreations := 0
	for _, current := range result.Result {
		if current.Response != nil {
			successCreations++
		}

	}

	if successCreations == 0 {
		result.StatusCode = result.Result[0].Error.Status()
	} else if successCreations != len(request) {
		result.StatusCode = http.StatusPartialContent
	} else {
		result.StatusCode = http.StatusCreated
	}

	return result, nil
}
func (r *repoService) handleRequest(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingRepo := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingRepo.Response,
			Error:    incomingRepo.Error,
		}
		results.Result = append(results.Result, repoResult)
		wg.Done()
	}
	output <- results
}

func (r *repoService) createRepos(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	result, err := r.CreateRepo(input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}
