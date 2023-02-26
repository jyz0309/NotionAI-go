package client

import "github.com/jyz0309/notionAI-go/common"

type PromptContext struct {
	SelectedText string            `json:"selectedText"`
	Type         common.PromptType `json:"type"`
}

// WriteWithPrompt writes with special prompt, like summarize, explain_this, improve_writing
func (cli *NotionClient) WriteWithPrompt(content string, promptType common.PromptType) (string, error) {
	req := &NotionRequest{
		Context: PromptContext{
			SelectedText: content,
			Type:         promptType,
		},
		Model: OpenAIModel,
	}
	return cli.Post(req)
}
