package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/patrickmn/go-wikimedia"
)

// A WikipediaAPIClient makes calls to Wikipedia for content.
type WikipediaAPIClient interface {
	GetExtracts(titles []string) ([]WikipediaPageFull, error)
	GetPrefixResults(pfx string, limit int) ([]WikipediaPage, error)
}

// A WikipediaPage returned from the API
type WikipediaPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// A WikipediaPageFull with extract returned from the API
type WikipediaPageFull struct {
	Meta    WikipediaPage `json:"metadata"`
	Extract string        `json:"extract"`
}

type wkAPI struct {
	w *wikimedia.Wikimedia
}

// NewWikipediaClient instantiates an instance of the WikipediaAPIClient.
func NewWikipediaClient() (WikipediaAPIClient, error) {
	w, err := wikimedia.New("https://en.wikipedia.org/w/api.php")
	if err != nil {
		return nil, err
	}

	return &wkAPI{
		w: w,
	}, nil
}

// GetPrefixResults retrieves a list of Wikipedia pages based on a query string
func (wk *wkAPI) GetPrefixResults(pfx string, limit int) ([]WikipediaPage, error) {
	if limit == 0 {
		limit = 50
	}

	f := url.Values{
		"action":       {"query"},
		"generator":    {"prefixsearch"},
		"prop":         {"pageprops|pageimages|description"},
		"ppprop":       {"displaytitle"},
		"gpssearch":    {pfx},
		"gpsnamespace": {"0"},
		"gpslimit":     {strconv.Itoa(limit)},
	}

	res, err := wk.w.Query(f)
	if err != nil {
		return nil, err
	}

	var values []WikipediaPage
	for _, p := range res.Query.Pages {
		values = append(values, WikipediaPage{
			ID:    p.PageId,
			Title: p.Title,
			URL:   getWikipediaURL(p),
		})
	}

	return values, nil
}

// GetExtracts retrieves the extracts for a given list of titles.
func (wk *wkAPI) GetExtracts(titles []string) ([]WikipediaPageFull, error) {
	f := url.Values{
		"action": {"query"},
		"prop":   {"extracts"},
		"titles": {strings.Join(titles[:], "|")},
	}
	res, err := wk.w.Query(f)
	if err != nil {
		return nil, err
	}

	var values []WikipediaPageFull
	for _, p := range res.Query.Pages {
		values = append(values, WikipediaPageFull{
			Meta: WikipediaPage{
				ID:    p.PageId,
				Title: p.Title,
				URL:   getWikipediaURL(p),
			},
			Extract: p.Extract,
		})
	}

	return values, nil
}

// getWikipediaURL returns a Wikipedia URL from a `wikimedia.ApiPage`
func getWikipediaURL(p wikimedia.ApiPage) string {
	return fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.Replace(p.Title, " ", "_", -1))
}
