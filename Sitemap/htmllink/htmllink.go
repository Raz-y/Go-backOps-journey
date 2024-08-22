// Package htmllink provides functions to parse and extract hyperlinks from HTML documents.
// It supports parsing HTML content from both `io.Reader` inputs and string inputs,
// making it versatile for various use cases, such as web scraping or HTML content analysis.

package htmllink

import (
	"errors"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a hyperlink (<a href="...">) found within an HTML document.
// The Href field contains the URL or path specified in the `href` attribute of the <a> tag,
// and the Text field contains the inner text enclosed by the <a> tag.

type Link struct {
	Href string // The href attribute of the <a> tag
	Text string // The inner text within the <a> tag
}

// ParseLinks parses an HTML document provided by an `io.Reader` and returns a slice of Link structs.
// Each Link struct corresponds to an <a> element found in the document.
// The function excludes nested links, meaning only the outermost <a> element is considered
// if there are nested <a> tags.
//
// It returns an error if the HTML document cannot be parsed.
//
// Example usage:
//
//    reader := strings.NewReader(`<a href="/home">Home</a>`)
//    links, err := htmllink.ParseLinks(reader)
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Println(links) // Output: [{Href:/home Text:Home}]

func ParseLinks(r io.Reader) ([]Link, error) {
	rootNode, err := html.Parse(r) // Parse the entire HTML document into a node tree
	if err != nil {
		return nil, errors.New("failed to parse HTML: " + err.Error())
	}
	if rootNode == nil {
		return nil, nil
	}
	ls := []Link{}                    // Initialize an empty slice
	ls = parseLinkNodes(rootNode, ls) // Capture the returned slice
	return ls, nil

}

// ParseLinks parses an HTML document provided by an `io.Reader` and returns a slice of Link structs.
// Each Link struct corresponds to an <a> element found in the document.
// The function excludes nested links, meaning only the outermost <a> element is considered
// if there are nested <a> tags.
//
// It returns an error if the HTML document cannot be parsed.
//
// Example usage:
//
//    reader := strings.NewReader(`<a href="/home">Home</a>`)
//    links, err := htmllink.ParseLinks(reader)
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Println(links) // Output: [{Href:/home Text:Home}]

func ParseLinksFromString(htmlStr string) ([]Link, error) {
	reader := strings.NewReader(htmlStr)
	return ParseLinks(reader)
}

// parseLinkNodes traverses the HTML node tree starting from the given node `n`,
// collecting <a> elements and converting them into Link structs.
//
// The function accumulates the Link structs into the provided slice `list` and returns it.

func parseLinkNodes(n *html.Node, list []Link) []Link {
	nodes := linkNodes(n) // Collect all <a> elements starting from the given node
	for _, node := range nodes {
		if link := buildLink(node); link != nil {
			list = append(list, *link)
		}

	}
	return list
}

// linkNodes traverses the HTML node tree recursively, starting from the node `n`,
// and collects all <a> elements into a slice.
//
// The function returns a slice of pointers to the collected <a> elements.

func linkNodes(n *html.Node) []*html.Node {
	var links []*html.Node

	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n} // Return immediately if the current node is an <a> element
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childLinks := linkNodes(c)
		links = append(links, childLinks...)
	}

	return links
}

// buildLink constructs a Link struct from an <a> element node `n`.
// It extracts the `href` attribute and the inner text of the <a> tag.

// If the `href` attribute is missing or empty, the function returns nil.

func buildLink(n *html.Node) *Link {
	var href string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			href = attr.Val // Extract the href attribute from the <a> element
			break
		}
	}
	if href == "" {
		return nil
	}
	return &Link{
		Href: href,
		Text: extractText(n),
	}
}

// extractText retrieves and concatenates the text content from an HTML node `n`,
// excluding any text within nested <a> elements. The function uses a strings.Builder
// to efficiently concatenate text nodes.
//
// The final result is a trimmed string containing the concatenated text.

func extractText(n *html.Node) string {
	var sb strings.Builder
	extractTextRecursive(n, &sb)
	return strings.TrimSpace(sb.String())
}

// extractTextRecursive is a helper function that traverses an HTML node tree
// starting from node `n`, and writes the text content to a provided `strings.Builder`.
//
// This function is called recursively for each child node, excluding text within nested <a> elements.

func extractTextRecursive(n *html.Node, sb *strings.Builder) {
	if n.Type == html.TextNode {
		sb.WriteString(n.Data) // Write text content to the strings.Builder
		sb.WriteString(" ")
	} else if n.Type == html.ElementNode {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractTextRecursive(c, sb)
		}
	}
}
