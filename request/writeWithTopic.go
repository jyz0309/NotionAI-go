package request

import "github.com/jyz0309/notionAI-go/common"

type TopicContext struct {
	Topic string           `json:"topic"`
	Type  common.TopicType `json:"type"`
}

// WriteWithTopic writes for special topic, like Blog, Essay, Todo lis
func (cli *NotionClient) WriteWithTopic(topic common.TopicType, prompt string) (string, error) {
	var content string
	switch topic {
	case common.BrainstormIdeas:
		content += string(common.BrainstormIdeas_Prefix) + prompt
	case common.MeetingAgenda:
		content += string(common.MeetingAgenda_Prefix) + prompt
	case common.BlogPost:
		content += string(common.BlogPost_Prefix) + prompt
	case common.CreativeStory:
		content += string(common.CreativeStory_Prefix) + prompt
	case common.Essay:
		content += string(common.Essay_Prefix) + prompt
	case common.Poem:
		content += string(common.Poem_Prefix) + prompt
	case common.JobDescription:
		content += string(common.JobDescription_Prefix) + prompt
	case common.Outline:
		content += string(common.Outline_Prefix) + prompt
	case common.PressRelease:
		content += string(common.PressRelease_Prefix) + prompt
	case common.ProsConsList:
		content += string(common.ProsConsList_Prefix) + prompt
	case common.RecruitingEmail:
		content += string(common.RecruitingEmail_Prefix) + prompt
	case common.SalesEmail:
		content += string(common.SalesEmail_Prefix) + prompt
	case common.SocialMediaPost:
		content += string(common.SalesEmail_Prefix) + prompt
	case common.TodoList:
		content += string(common.TodoList_Prefix) + prompt
	default:
		content += prompt
	}
	req := &NotionRequest{
		Context: TopicContext{
			Topic: content,
			Type:  topic,
		},
		Model: OpenAIModel,
	}
	return cli.Post(req)
}
