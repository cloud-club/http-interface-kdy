package types

import (
	"fmt"
	"net/http"
)

func GetCommonHeader(request *http.Request) {
	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("cache-control", "max-age=0")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer "))
	request.Header.Set("X-GitHub-Api-Version", "2022-11-28")
}
