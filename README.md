sitemap-generator
=================

An awesome sitemap-generator Go module which is a comprehensive tool to create
and manage sitemap_index and sitemap files in a beautiful way. :)

Please see http://www.sitemaps.org/ for description of sitemap contents.

## Installation

# How to Use sitemap-generator

You can use the module in either Single-file sitemap or Multiple-files 
sitemaps with a sitemap_index file.

### Single sitemap usage

## TODO list
- [x] Develop: add new functionalities:
  - [x] Write the sitemap_index and sitemap files in xml format
  - [x] Compress option
  - [x] Break the sitemap xml file in case of exceeding 
    the sitemaps.org limits (50,000 urls OR 50MB uncompressed file)
  - [x] Ping search engines
  - [ ] Break the sitemap_index xml file in case of exceeding
    the sitemaps.org limits (50,000 urls OR 50MB uncompressed file)
- [ ] Support: Additional content types:
  - [ ] Video sitemaps
  - [ ] Image sitemaps
  - [ ] News sitemaps
  - [ ] Alternate Links
- [ ] Module Stability:
  - [x] Increase test coverage to more than %80.
  current coverage is: 80.6% of statements
  - [ ] Write more test files.


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
