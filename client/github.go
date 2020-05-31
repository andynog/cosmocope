package client

import (
	"encoding/json"
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
	client := &http.Client {
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
	url := "https://api.github.com/search/repositories?q=topic:" + topic + "&page=1&per_page=1000"
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