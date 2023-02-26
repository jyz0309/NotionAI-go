package request

import "github.com/jyz0309/notionAI-go/common"

type HelpDraftContext struct {
	Prompt string            `json:"prompt"`
	Type   common.PromptType `json:"type"`
}

// HelpMeDraft can generate a draft based on the given prompt.
func (cli *NotionClient) HelpMeDraft(prompt string) (string, error) {
	req := &NotionRequest{
		Context: HelpDraftContext{
			Prompt: prompt,
			Type:   common.HelpMeDraft,
		},
		Model: OpenAIModel,
	}
	return cli.Post(req)
}

type HelpEditContext struct {
	Prompt       string            `json:"prompt"`
	Type         common.PromptType `json:"type"`
	SelectedText string            `json:"selectedText"`
}

// HelpMeEdit can optimize content based on the given prompt.
func (cli *NotionClient) HelpMeEdit(content, prompt string) (string, error) {
	req := &NotionRequest{
		Context: HelpEditContext{
			Prompt:       prompt,
			Type:         common.HelpMeDraft,
			SelectedText: content,
		},
		Model: OpenAIModel,
	}
	return cli.Post(req)
}
