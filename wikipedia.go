package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/patmigliaccio/go-wikimedia"
)

// A WikipediaAPIClient makes calls to Wikipedia for content.
type WikipediaAPIClient interface {
	GetExtracts(titles []string) ([]WikipediaPageFull, error)
	GetPrefixResults(pfx string, limit int) ([]WikipediaPage, error)
	GetCategories(pageid int) (WikipediaPageFull, error)
	GetSections(pageid int) (WikipediaPageFull, error)
}

// A WikipediaPage returned from the API
type WikipediaPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// A WikipediaPageFull with extract returned from the API
type WikipediaPageFull struct {
	Meta       WikipediaPage `json:"metadata"`
	Extract    string        `json:"extract"`
	Categories []string      `json:"categories"`
	Sections   []string      `json:"sections"`
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
			URL:   getWikipediaURL(p.Title),
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
				URL:   getWikipediaURL(p.Title),
			},
			Extract: p.Extract,
		})
	}

	return values, nil
}

// GetCategories retrieves the categories associated with a specified Wikipedia article.
func (wk *wkAPI) GetCategories(pageid int) (WikipediaPageFull, error) {
	var value WikipediaPageFull

	f := url.Values{
		"action": {"parse"},
		"pageid": {strconv.Itoa(pageid)},
		"prop":   {"categories"},
	}

	res, err := wk.w.Query(f)
	if err != nil {
		return value, err
	}

	value = WikipediaPageFull{
		Meta: WikipediaPage{
			ID:    res.Parse.PageId,
			Title: res.Parse.Title,
			URL:   getWikipediaURL(res.Parse.Title),
		},
		Categories: getCategoryNames(res.Parse.Categories),
	}

	return value, nil
}

// GetSections retrieves the sections within a specified Wikipedia article.
func (wk *wkAPI) GetSections(pageid int) (WikipediaPageFull, error) {
	var value WikipediaPageFull

	f := url.Values{
		"action": {"parse"},
		"pageid": {strconv.Itoa(pageid)},
		"prop":   {"sections"},
	}

	res, err := wk.w.Query(f)
	if err != nil {
		return value, err
	}

	value = WikipediaPageFull{
		Meta: WikipediaPage{
			ID:    res.Parse.PageId,
			Title: res.Parse.Title,
			URL:   getWikipediaURL(res.Parse.Title),
		},
		Sections: getSectionAnchors(res.Parse.Sections),
	}

	return value, nil
}

// getWikipediaURL returns a Wikipedia URL from a `wikimedia.ApiPage`
func getWikipediaURL(pageTitle string) string {
	return fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.Replace(pageTitle, " ", "_", -1))
}

// getCategoryNames returns the string name from a `wikimedia.ApiPageCategory`
func getCategoryNames(categories []wikimedia.ApiPageCategory) []string {
	var values []string
	for _, cat := range categories {
		values = append(values, cat.Name)
	}

	return values
}

// getSectionAnchors returns the anchor name from a `wikimedia.ApiPageSection`
func getSectionAnchors(sections []wikimedia.ApiPageSection) []string {
	var values []string
	for _, section := range sections {
		values = append(values, section.Anchor)
	}

	return values
}
