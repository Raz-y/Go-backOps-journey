// Package sitemap provides tools to generate XML sitemaps by crawling websites and collecting URLs.
// It includes functions for validating URLs, performing website crawls, and generating sitemaps
// in compliance with the sitemap protocol.
package sitemap

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sitemap/htmllink"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

// url represents a single URL entry in the sitemap.
type URL struct {
	Loc string `xml:"loc"`
}

// urlset represents the entire sitemap XML structure.

type URLSet struct {
	Xmlns string `xml:"xmlns,attr"`
	URLs  []URL  `xml:"url"`
}

var baseURL *url.URL // Base URL used to resolve relative links

// ParseAndValidateURL parses and validates a domain name, ensuring it has a valid scheme, host, and TLD.
// If no scheme is provided, "https://" is assumed. Returns the parsed *url.URL or an error if validation fails.
// Example: u, err := ParseAndValidateURL("example.com")
func ParseAndValidateURL(domainName string) (*url.URL, error) {

	// If no scheme is provided, assume it's an HTTPS URL.
	if !strings.HasPrefix(domainName, "http://") && !strings.HasPrefix(domainName, "https://") {
		domainName = "https://" + domainName
	}

	// Parse the URL
	u, err := url.Parse(domainName)
	if err != nil || u.Scheme == "" || u.Host == "" || !strings.Contains(u.Host, ".") {
		return nil, fmt.Errorf("invalid domain: %s", domainName)
	}

	// Normalize the URL (lowercase scheme and host)
	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)

	return u, nil
}

// BuildSiteMap crawls a website starting from the given URL, collecting unique URLs up to the specified depth.
// Returns a slice of URLs found or an error if the starting URL is invalid.
func BuildSiteMap(startURL *url.URL, maxDepth int) []string {
	baseURL = startURL // Set the base URL for resolving relative links

	seen := make(map[string]struct{}) // Track seen URLs to avoid duplicates
	var q map[string]struct{}
	nq := map[string]struct{}{
		startURL.String(): {}, // Initialize the queue with the starting URL
	}

	for d := 0; d <= maxDepth; d++ {
		if len(nq) == 0 {
			break // Stop if there are no more URLs to process
		}

		q, nq = nq, make(map[string]struct{}) // Swap current and next queue

		for currURL := range q {
			if _, ok := seen[currURL]; ok {
				continue // Skip URLs that have already been seen
			}
			seen[currURL] = struct{}{} // Mark the URL as seen

			links := fetchLinks(currURL) // Fetch links from the current URL

			for _, link := range links {
				normalizedURL := normalizeURL(link) // Normalize the URL

				if isValidLink(normalizedURL) && normalizedURL != nil {

					nq[normalizedURL.String()] = struct{}{} // Add valid, normalized URLs to the next queue
				}
			}
		}
	}

	ret := make([]string, 0, len(seen)) // Prepare the slice to return collected URLs
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

// BuildSiteMapXML generates a sitemap in XML format from a list of URLs and writes it to "sitemap.xml".
// Returns an error if XML generation, file creation, or writing fails.
func BuildSiteMapXML(urls []string) error {
	toXml := URLSet{
		Xmlns: xmlns,
	}

	// Populate the URLSet with URLs
	for _, u := range urls {
		toXml.URLs = append(toXml.URLs, URL{
			Loc: u,
		})
	}

	// Generate XML output with indentation
	output, err := xml.MarshalIndent(toXml, "", "  ")
	if err != nil {
		return fmt.Errorf("error generating XML: %v", err)
	}

	// Add XML header
	xmlHeader := []byte(xml.Header)
	output = append(xmlHeader, output...)

	// Create the sitemap.xml file
	file, err := os.Create("sitemap.xml")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	defer file.Close()

	// Write the XML output to the file
	_, err = file.Write(output)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)

	}
	return nil

}

// fetchLinks fetches a URL and returns all the valid links found on the page.
// It returns a slice of strings representing the absolute URLs found.
func fetchLinks(rawURL string) []string {
	resp, err := http.Get(rawURL)
	if err != nil {
		return []string{} // Return an empty slice on error
	}
	defer resp.Body.Close()

	// Parse and filter the links
	return filter(hrefs(resp.Body, baseURL.String()), withPrefix(baseURL.String()))
}

// hrefs parses the links from the response body and returns them as a slice of strings.
// It considers both absolute and relative URLs, converting relative URLs to absolute using the base URL.
func hrefs(r io.Reader, base string) []string {
	links, _ := htmllink.ParseLinks(r)
	var ret []string
	for _, l := range links {
		href := strings.TrimSpace(l.Href) // Clean up the link

		// Convert relative links to absolute using the base URL
		if strings.HasPrefix(href, "/") {
			ret = append(ret, base+href)
		} else if strings.HasPrefix(href, "http") {
			ret = append(ret, href)
		}
	}
	return ret
}

// normalizeURL removes query parameters, fragments, and trailing slashes to avoid duplicates.
// It returns a pointer to a normalized URL object.
func normalizeURL(href string) *url.URL {
	u, err := url.Parse(href)
	if err != nil {
		return nil
	}
	normalized := *u
	// Remove fragment, query parameters, and trailing slashes
	normalized.Fragment = ""
	normalized.RawQuery = ""
	normalized.Path = strings.TrimSuffix(normalized.Path, "/")

	return &normalized
}

// isValidLink checks if a URL is valid for inclusion in the sitemap.
// It excludes mailto links, comment pages, and URLs with different hosts.
func isValidLink(u *url.URL) bool {
	if strings.Contains(u.Path, "comment-page-") {
		return false // Exclude comment page URLs (wordpress)
	}
	// Exclude mailto links and URLs with different hosts

	return u.Scheme != "mailto" && u.Host == baseURL.Host && u.Fragment == ""
}

// filter applies a filtering function to a slice of strings and returns the filtered slice.
func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link) // Keep links that match the filter condition
		}
	}
	return ret
}

// withPrefix returns a filtering function that checks if a string starts with a given prefix.
func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx) // Check if the link starts with the specified prefix
	}
}
