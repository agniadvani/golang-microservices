package services

import (
	"strings"

	"github.com/agniadvani/golang-microservices/src/api/config"
	"github.com/agniadvani/golang-microservices/src/api/domain/github"
	"github.com/agniadvani/golang-microservices/src/api/providers/github_provider"

	"github.com/agniadvani/golang-microservices/src/api/domain/repositories"
	"github.com/agniadvani/golang-microservices/src/api/utilis/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (r *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
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
