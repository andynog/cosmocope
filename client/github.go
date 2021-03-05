package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Function to check if a git repository contains a folder named 'x'
// which indicates it is a Cosmos SDK project
func LookForModules(repo string) bool {
	url := repo + "/tree/master/x"
	method := "GET"
	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if res.StatusCode == http.StatusOK {
		return true
	}
	return false
}

// Function that calls the Github Search API and look for projects
// that contain the topic 'cosmos-sdk'. This assumes that owners of
// Cosmos SDK project add the 'cosmos-sdk' to their projects for better
// discovery
func SearchGithub(topic string) GithubSearchResult {
	url := "https://api.github.com/search/repositories?q=topic:" + topic + "&page=1&per_page=1000&sort:updated"
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Accept", "application/vnd.github.mercy-preview+json")
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var searchRslt GithubSearchResult
	err = json.Unmarshal(body, &searchRslt)
	if err != nil {
		log.Println(err)
	}
	return searchRslt
}


// Function that calls the Github API to retrieve contents
// (files and folders) information from a Github repo that
// contains a folder named 'x' where modules are stored.
func GetContentFromGithub(owner string, repo string) (result GithubContentResult, err error) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/contents/x?ref=master"
	method := "GET"

	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("no modules found")
	}
	var contentResult GithubContentResult
	err = json.Unmarshal(body, &contentResult)
	if err != nil {
		return nil, err
	}
	return contentResult, nil
}
