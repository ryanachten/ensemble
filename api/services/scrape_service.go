package services

import (
	"regexp"

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

func ScrapeBandMetadata(requestUrl string) BandMetadata {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

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
		url, hasUrl := s.Find("a").Attr("href")
		links = append(links, getLink(element, s, url, hasUrl))
	})

	// Handle cases where the links are not formatted in a list
	if len(links) == 0 {
		element.DOM.Find(".infobox-data a").Each(func(i int, s *goquery.Selection) {
			url, hasUrl := s.Attr("href")
			links = append(links, getLink(element, s, url, hasUrl))
		})
	}
	return links
}

func getLink(element *colly.HTMLElement, s *goquery.Selection, url string, hasUrl bool) Link {
	// Remove footnotes from link title
	footnoteRegex := regexp.MustCompile(`\[\d*\]`)
	title := footnoteRegex.ReplaceAllString(s.Text(), "")

	link := Link{
		Title: title,
	}

	if hasUrl {
		url = element.Request.AbsoluteURL(url)
		link.Url = &url
	}

	return link
}
