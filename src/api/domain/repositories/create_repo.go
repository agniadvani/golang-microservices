package repositories

import (
	"strings"

	"github.com/agniadvani/golang-microservices/src/api/utilis/errors"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (input *CreateRepoRequest) Validate() errors.ApiError {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int                        `json:"status"`
	Result     []CreateRepositoriesResult `json:"result"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"response"`
	Error    errors.ApiError     `json:"error"`
}
