package smg

import (
	"encoding/xml"
	"time"
)

//SitemapLoc todo
type SitemapLoc struct {
	XMLName    xml.Name   `xml:"url"`
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
}

// SitemapIndexLoc todo
type SitemapIndexLoc struct {
	XMLName xml.Name   `xml:"url"`
	Loc     string     `xml:"loc"`
	LastMod *time.Time `xml:"lastmod,omitempty"`
}
