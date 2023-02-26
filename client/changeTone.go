package client

import "github.com/jyz0309/notionAI-go/common"

const ChangeTone = "changeTone"

type ChangeToneContext struct {
	Text string               `json:"text"`
	Tone common.SupportedTone `json:"tone"`
	Type string               `json:"type"`
}

// ChangeTone uses specific tones to optimize content.
func (cli *NotionClient) ChangeTone(tone common.SupportedTone, content string) (string, error) {
	req := &NotionRequest{
		Context: ChangeToneContext{
			Text: content,
			Type: ChangeTone,
			Tone: tone,
		},
		Model: OpenAIModel,
	}
	return cli.Post(req)
}
