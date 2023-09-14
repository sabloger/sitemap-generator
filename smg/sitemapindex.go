package smg

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"
)

// SitemapIndex contains sitemap_index items which are SitemapURLs.
// New instances must be created with NewSitemapIndex() in order to set the
// Xmlns attribute correctly. Options is for general attributes
// Name is the filename which is used in Save method. Hostname is a prefix
// which wll be used for all URLs in SitemapIndex and it's Sitemaps.
// SitemapLocs is list of location structs of its Sitemaps.
// Sitemaps contains all Sitemaps which is belong to this SitemapIndex.
// ServerURI is used for making url of Sitemap in SitemapIndex.
type SitemapIndex struct {
	Options
	XMLName     xml.Name           `xml:"sitemapindex"`
	Xmlns       string             `xml:"xmlns,attr"`
	SitemapLocs []*SitemapIndexLoc `xml:"sitemap"`
	Sitemaps    []*Sitemap         `xml:"-"`
	ServerURI   string             `xml:"-"`
	finalURL    string
	mutex       sync.Mutex
	wg          sync.WaitGroup
}

var (
	searchEnginePingURLs = []string{
		"http://www.google.com/webmasters/tools/ping?sitemap=%s",
		"http://www.bing.com/webmaster/ping.aspx?siteMap=%s",
	}
)

// NewSitemapIndex builds returns new SitemapIndex.
// prettyPrint param makes the file easy to read and is
// recommended to be set to false for production use and
// is not changeable after initialization.
func NewSitemapIndex(prettyPrint bool) *SitemapIndex {
	s := &SitemapIndex{
		Xmlns:       "http://www.sitemaps.org/schemas/sitemap/0.9",
		SitemapLocs: make([]*SitemapIndexLoc, 0),
		Sitemaps:    make([]*Sitemap, 0),
		mutex:       sync.Mutex{},
		wg:          sync.WaitGroup{},
	}
	s.Name = "sitemap"
	s.Compress = true
	s.prettyPrint = prettyPrint
	return s
}

// Add adds an URL to a SitemapIndex.
func (s *SitemapIndex) Add(u *SitemapIndexLoc) {
	s.mutex.Lock()
	s.SitemapLocs = append(s.SitemapLocs, u)
	s.mutex.Unlock()
}

// SetSitemapIndexName sets the filename of SitemapIndex which be used to save the xml file.
// name param must not have .xml extension.
func (s *SitemapIndex) SetSitemapIndexName(name string) {
	s.Name = name
}

// NewSitemap builds a new instance of Sitemap and appends it in SitemapIndex's Sitemaps
// and sets it's Name nad Hostname
func (s *SitemapIndex) NewSitemap() *Sitemap {
	sm := NewSitemap(s.prettyPrint)
	s.Sitemaps = append(s.Sitemaps, sm)

	fileNum := len(s.Sitemaps)
	sm.SetName(fmt.Sprintf("sitemap%d", fileNum))
	sm.SetHostname(s.Hostname)
	sm.SetOutputPath(s.OutputPath)
	sm.SetCompress(s.Compress)
	return sm
}

// AppendSitemap appends a Sitemap instance into it's Sitemaps.
// Does not change and configurations.
func (s *SitemapIndex) AppendSitemap(sm *Sitemap) {
	s.Sitemaps = append(s.Sitemaps, sm)
}

// SetHostname sets the Hostname for SitemapIndex and it's Sitemaps
// and sets it as Hostname of new Sitemap entries built using NewSitemap method.
func (s *SitemapIndex) SetHostname(hostname string) {
	s.Hostname = hostname
	for _, sitemap := range s.Sitemaps {
		sitemap.SetHostname(s.Hostname)
	}
}

// SetOutputPath sets the OutputPath for SitemapIndex and it's Sitemaps
// and sets it as OutputPath of new Sitemap entries built using NewSitemap method.
// this path can be a multi-level dir path and will be used in Save method.
func (s *SitemapIndex) SetOutputPath(outputPath string) {
	s.OutputPath = outputPath
	for _, sitemap := range s.Sitemaps {
		sitemap.SetOutputPath(s.OutputPath)
	}
}

// SetServerURI sets the ServerURI for SitemapIndex and it's Sitemaps
// and sets it as OutputPath of new Sitemap entries built using NewSitemap method.
func (s *SitemapIndex) SetServerURI(serverURI string) {
	s.ServerURI = serverURI
}

// SetCompress sets the Compress option to be either enabled or disabled for SitemapIndex
// and it's Sitemaps and sets it as Compress of new Sitemap entries built using NewSitemap method.
// When Compress is enabled, the output file is compressed using gzip with .xml.gz extension.
func (s *SitemapIndex) SetCompress(compress bool) {
	s.Compress = compress
	for _, sitemap := range s.Sitemaps {
		sitemap.SetCompress(s.Compress)
	}
}

// WriteTo writes XML encoded sitemap to given io.Writer.
// Implements io.WriterTo interface.
func (s *SitemapIndex) WriteTo(writer io.Writer) (int64, error) {
	headerCount, err := writer.Write([]byte(xml.Header))
	if err != nil {
		return 0, err
	}
	encoder := xml.NewEncoder(writer)
	if s.prettyPrint {
		encoder.Indent("", "  ")
	}
	err = encoder.Encode(s)
	if err != nil {
		return 0, err
	}

	bodyCount, err := writer.Write([]byte{'\n'})
	if err != nil {
		return 0, err
	}
	return int64(headerCount + bodyCount), err
}

// Save makes the OutputPath in case of absence and saves the SitemapIndex
// and it's Sitemaps into OutputPath as separate files using their Name.
func (s *SitemapIndex) Save() (string, error) {
	err := checkAndMakeDir(s.OutputPath)
	if err != nil {
		return "", err
	}

	err = s.saveSitemaps()
	if err != nil {
		return "", err
	}

	var filename string
	if s.Compress {
		filename = s.Name + fileGzExt
	} else {
		filename = s.Name + fileExt
	}

	buf := bytes.Buffer{}
	_, err = s.WriteTo(&buf)
	if err != nil {
		return "", err
	}
	_, err = writeToFile(filename, s.OutputPath, s.Compress, buf.Bytes())
	// s.finalURL = filepath.Join(s.Hostname, s.OutputPath, filename)

	output, err := url.Parse(s.Hostname)
	if err != nil {
		return "", err
	}
	output.Path = path.Join(output.Path, s.OutputPath, filename)
	s.finalURL = output.String()

	return filename, err
}

func (s *SitemapIndex) saveSitemaps() error {
	for _, sitemap := range s.Sitemaps {
		s.wg.Add(1)
		go func(sm *Sitemap) {
			defer s.wg.Done()

			smFilenames, err := sm.Save()
			if err != nil {
				log.Println("Error while saving this sitemap:", sm.Name, err)
				return
			}
			for _, smFilename := range smFilenames {
				// sm.SitemapIndexLoc.Loc = filepath.Join(s.Hostname, s.ServerURI, smFilename)

				output, err := url.Parse(s.Hostname)
				if err != nil {
					log.Println("Error while saving this sitemap:", sm.Name, err)
					return
				}
				output.Path = path.Join(output.Path, s.ServerURI, smFilename)
				smIndexLoc := &SitemapIndexLoc{
					Loc: output.String(),
				}
				s.Add(smIndexLoc)
			}
		}(sitemap)
	}
	s.wg.Wait()
	return nil
}

// PingSearchEngines pings search engines
func (s *SitemapIndex) PingSearchEngines(pingURLs ...string) error {
	if s.finalURL == "" {
		return errors.New("the save method must be called before ping")
	}
	pingURLs = append(pingURLs, searchEnginePingURLs...)

	wg := sync.WaitGroup{}
	client := http.Client{Timeout: 5 * time.Second}
	for _, pingURL := range pingURLs {
		wg.Add(1)
		go func(urlFormat string) {
			defer wg.Done()
			
			urlStr := fmt.Sprintf(urlFormat, s.finalURL)
			log.Println("Pinging", urlStr)

			resp, err := client.Get(urlStr)
			if err != nil {
				log.Println("Failed to Ping:", urlStr)
				return
			}
			resp.Body.Close()
			log.Println("Successful Ping:", urlStr)
		}(pingURL)
	}
	wg.Wait()
	return nil
}
