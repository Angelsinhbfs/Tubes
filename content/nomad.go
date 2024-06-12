package content

import (
	"encoding/json"
	"html/template"
)

// NomadContent represents the structure of Nomad content.
type NomadContent struct {
	Body string `json:"body"`
	HTML string `json:"html"`
	Type string `json:"type"`
}

// RenderHTML renders the HTML content based on the observer's context.
func (nc *NomadContent) RenderHTML(observerContext map[string]interface{}) (template.HTML, error) {
	if nc.Type == "text/x-multicode" {
		// Dynamically generate HTML based on observer context
		// This is a placeholder implementation
		return template.HTML(nc.Body), nil
	}
	return template.HTML(nc.HTML), nil
}

// ParseNomadContent parses the JSON content into a NomadContent struct.
func ParseNomadContent(jsonContent []byte) (*NomadContent, error) {
	var nc NomadContent
	err := json.Unmarshal(jsonContent, &nc)
	if err != nil {
		return nil, err
	}
	return &nc, nil
}
