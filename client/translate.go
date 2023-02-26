package client

import "github.com/jyz0309/notionAI-go/common"

const Translate = "Translate"

type TranslateContext struct {
	Text     string                   `json:"text"`
	Language common.SupportedLanguage `json:"language"`
	Type     string                   `json:"type"`
}

// Translate will translate the content to given language
func (cli *NotionClient) Translate(lang common.SupportedLanguage, content string) (string, error) {
	req := &NotionRequest{
		Context: TranslateContext{
			Text:     content,
			Type:     ChangeTone,
			Language: lang,
		},
		Model: OpenAIModel,
	}
	return cli.Post(req)
}
