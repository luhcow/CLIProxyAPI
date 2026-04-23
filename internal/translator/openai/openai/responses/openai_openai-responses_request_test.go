package responses

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestConvertOpenAIResponsesRequestToOpenAIChatCompletions_PreservesCustomAndBuiltinTools(t *testing.T) {
	inputJSON := []byte(`{
		"model": "gpt-5.4",
		"input": "apply a patch and search the web",
		"tools": [
			{
				"type": "custom",
				"name": "ApplyPatch",
				"description": "Apply a patch to files",
				"format": {"type": "text"}
			},
			{
				"type": "web_search"
			}
		]
	}`)

	output := ConvertOpenAIResponsesRequestToOpenAIChatCompletions("gpt-5.4", inputJSON, true)

	if got := gjson.GetBytes(output, "tools.0.type").String(); got != "custom" {
		t.Fatalf("tools.0.type = %q, want %q: %s", got, "custom", string(output))
	}
	if got := gjson.GetBytes(output, "tools.0.name").String(); got != "ApplyPatch" {
		t.Fatalf("tools.0.name = %q, want %q: %s", got, "ApplyPatch", string(output))
	}
	if got := gjson.GetBytes(output, "tools.0.format.type").String(); got != "text" {
		t.Fatalf("tools.0.format.type = %q, want %q: %s", got, "text", string(output))
	}
	if got := gjson.GetBytes(output, "tools.1.type").String(); got != "web_search" {
		t.Fatalf("tools.1.type = %q, want %q: %s", got, "web_search", string(output))
	}
}
