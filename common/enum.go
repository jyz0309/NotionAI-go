package common

type SupportedTone string

var (
	Professional    SupportedTone = "professional"
	Casual          SupportedTone = "casual"
	Straightforward SupportedTone = "straightforward"
	Confident       SupportedTone = "confident"
	Friendly        SupportedTone = "friendly"
)

type PromptType string

var (
	HelpMeDraft        PromptType = "helpMeDraft"
	HelpMeEdit         PromptType = "helpMeEdit"
	MakeLonger         PromptType = "makeLonger"
	ImpoveWriting      PromptType = "impoveWriting"
	FixSpellingGrammar PromptType = "fixSpellingGrammar"
	MakeShorter        PromptType = "makeShorter"
	SimplifyLanguage   PromptType = "simplifyLanguage"
	ExplainThis        PromptType = "explainThis"
	Summarize          PromptType = "summarize"
)

type SupportedLanguage string

var (
	English    SupportedLanguage = "english"
	Chinese    SupportedLanguage = "chinese"
	Korean     SupportedLanguage = "korean"
	Japanese   SupportedLanguage = "japanese"
	Spanish    SupportedLanguage = "spanish"
	Russian    SupportedLanguage = "russian"
	French     SupportedLanguage = "french"
	Portuguese SupportedLanguage = "portuguese"
	German     SupportedLanguage = "german"
	Italian    SupportedLanguage = "italian"
	Dutch      SupportedLanguage = "dutch"
	Indonesian SupportedLanguage = "indonesian"
	Filipino   SupportedLanguage = "filipino"
	Vietnamese SupportedLanguage = "vietnamese"
)

type TopicType string

var (
	BrainstormIdeas TopicType = "brainstormIdeas"
	BlogPost        TopicType = "blogPost"
	Outline         TopicType = "outline"
	SocialMediaPost TopicType = "socialMediaPost"
	PressRelease    TopicType = "pressRelease"
	CreativeStory   TopicType = "creativeStory"
	Essay           TopicType = "essay"
	Poem            TopicType = "poem"
	TodoList        TopicType = "todoList"
	MeetingAgenda   TopicType = "meetingAgenda"
	ProsConsList    TopicType = "prosConsList"
	JobDescription  TopicType = "jobDescription"
	SalesEmail      TopicType = "salesEmail"
	RecruitingEmail TopicType = "recruitingEmail"
)

type TopicPrefix string

var (
	BrainstormIdeas_Prefix TopicPrefix = "Brainstorm ideas on "
	BlogPost_Prefix        TopicPrefix = "Write a blog post about "
	Outline_Prefix         TopicPrefix = "Write an outline about "
	SocialMediaPost_Prefix TopicPrefix = "Write a social media post about "
	PressRelease_Prefix    TopicPrefix = "Write a press release about "
	CreativeStory_Prefix   TopicPrefix = "Write a creative story about "
	Essay_Prefix           TopicPrefix = "Write an essay about "
	Poem_Prefix            TopicPrefix = "Write a poem about"
	TodoList_Prefix        TopicPrefix = "Write a todo list about "
	MeetingAgenda_Prefix   TopicPrefix = "Write a meeting agenda about "
	ProsConsList_Prefix    TopicPrefix = "Write a pros and cons list about "
	JobDescription_Prefix  TopicPrefix = "Write a job description about "
	SalesEmail_Prefix      TopicPrefix = "Write a sales email about "
	RecruitingEmail_Prefix TopicPrefix = "Write a recruiting email about "
)
