package smg

// Options contains general attributes of Sitemap and SitemapIndex.
// OutputPath is the dir path to save the SitemapIndex file and it's
// sitemaps. Name of Sitemap output xml file which must be without ".xml" extension.
// Hostname of Sitemap urls which be prepended to all URLs. Compress option can be
// either enabled or disabled for Sitemap and SitemapIndex.
// ServerURI is used for making url of Sitemap in SitemapIndex.
type Options struct {
	Compress    bool   `xml:"-"`
	Name        string `xml:"-"`
	Hostname    string `xml:"-"`
	ServerURI   string `xml:"-"`
	OutputPath  string `xml:"-"`
	prettyPrint bool
}
