package smg

import (
	"time"
)

//SitemapLoc todo
type SitemapLoc struct {
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
}

// SitemapIndexLoc todo
type SitemapIndexLoc struct {
	Loc     string     `xml:"loc"`
	LastMod *time.Time `xml:"lastmod,omitempty"`
}

//func (u *SitemapURL) toXMLBytes() []byte {
//	buffer := bytes.Buffer{}
//
//}
