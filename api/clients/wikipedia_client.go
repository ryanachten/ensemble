package clients

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type WikipediaSearchItem struct {
	Title  string `json:"title"`
	PageId int    `json:"pageid"`
}

type WikipediaQuery struct {
	Search []WikipediaSearchItem `json:"search"`
}

type WikipediaSearchResult struct {
	Query WikipediaQuery `json:"query"`
}

func GetSearchResults(searchTerm string) ([]WikipediaSearchItem, error) {
	encodedSearchQuery := url.QueryEscape(searchTerm)
	requestUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch=%s&format=json", encodedSearchQuery)
	res, err := GetRequest(requestUrl)
	if err != nil {
		return nil, err
	}

	var result WikipediaSearchResult
	err = json.NewDecoder(strings.NewReader(string(res))).Decode(&result)
	return result.Query.Search, err
}
