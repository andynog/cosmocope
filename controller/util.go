package controller

import (
	"fmt"
	"net/url"
	"strings"
)

func ParseGithubURL(repo string) (owner string, name string, err error) {
	u, err := url.Parse(repo)
	if err != nil {
		return "", "", err
	}
	paths := strings.Split(u.Path, "/")
	if strings.ToLower(paths[0]) != "github.com" {
		return "", "", fmt.Errorf("invalid Github URL")
	}
	return paths[1], paths[2], nil
}