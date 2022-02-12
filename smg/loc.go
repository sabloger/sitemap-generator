package smg

import (
	"encoding/xml"
	"time"
)

// SitemapLoc contains data related to <url> tag in Sitemap.
type SitemapLoc struct {
	XMLName    xml.Name   `xml:"url"`
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
}

// SitemapIndexLoc contains data related to <sitemap> tag in SitemapIndex.
type SitemapIndexLoc struct {
	XMLName xml.Name   `xml:"url"`
	Loc     string     `xml:"loc"`
	LastMod *time.Time `xml:"lastmod,omitempty"`
}
