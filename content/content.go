package content

import (
	"encoding/json"
	"html/template"
)

// ActivityStreamsContent represents the structure of ActivityStreams content.
type ActivityStreamsContent struct {
	Type       string            `json:"type"`
	Content    string            `json:"content"`
	ContentMap map[string]string `json:"contentMap"`
	Source     Source            `json:"source"`
}

// Source represents the source of the rendered HTML.
type Source struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

// RenderHTML renders the HTML content based on the observer's context.
func (asc *ActivityStreamsContent) RenderHTML(observerContext map[string]interface{}) (template.HTML, error) {
	if asc.Source.Type == "text/x-multicode" {
		// Dynamically generate HTML based on observer context
		// This is a placeholder implementation
		return template.HTML(asc.Source.Content), nil
	}
	return template.HTML(asc.Content), nil
}

// ParseActivityStreamsContent parses the JSON content into an ActivityStreamsContent struct.
func ParseActivityStreamsContent(jsonContent []byte) (*ActivityStreamsContent, error) {
	var asc ActivityStreamsContent
	err := json.Unmarshal(jsonContent, &asc)
	if err != nil {
		return nil, err
	}
	return &asc, nil
}
