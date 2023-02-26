package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

const (
	UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3483.0 Safari/537.36"
	API         = "https://www.notion.so/api/v3/getCompletion"
	OpenAIModel = "openai-3"
)

type NotionContext interface{}

type NotionRequest struct {
	Context           NotionContext `json:"context"`
	ID                string        `json:"id"`
	IsSpacePermission bool          `json:"isSpacePermission"`
	Model             string        `json:"model"`
	SpaceId           string        `json:"spaceId"`
}

type NotionAIResp struct {
	Type       string `json:"type"`
	Completion string `json:"completion"`
}

type NotionClient struct {
	token   string
	spaceID string
}

func NewClient(token, spaceID string) *NotionClient {
	return &NotionClient{
		token:   token,
		spaceID: spaceID,
	}
}

func (cli *NotionClient) Post(req *NotionRequest) (string, error) {
	req.ID = uuid.NewString()
	req.SpaceId = cli.spaceID
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	r, err := http.NewRequest("POST", API, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	r.Header.Set("accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", UserAgent)
	r.Header.Set("cookie", fmt.Sprintf("token_v2=%v", cli.token))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	conpletions := strings.Split(string(body), "\n")
	resps := make([]*NotionAIResp, 0)
	for _, conpletion := range conpletions {
		if len(conpletion) > 0 {
			var aiResp *NotionAIResp
			err := json.Unmarshal([]byte(conpletion), &aiResp)
			if err != nil {
				return "", err
			}
			resps = append(resps, aiResp)
		}
	}
	var content string
	for _, resp := range resps {
		content += resp.Completion
	}
	return content, err
}
