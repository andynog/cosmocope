package client

import (
	"encoding/json"
	"fmt"
	"github.com/hako/durafmt"
	_ "github.com/hako/durafmt"
	"github.com/rogpeppe/go-internal/modfile"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	repoTendermint = "github.com/tendermint/tendermint"
	repoCosmosSDK = "github.com/cosmos/cosmos-sdk"
	repoIBC = "github.com/cosmos/ibc-go"
)

type RateLimitError struct {
	Remaining int64
}

func (e *RateLimitError) Error() string {
	if e.Remaining == 0 {
		return "rate limit reached, please try again later..."
	} else {
		//reset := fmt.Sprintf("%s", time.Unix(e.Remaining, 0))
		remainingTime := time.Unix(e.Remaining, 0)
		until := time.Until(remainingTime)
		return fmt.Sprintf("rate limit reached, please try again in %s", durafmt.Parse(until).LimitFirstN(2))
	}
}

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

	res, _ := client.Do(req)
	return res.StatusCode == http.StatusOK
}

// Function that calls the Github Search API and look for projects
// that contain the topic 'cosmos-sdk'. This assumes that owners of
// Cosmos SDK project add the 'cosmos-sdk' to their projects for better
// discovery
func SearchGithub(topic string) (result GithubSearchResult, err error) {
	var searchRslt GithubSearchResult
	url := "https://api.github.com/search/repositories?q=topic:" + topic + "&page=1&per_page=1000&sort:updated"
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Accept", "application/vnd.github.mercy-preview+json")
	if err != nil {
		fmt.Println(err)
	}
	// Authenticate the API call
	id, secret := GetCredentials()
	req.SetBasicAuth(id,secret)

	res, err := client.Do(req)

	// Check if rate limit reached
	if (err != nil) || (res.StatusCode != 200)  {
		val, ok := res.Header["X-Ratelimit-Remaining"]
		if ok && (val[0] == "0") {
			val, ok := res.Header["X-Ratelimit-Reset"]
			if ok {
				remaining, err := strconv.Atoi(val[0])
				if err != nil {
					return searchRslt, &RateLimitError{	Remaining: 0 }
				} else {
					return searchRslt, &RateLimitError{	Remaining: int64(remaining) }
				}
			}
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &searchRslt)
	if err != nil {
		return searchRslt, err
	} else {
		return searchRslt, nil
	}
}


// Function that calls the Github API to retrieve contents
// (files and folders) information from a Github repo that
// contains a folder named 'x' where modules are stored.
func GetContentFromGithub(owner string, repo string) (result GithubContentResult, err error) {
	var contentResult GithubContentResult

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/contents/x?ref=master"
	method := "GET"

	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// Authenticate the API call
	id, secret := GetCredentials()
	req.SetBasicAuth(id,secret)

	res, err := client.Do(req)

	// Check if rate limit reached
	if (err != nil) || (res.StatusCode != 200)  {
		val, ok := res.Header["X-Ratelimit-Remaining"]
		if ok && (val[0] == "0") {
			val, ok := res.Header["X-Ratelimit-Reset"]
			if ok {
				remaining, err := strconv.Atoi(val[0])
				if err != nil {
					err = &RateLimitError{ Remaining: 0 }
					fmt.Println("\r\n", err)
					os.Exit(1)
				} else {
					err = &RateLimitError{ Remaining: int64(remaining) }
					fmt.Println("\r\n", err)
					os.Exit(1)
				}
			}
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusNotFound {
		return contentResult, nil
	}
	err = json.Unmarshal(body, &contentResult)
	if err != nil {
		return nil, err
	}
	return contentResult, nil
}

// Function that calls the Github API to retrieve releases
// from a Github repo
func GetReleasesFromGithub(owner string, repo string) (result GithubReleasesResult, err error) {
	var releaseResult GithubReleasesResult
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/releases"

	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Accept", "application/vnd.github.mercy-preview+json")
	if err != nil {
		fmt.Println(err)
	}
	// Authenticate the API call
	id, secret := GetCredentials()
	req.SetBasicAuth(id,secret)

	res, err := client.Do(req)

	// Check if rate limit reached
	if (err != nil) || (res.StatusCode != 200)  {
		val, ok := res.Header["X-Ratelimit-Remaining"]
		if ok && (val[0] == "0") {
			val, ok := res.Header["X-Ratelimit-Reset"]
			if ok {
				remaining, err := strconv.Atoi(val[0])
				if err != nil {
					return releaseResult, &RateLimitError{	Remaining: 0 }
				} else {
					return releaseResult, &RateLimitError{	Remaining: int64(remaining) }
				}
			}
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &releaseResult)
	if err != nil {
		return releaseResult, err
	} else {
		return releaseResult, nil
	}
}

// Function that fetches the go.mod from a repository
// If it finds a dependency.go on Cosmos SDK or Tendermint
// return the versions
func GetDependencies(owner string, repo string, branch string) (sdk string, tendermint string, ibc string, err error) {
	sdkVersion := ""
	tendermintVersion := ""
	url := "https://raw.githubusercontent.com/" + owner + "/" + repo + "/" + branch + "/go.mod"
	method := "GET"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)
	res, err := client.Do(req)

	// Check if rate limit reached
	if (err != nil) || (res.StatusCode != 200) {
		return "", "", "", err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", "", err
	}
	if res.StatusCode == http.StatusNotFound {
		return "", "", "", nil
	}

	f, err := modfile.Parse(url, data, nil)
	if err != nil {
		// ignore err
		return "", "", "", err
	}

	for _, r := range f.Require {
		if strings.ToLower(r.Mod.Path) == repoCosmosSDK {
			sdkVersion = r.Mod.Version
		}
		if strings.ToLower(r.Mod.Path) == repoTendermint {
			tendermintVersion = r.Mod.Version
		}
		if strings.ToLower(r.Mod.Path) == "repoIBC" {
			ibc = r.Mod.Version
		}
	}
	return sdkVersion, tendermintVersion, ibc,nil
}

func GetCredentials() (clientID string, clientSecret string) {
	// Get the CC_GITHUB_CLIENT_ID environment variable
	var found bool
	clientID, found = os.LookupEnv("CC_GITHUB_CLIENT_ID")
	if !found {
		fmt.Println("Failed to retrieve Github Client ID. Please set the 'CC_GITHUB_CLIENT_ID' environment variable")
		os.Exit(1)
	}
	// Get the CC_GITHUB_CLIENT_SECRET environment variable
	clientSecret, found = os.LookupEnv("CC_GITHUB_CLIENT_SECRET")
	if !found {
		fmt.Println("Failed to retrieve Github Client Secret. Please set the 'CC_GITHUB_CLIENT_SECRET' environment variable")
		os.Exit(1)
	}
	return clientID, clientSecret
}