// Package client provides functions to interact with the Github API
package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hako/durafmt"
	"golang.org/x/mod/modfile"
)

// Add constant for HTTP GET
const (
	GET = "GET"
)

// RateLimitError is a custom error type that represents a rate limit error and stores
// the remaining time until the rate limit resets.
type RateLimitError struct {
	Remaining int64
}

func (e *RateLimitError) Error() string {
	if e.Remaining == 0 {
		return "rate limit reached, please try again later..."
	} 

	remainingTime := time.Unix(e.Remaining, 0)
	until := time.Until(remainingTime)
	return fmt.Sprintf("rate limit reached, please try again in %s", durafmt.Parse(until).LimitFirstN(2))
	
}

// LookForModules checks if a given repository contains a folder named 'x' which indicates it is a Cosmos SDK project
func LookForModules(repo string) bool {
	url := repo + "/tree/master/x"
	client := &http.Client{
	}
	req, err := http.NewRequest(GET, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, _ := client.Do(req)
	return res.StatusCode == http.StatusOK
}

// SearchGithub calls the Github Search API and look for projects
// that contain the topic 'cosmos-sdk'. This assumes that owners of
// Cosmos SDK project add the 'cosmos-sdk' to their projects for better
// discovery
func SearchGithub(topic string) (result GithubSearchResult, err error) {
	var searchRslt GithubSearchResult
	url := "https://api.github.com/search/repositories?q=topic:" + topic + "&page=1&per_page=1000&sort:updated"
	client := &http.Client{
	}
	req, err := http.NewRequest(GET, url, nil)
	req.Header.Add("Accept", "application/vnd.github.mercy-preview+json")
	if err != nil {
		fmt.Println(err)
	}
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
				} 
				return searchRslt, &RateLimitError{	Remaining: int64(remaining) }
			}
		}
	}

	defer func() {
        closeErr := res.Body.Close()
        if closeErr != nil {
            // If the function wasn't already returning an error,
            // assign the close error to the named return 'err'.
            if err == nil {
                err = fmt.Errorf("error closing response body: %w", closeErr)
            } else {
                // Optionally log or wrap if another error already occurred
                log.Printf("Error closing response body (original error: %v): %v", err, closeErr)
            }
        }
    }()
	
	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &searchRslt)
	if err != nil {
		return searchRslt, err
	} 
	
	return searchRslt, nil
	
}


// GetContentFromGithub calls the Github API to retrieve contents (files and folders) information from a Github repo that
// contains a folder named 'x' where modules are stored.
func GetContentFromGithub(owner string, repo string) (result GithubContentResult, err error) {
	var contentResult GithubContentResult

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/contents/x?ref=master"

	client := &http.Client {}
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, err
	}

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
				} 
				err = &RateLimitError{ Remaining: int64(remaining) }
				fmt.Println("\r\n", err)
				os.Exit(1)		
			}
		}
	}

	defer func() {
        closeErr := res.Body.Close()
        if closeErr != nil {
            // If the function wasn't already returning an error,
            // assign the close error to the named return 'err'.
            if err == nil {
                err = fmt.Errorf("error closing response body: %w", closeErr)
            } else {
                // Optionally log or wrap if another error already occurred
                log.Printf("Error closing response body (original error: %v): %v", err, closeErr)
            }
        }
    }()

	body, err := io.ReadAll(res.Body)
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

// GetReleasesFromGithub calls the Github API to retrieve releases from a Github repo
func GetReleasesFromGithub(owner string, repo string) (result GithubReleasesResult, err error) {
	var releaseResult GithubReleasesResult
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/releases"

	client := &http.Client{
	}
	req, err := http.NewRequest(GET, url, nil)
	req.Header.Add("Accept", "application/vnd.github.mercy-preview+json")
	if err != nil {
		fmt.Println(err)
	}
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
				} 
				return releaseResult, &RateLimitError{	Remaining: int64(remaining) }				
			}
		}
	}

	defer func() {
        closeErr := res.Body.Close()
        if closeErr != nil {
            // If the function wasn't already returning an error,
            // assign the close error to the named return 'err'.
            if err == nil {
                err = fmt.Errorf("error closing response body: %w", closeErr)
            } else {
                // Optionally log or wrap if another error already occurred
                log.Printf("Error closing response body (original error: %v): %v", err, closeErr)
            }
        }
    }()
	
	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &releaseResult)
	if err != nil {
		return releaseResult, err
	} 
		
	return releaseResult, nil
	
}

// IsCosmosSDK fetches the go.mod from a repository
// If it finds a dependency on Cosmos SDK return a
// boolean value indicating this is a project that
// uses the Cosmos SDK
func IsCosmosSDK(owner string, repo string, branch string) (result string, err error) {
	url := "https://raw.githubusercontent.com/" + owner + "/" + repo + "/" + branch + "/go.mod"
	client := &http.Client {}
	req, err := http.NewRequest(GET, url, nil)
	if (err != nil) {
		return "", err
	}
	res, err := client.Do(req)

	// Check if rate limit reached
	if (err != nil) || (res.StatusCode != 200) {
		return "", err
	}

	defer func() {
        closeErr := res.Body.Close()
        if closeErr != nil {
            // If the function wasn't already returning an error,
            // assign the close error to the named return 'err'.
            if err == nil {
                err = fmt.Errorf("error closing response body: %w", closeErr)
            } else {
                // Optionally log or wrap if another error already occurred
                log.Printf("Error closing response body (original error: %v): %v", err, closeErr)
            }
        }
    }()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode == http.StatusNotFound {
		return "", nil
	}

	f, err := modfile.Parse(url, data, nil)
	if err != nil {
		// ignore err
		return "", err
	}

	for _, r := range f.Require {
		if strings.ToLower(r.Mod.Path) == "github.com/cosmos/cosmos-sdk" {
			return r.Mod.Version, nil
		}
	}
	return "", nil
}