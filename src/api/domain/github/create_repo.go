package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProject  bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPush bool `json:"push"`
	HasPull bool `json:"pull"`
}
