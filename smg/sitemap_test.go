package smg

import (
	"testing"
	"time"
)

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
