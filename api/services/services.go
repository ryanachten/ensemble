package services

import (
	clients "ensemble/clients"
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Link struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type BandMetadata struct {
	Members     []Link `json:"members"`
	PastMembers []Link `json:"pastMembers"`
}

func GetBandMetadata(name string) *BandMetadata {
	searchResults, err := clients.GetSearchResults(name)
	if err != nil {
		return nil
	}
	metadata := ScrapeBandMetadata(searchResults[0].Title)
	return &metadata
}

func ScrapeBandMetadata(pageTitle string) BandMetadata {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	encodedTitle := url.QueryEscape(pageTitle)
	requestUrl := fmt.Sprintf("https://en.wikipedia.org/w/index.php?title=%s", encodedTitle)

	var metadata BandMetadata

	c.OnHTML(".infobox tr", func(e *colly.HTMLElement) {
		label := e.DOM.Find(".infobox-label").Text()
		if label == "Members" {
			metadata.Members = ScrapeInfoBoxDataLinks(e)
		}
		if label == "Past members" {
			metadata.PastMembers = ScrapeInfoBoxDataLinks(e)
		}
	})

	c.Visit(requestUrl)

	return metadata
}

func ScrapeInfoBoxDataLinks(element *colly.HTMLElement) []Link {
	var links []Link
	element.DOM.Find(".infobox-data li").Each(func(i int, s *goquery.Selection) {
		url, urlExists := s.Find("a").Attr("href")
		link := Link{
			Title: s.Text(),
		}
		if urlExists {
			link.Url = element.Request.AbsoluteURL(url)
		}
		links = append(links, link)
	})
	return links
}
