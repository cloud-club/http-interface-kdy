package internal

import (
	"fmt"
	"github.com/cloud-club/http-interface-kdy/pkg"
	"github.com/cloud-club/http-interface-kdy/types"
	"log"
)

const GITHUB_URL = "https://api.github.com"

type Client struct {
	Github *types.GithubClient
}

func NewClientService() *Client {
	client := &Client{Github: types.NewGithubClientClient()}

	return client
}

func Run() {
	client := NewClientService()
	resp, err := client.Github.Get(GITHUB_URL)
	defer resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("error : %v", err))
	}
	log.Printf("status code : %d\n", resp.StatusCode)
	log.Println()
	resp, err = client.Github.Get("https://api.github.com/repos/cloud-club/http-interface-kdy")
	defer resp.Body.Close()
	if err != nil {
		panic(fmt.Sprintf("error : %v", err))
	}

	log.Printf("status code : %d\n", resp.StatusCode)
	githubRepo, err := pkg.GetGithubInfo(resp)
	if err != nil {
		panic(err)
	}
	log.Println("repo name : ", githubRepo.Name)
	log.Println("repo full name : ", githubRepo.FullName)
	log.Println("repo owner : ", githubRepo.Owner.Login)
	log.Println("owner url : ", githubRepo.Owner.Url)
}
