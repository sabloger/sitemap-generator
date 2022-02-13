sitemap-generator
=================

A high-performance sitemap-generator Go module which is a comprehensive tool to create
and manage sitemap_index and sitemap files in a beautiful way. :)

Please see http://www.sitemaps.org/ for description of sitemap contents.

## Installation
Use `go get`:

`go get github.com/sabloger/sitemap-generator`

# How to Use sitemap-generator

You can use the module in either Single-file sitemap or Multiple-files 
sitemaps with a sitemap_index file.

### Single sitemap usage
```go
package main

import (
  "fmt"
  "github.com/sabloger/sitemap-generator/smg"
  "log"
  "time"
)

func main() {
  now := time.Now().UTC()

  sm := smg.NewSitemap(true) // The argument is PrettyPrint which must be set on initializing
  sm.SetName("single_sitemap") // Optional
  sm.SetHostname("https://www.example.com")
  sm.SetOutputPath("./some/path")
  sm.SetLastMod(&now)
  sm.SetCompress(false) // Default is true

  // Adding URL items
  err := sm.Add(&smg.SitemapLoc{
    Loc:        "some/uri.html",
    LastMod:    &now,
    ChangeFreq: smg.Always,
    Priority:   0.4,
  })
  if err != nil {
    log.Fatal("Unable to add SitemapLoc:", err)
  }

  // Save func saves the xml files and returns more than one filename in case of split large files.
  filenames, err := sm.Save()
  if err != nil {
    log.Fatal("Unable to Save Sitemap:", err)
  }
  for i, filename := range filenames {
    fmt.Println("file no.", i+1, filename)
  }
}
```
`single_sitemap.xml` will look like:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<url>
  <loc>https:/www.example.com/some/uri.html</loc>
  <lastmod>2022-02-12T16:29:46.45013Z</lastmod>
  <changefreq>always</changefreq>
  <priority>0.4</priority>
</url>
</urlset>
```


### SitemapIndex usage
```go
package main

import (
  "fmt"
  "github.com/sabloger/sitemap-generator/smg"
  "log"
  "time"
)

func main() {
  now := time.Now().UTC()

  smi := smg.NewSitemapIndex(true)
  smi.SetCompress(false)
  smi.SetSitemapIndexName("an_optional_name_for_sitemap_index")
  smi.SetHostname("https://www.example.com")
  smi.SetOutputPath("./sitemap_index_example/")
  smi.SetServerURI("/sitemaps/") // Optional

  smBlog := smi.NewSitemap()
  smBlog.SetName("blog_sitemap")
  smBlog.SetLastMod(&now)
  err := smBlog.Add(&smg.SitemapLoc{
    Loc:        "blog/post/1231",
    LastMod:    &now,
    ChangeFreq: smg.Weekly,
    Priority:   0.8,
  })
  if err != nil {
    log.Fatal("Unable to add SitemapLoc:", err)
  }

  smNews := smi.NewSitemap()
  smNews.SetLastMod(&now)
  err = smNews.Add(&smg.SitemapLoc{
    Loc:        "news/2021-01-05/a-news-page",
    LastMod:    &now,
    ChangeFreq: smg.Weekly,
    Priority:   1,
  })
  if err != nil {
    log.Fatal("Unable to add SitemapLoc:", err)
  }

  filename, err := smi.Save()
  if err != nil {
    log.Fatal("Unable to Save Sitemap:", err)
  }

  fmt.Println("sitemap_index file:", filename)
}
```
the output directory will be like this:
```
sitemap_index_example
├── an_optional_name_for_sitemap_index.xml
├── blog_sitemap.xml
└── sitemap2.xml
```
`an_optional_name_for_sitemap_index.xml` will look like:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https:/www.example.com/sitemaps/blog_sitemap.xml</loc>
    <lastmod>2022-02-12T18:38:06.671183Z</lastmod>
  </url>
  <url>
    <loc>https:/www.example.com/sitemaps/sitemap2.xml</loc>
    <lastmod>2022-02-12T18:38:06.671183Z</lastmod>
  </url>
</sitemapindex>
```

## TODO list
- [x] Develop: add new functionalities:
  - [x] Write the sitemap_index and sitemap files in xml format
  - [x] Compress option
  - [x] Break the sitemap xml file in case of exceeding 
    the sitemaps.org limits (50,000 urls OR 50MB uncompressed file)
  - [x] Ability to set Sitemap uri on server to set on it's url
    in sitemap_index file
  - [x] Ping search engines for sitemap_index
  - [ ] Ping search engines for single sitemap
  - [ ] Break the sitemap_index xml file in case of exceeding
    the sitemaps.org limits (50,000 urls OR 50MB uncompressed file)
- [ ] Support: Additional content types:
  - [ ] Video sitemaps
  - [ ] Image sitemaps
  - [ ] News sitemaps
  - [ ] Alternate Links
- [ ] Module Stability:
  - [x] Increase test coverage to more than %80.
  current coverage is: 86.6% of statements
  - [X] Write tests for different usages.


## LINKS
GoDoc documentation:
https://godoc.org/github.com/sabloger/sitemap-generator

Git repository:
https://github.com/sabloger/sitemap-generator


## License
MIT


## THANKS
Special thanks to authors of these repos whom I inspired from their sitemap modules to create this awesome module. :)
https://github.com/snabb/sitemap
https://github.com/ikeikeikeike/go-sitemap-generator
