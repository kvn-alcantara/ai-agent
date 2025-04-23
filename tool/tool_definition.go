package tool

import (
	"encoding/json"

	"github.com/anthropics/anthropic-sdk-go"
)

// ToolDefinition represents a tool with a name, description, input schema, and a function to process input.
type ToolDefinition struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	InputSchema anthropic.ToolInputSchemaParam `json:"input_schema"`
	Function    func(input json.RawMessage) (string, error)
}
