package cyoa

import (
	"html/template"
	"log"
)

// LoadTemplates loads and parses the template from the specified directory
func LoadTemplates(templateDir string) (*template.Template, error) {
	tpl, err := template.ParseGlob(templateDir + "/*.html")
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

// MustLoadTemplates is a helper that panics on error, useful for initialization.
func MustLoadTemplates(templateDir string) *template.Template {
	tpl, err := LoadTemplates(templateDir)
	if err != nil {
		log.Fatalf("Error loading templates from %s: %v", templateDir, err)
	}
	return tpl
}
