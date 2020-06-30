package services

import (
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
	go r.handleRequest(input, output)
	for _, currentRepo := range request {
		go r.createRepos(currentRepo, input)
	}
	result := <-output
	return result, nil
}
func (r *repoService) handleRequest(input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingRepo := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingRepo.Response,
			Error:    incomingRepo.Error,
		}
		results.Result = append(results.Result, repoResult)
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
