package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type searchResult struct {
	Name       string `json:"name"`
	Authors    string `json:"authors"`
	ProjectURI string `json:"project_uri"`
}

func search(query string) ([]*searchResult, error) {
	escapedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("https://rubygems.org/api/v1/search.json?query=%s", escapedQuery)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var results []*searchResult
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatal(err)
	}
	return results, nil
}
