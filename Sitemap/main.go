package main

import (
	"flag"
	"log"
	"sitemap/sitemap"
)

func main() {
	// Define flags for the domain name input and maximum depth crawl
	domainName := flag.String("site", "", "the domain to build a sitemap from. Example: 'google.com' or 'https://google.com'")
	maxDepth := flag.Int("depth", 5, "the maximum depth to build the sitemap for")
	flag.Parse()

	// Check if the domain name was provided
	if *domainName == "" {
		log.Fatal("No domain provided, use -h for help")
	}

	// Parse and validate the provided domain name
	u, err := sitemap.ParseAndValidateURL(*domainName)
	if err != nil {
		log.Fatalf("Error in parsing URL: %v", err)
	}

	// Build the sitemap by crawling the site up to the specified depth
	s := sitemap.BuildSiteMap(u, *maxDepth)

	// Generate the sitemap XML file from the collected URLs
	err = sitemap.BuildSiteMapXML(s)
	if err != nil {
		log.Fatalf("Failed to build sitemap XML: %v\n", err)
	}

	// Successfully completed
	log.Println("Sitemap generation completed successfully.")

}
