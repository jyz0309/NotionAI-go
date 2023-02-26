package request

const ContinueWriting = "continueWriting"

type ContinueWritingContext struct {
	PreviousContent string `json:"previousContent"`
	Type            string `json:"type"`
}

// ContinueWriting can generate the following content based on the given input content.
func (cli *NotionClient) ContinueWriting(content string) (string, error) {
	req := &NotionRequest{
		Context: ContinueWritingContext{
			PreviousContent: content,
			Type:            ContinueWriting,
		},
		Model: OpenAIModel,
	}
	return cli.Post(req)
}
