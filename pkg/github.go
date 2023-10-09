package pkg

import (
	"encoding/json"
	"io"
	"net/http"
)

type GithubRepo struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login     string `json:"login"`
		Id        int    `json:"id"`
		Url       string `json:"url"`
		SiteAdmin bool   `json:"site_admin"`
	} `json:"owner"`
}

func GetGithubInfo(resp *http.Response) (*GithubRepo, error) {
	repo := &GithubRepo{}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, repo); err != nil {
		return nil, err
	}
	return repo, nil
}
