package models

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
	Genres      []Link
	Members     []Link
	PastMembers []Link
}

type ArtistMetadata struct {
	ImageUrl   string
	MemberOf   []Link
	FormerlyOf []Link
}

type GenreMetadata struct {
	ImageUrl         string
	StylisticOrigins []Link
	DerivativeForms  []Link
	Subgenres        []Link
	FusionGenres     []Link
}

// Web-scraper for obtaining information from Wikipedia pages
type WikiScraper struct {
	Domains colly.CollectorOption
}

func NewWikiScraper() WikiScraper {
	return WikiScraper{
		Domains: colly.AllowedDomains("en.wikipedia.org"),
	}
}

func (scraper *WikiScraper) GetBandMetadata(requestUrl string) BandMetadata {
	collector := colly.NewCollector(scraper.Domains)
	var metadata BandMetadata

	collector.OnHTML(".infobox-image img", func(e *colly.HTMLElement) {
		metadata.ImageUrl = "https:" + e.Attr("src")
	})

	collector.OnHTML(".infobox tr", func(e *colly.HTMLElement) {
		label := e.DOM.Find(".infobox-label").Text()
		if label == "Members" {
			metadata.Members = scrapeInfoBoxDataLinks(e)
		}
		if label == "Past members" {
			metadata.PastMembers = scrapeInfoBoxDataLinks(e)
		}
		if label == "Genres" {
			metadata.Genres = scrapeInfoBoxDataLinks(e)
		}
	})

	collector.Visit(requestUrl)

	return metadata
}

func (scraper *WikiScraper) GetArtistMetadata(requestUrl string) ArtistMetadata {
	collector := colly.NewCollector(scraper.Domains)
	var metadata ArtistMetadata

	collector.OnHTML(".infobox-image img", func(e *colly.HTMLElement) {
		metadata.ImageUrl = "https:" + e.Attr("src")
	})

	collector.OnHTML(".infobox tr", func(e *colly.HTMLElement) {
		label := e.DOM.Find(".infobox-label").Text()
		if label == "Member of" {
			metadata.MemberOf = scrapeInfoBoxDataLinks(e)
		}
		if label == "Formerly of" {
			metadata.FormerlyOf = scrapeInfoBoxDataLinks(e)
		}
	})

	collector.Visit(requestUrl)

	return metadata
}

func (scraper *WikiScraper) GetGenreMetadata(requestUrl string) GenreMetadata {
	collector := colly.NewCollector(scraper.Domains)
	var metadata GenreMetadata

	collector.OnHTML(".infobox-image img", func(e *colly.HTMLElement) {
		metadata.ImageUrl = "https:" + e.Attr("src")
	})

	collector.OnHTML(".infobox tr", func(e *colly.HTMLElement) {
		label := e.DOM.Find(".infobox-label").Text()
		header := e.DOM.Find(".infobox-header").Text()
		if label == "Stylistic origins" {
			metadata.StylisticOrigins = scrapeInfoBoxDataLinks(e)
		}
		if label == "Derivative forms" {
			metadata.DerivativeForms = scrapeInfoBoxDataLinks(e)
		}
		if header == "Subgenres" {
			metadata.Subgenres = scrapeInfoBoxFullDataLinks(e, e.DOM.Next())
		}
		if header == "Fusion genres" {
			metadata.FusionGenres = scrapeInfoBoxFullDataLinks(e, e.DOM.Next())
		}
	})

	collector.Visit(requestUrl)

	return metadata
}

func scrapeInfoBoxDataLinks(element *colly.HTMLElement) []Link {
	var links []Link
	element.DOM.Find(".infobox-data li").Each(func(i int, s *goquery.Selection) {
		url, hasUrl := s.Find("a").Attr("href")
		link, hasLink := getLink(element, s, url, hasUrl)
		if hasLink {
			links = append(links, link)
		}
	})

	// Handle cases where the links are not formatted in a list
	if len(links) == 0 {
		element.DOM.Find(".infobox-data a").Each(func(i int, s *goquery.Selection) {
			url, hasUrl := s.Attr("href")
			link, hasLink := getLink(element, s, url, hasUrl)
			if hasLink {
				links = append(links, link)
			}
		})
	}
	return links
}

func scrapeInfoBoxFullDataLinks(element *colly.HTMLElement, selection *goquery.Selection) []Link {
	var links []Link
	selection.Find(".infobox-full-data li").Each(func(i int, s *goquery.Selection) {
		url, hasUrl := s.Find("a").Attr("href")
		link, hasLink := getLink(element, s, url, hasUrl)
		if hasLink {
			links = append(links, link)
		}
	})

	return links
}

func getLink(element *colly.HTMLElement, s *goquery.Selection, url string, hasUrl bool) (Link, bool) {
	// Remove footnotes from link title
	footnoteRegex := regexp.MustCompile(`\[\d*\]`)
	title := footnoteRegex.ReplaceAllString(s.Text(), "")
	if title == "" {
		return Link{}, false
	}

	link := Link{
		Title: title,
	}

	if hasUrl {
		url = element.Request.AbsoluteURL(url)
		link.Url = &url
	}

	return link, true
}
