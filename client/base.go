package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

const (
	UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3483.0 Safari/537.36"
	API         = "https://www.notion.so/api/v3/getCompletion"
	OpenAIModel = "openai-4"
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
	client  *http.Client
}

func NewClient(token, spaceID string) *NotionClient {
	return &NotionClient{
		token:   token,
		spaceID: spaceID,
		client:  &http.Client{},
	}
}

func NewProxyClient(token, spaceID string, proxy string) *NotionClient {
	proxyAddress, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	return &NotionClient{
		token:   token,
		spaceID: spaceID,
		client:  &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyAddress)}},
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
	resp, err := cli.client.Do(r)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		if err != nil {
			return "", err
		}
		scanner := bufio.NewScanner(resp.Body)
		conpletions := make([]string, 0)
		for scanner.Scan() {
			conpletions = append(conpletions, scanner.Text())
		}
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
	} else {
		return "", fmt.Errorf("failed to get response, statusCode[%v], body[%v]", resp.StatusCode, resp.Body)
	}

}
