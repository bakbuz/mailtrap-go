package emails

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bakbuz/mailtrap-go/core/constants"
	"github.com/bakbuz/mailtrap-go/emails/request"
	"github.com/bakbuz/mailtrap-go/emails/response"
)

type EmailClient interface {
	Send(ctx context.Context, req *request.SendEmailRequest) (*response.SendEmailResponse, error)
}

type emailClient struct {
	apiKey string
}

func NewEmailClient(apiKey string) EmailClient {
	return &emailClient{apiKey: apiKey}
}

func (c *emailClient) Send(ctx context.Context, req *request.SendEmailRequest) (*response.SendEmailResponse, error) {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewReader(jsonBytes)

	url := constants.Endpoints_SendDefaultUrl + "/api/send"

	client := &http.Client{}
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, payload)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", string(constants.MimeType_ApplicationJSON))
	httpReq.Header.Add("Authorization", "Bearer "+c.apiKey)

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	var res *response.SendEmailResponse
	if err := json.Unmarshal(body, &res); err != nil {
		panic(err)
	}
	return res, nil
}
