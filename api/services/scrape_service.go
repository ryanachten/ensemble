package services

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Link struct {
	Title string
	Url   *string
}

type BandMetadata struct {
	ImageUrl    string
	Members     []Link
	PastMembers []Link
}

type ArtistMetadata struct {
	ImageUrl   string
	MemberOf   []Link
	FormerlyOf []Link
}

func ScrapeBandMetadata(pageTitle string) BandMetadata {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	encodedTitle := url.QueryEscape(pageTitle)
	requestUrl := fmt.Sprintf("https://en.wikipedia.org/w/index.php?title=%s", encodedTitle)

	var metadata BandMetadata
	c.OnHTML(".infobox-image img", func(e *colly.HTMLElement) {
		metadata.ImageUrl = "https:" + e.Attr("src")
	})

	c.OnHTML(".infobox tr", func(e *colly.HTMLElement) {
		label := e.DOM.Find(".infobox-label").Text()
		if label == "Members" {
			metadata.Members = scrapeInfoBoxDataLinks(e)
		}
		if label == "Past members" {
			metadata.PastMembers = scrapeInfoBoxDataLinks(e)
		}
	})

	c.Visit(requestUrl)

	return metadata
}

func ScrapeArtistMetadata(requestUrl string) ArtistMetadata {
	// TODO: do we want to store the collector as a variable to prevent instating each time?
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	var metadata ArtistMetadata

	c.OnHTML(".infobox-image img", func(e *colly.HTMLElement) {
		metadata.ImageUrl = "https:" + e.Attr("src")
	})

	c.OnHTML(".infobox tr", func(e *colly.HTMLElement) {
		label := e.DOM.Find(".infobox-label").Text()
		if label == "Member of" {
			metadata.MemberOf = scrapeInfoBoxDataLinks(e)
		}
		if label == "Formerly of" {
			metadata.FormerlyOf = scrapeInfoBoxDataLinks(e)
		}
	})

	c.Visit(requestUrl)

	return metadata
}

func scrapeInfoBoxDataLinks(element *colly.HTMLElement) []Link {
	var links []Link
	element.DOM.Find(".infobox-data li").Each(func(i int, s *goquery.Selection) {
		url, urlExists := s.Find("a").Attr("href")
		link := Link{
			Title: s.Text(),
		}
		url = element.Request.AbsoluteURL(url)
		if urlExists {
			link.Url = &url
		}
		links = append(links, link)
	})
	return links
}
