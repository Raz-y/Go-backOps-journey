package htmllink

import (
	"reflect"
	"strings"
	"testing"
)

// linkTestCase represents a single test case for testing the ParseLinks function.
type linkTestCase struct {
	name     string
	html     string
	expected []Link
}

// TestParsedLinks_SingleLink runs a series of tests to validate the behavior of the ParseLinks function.
// The tests cover scenarios including:
// 1. Parsing a single link
// 2. Parsing multiple links
// 3. Ignoring nested links (since they are considered invalid HTML)
// 4. Handling links without an href attribute
// 5. Handling empty documents
// 6. Handling documents containing only text nodes without links
func TestParseLinks_VariousScenarios(t *testing.T) {
	cases := []linkTestCase{
		{
			name: "Single link",
			html: `<a href="/home">Home</a>`,
			expected: []Link{
				{Href: "/home", Text: "Home"},
			},
		}, {
			name: "Multiple links",
			html: `<a href="/about">About</a> <a href="/contact">Contact</a>`,
			expected: []Link{
				{Href: "/about", Text: "About"},
				{Href: "/contact", Text: "Contact"},
			},
		},
		{
			name: "Nested links",
			html: `<a href="/outer-link">Outer <a href="/inner-link">Inner Link</a></a>`,
			expected: []Link{
				{Href: "/outer-link", Text: "Outer"},
				{Href: "/inner-link", Text: "Inner Link"},
			},
		},
		{
			name:     "Link without href",
			html:     `<a>Just text</a>`,
			expected: []Link{},
		},
		{
			name:     "Empty document",
			html:     ``,
			expected: []Link{},
		},
		{
			name:     "Document with only text nodes",
			html:     `<div>No links here</div>`,
			expected: []Link{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.html)
			links, err := ParseLinks(reader)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(links, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, links)
			}
		})
	}
}

// TestParseLinks_MalformedHTML verifies that the ParseLinks function handles malformed HTML gracefully.
// It checks that when the HTML is improperly closed or malformed, the parser still extracts the expected
// links without crashing. This test reflects the behavior of the html.Parse function, which is lenient
// with HTML errors.
func TestParseLinks_MalformedHTML(t *testing.T) {
	html := `<a href="/home">Home<a>More text without closing the link properly`
	reader := strings.NewReader(html)
	links, err := ParseLinks(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []Link{
		{Href: "/home", Text: "Home"},
	}

	if !reflect.DeepEqual(links, expected) {
		t.Errorf("Expected %v, got %v", expected, links)
	}
}
