package smg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Urls []UrlData `xml:"url"`
}

type UrlData struct {
	XMLName xml.Name `xml:"url"`
	Loc string `xml:"loc"`
	LasMod string `xml:"lastmod"`
	ChangeFreq string `xml:"changefreq"`
	Priority string `xml:"priority"`
}


// TestSingleSitemap tests the module against Single-file sitemap usage format.
func TestSingleSitemap(t *testing.T) {
	path := getNewPath()
	now := time.Now().UTC()
	routes := buildRoutes(10, 40, 10)

	sm := NewSitemap(true)
	sm.SetName("single_sitemap")
	sm.SetHostname(baseURL)
	sm.SetOutputPath(path)
	sm.SetLastMod(&now)
	sm.SetCompress(false)

	for _, route := range routes {
		err := sm.Add(&SitemapLoc{
			Loc:        route,
			LastMod:    &now,
			ChangeFreq: Always,
			Priority:   0.4,
		})
		if err != nil {
			t.Fatal("Unable to add SitemapLoc:", err)
		}
	}
	// -----------------------------------------------------------------

	// Compressed files;
	filenames, err := sm.Save()
	if err != nil {
		t.Fatal("Unable to Save Compressed Sitemap:", err)
	}
	for _, filename := range filenames {
		assertOutputFile(t, path, filename)
	}

	// Plain files:
	sm.SetCompress(false)
	filenames, err = sm.Save()
	if err != nil {
		t.Fatal("Unable to Save Sitemap:", err)
	}
	for _, filename := range filenames {
		assertOutputFile(t, path, filename)
	}
	// -----------------------------------------------------------------
	// Removing the generated path and files
	removeTmpFiles(t, path)
}

// TestSitemapAdd tests that the Add function produces a proper URL
func TestSitemapAdd(t *testing.T) {
	path := "./tmp/sitemap_test"
	testLocation := "/test"
	now := time.Now().UTC()
	sm := NewSitemap(true)
	sm.SetName("single_sitemap")
	sm.SetHostname(baseURL)
	sm.SetOutputPath(path)
	sm.SetLastMod(&now)
	sm.SetCompress(false)

	err := sm.Add(&SitemapLoc{
		Loc:        testLocation,
		LastMod:    &now,
		ChangeFreq: Always,
		Priority:   0.4,
	})
	if err != nil {
		t.Fatal("Unable to add SitemapLoc:", err)
	}
	expectedUrl := fmt.Sprintf("%s%s", baseURL, testLocation)
	filepath, err := sm.Save()
	if err != nil {
		t.Fatal("Unable to Save Sitemap:", err)
	}
	xmlFile, err := os.Open(fmt.Sprintf("%s/%s",path, filepath[0]))
	if err != nil {
		t.Fatal("Unable to open file:", err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var urlSet UrlSet
	err = xml.Unmarshal(byteValue, &urlSet)
	if err != nil {
		t.Fatal("Unable to unmarhsall sitemap byte array into xml: ", err)
	}
	actualUrl := urlSet.Urls[0].Loc
	if actualUrl != expectedUrl {
		t.Fatal(fmt.Sprintf("URL Mismatch: \nActual: %s\nExpected: %s", actualUrl, expectedUrl))
	}

	removeTmpFiles(t, "./tmp")

}
