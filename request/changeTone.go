package request

const ChangeTone = "changeTone"

type SupportedTone string

var (
	Professional    SupportedTone = "professional"
	Casual          SupportedTone = "casual"
	Straightforward SupportedTone = "straightforward"
	Confident       SupportedTone = "confident"
	Friendly        SupportedTone = "friendly"
)

type ChangeToneContext struct {
	Text string        `json:"text"`
	Tone SupportedTone `json:"tone"`
	Type string        `json:"type"`
}

// ChangeTone uses specific tones to optimize content.
func (cli *NotionClient) ChangeTone(content string, tone SupportedTone) (string, error) {
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
