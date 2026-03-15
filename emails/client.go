package emails

import (
	"github.com/Fremenkiel/go-azure-communication-services/azureclient"
	"context"
	"encoding/json"
)

type Client interface {
	SendEmail(
		ctx context.Context,
		payload Payload,
	) (EmailResult, error)
}

type _client struct {
	cl      azureclient.Client
	host    string
	version string
}

func NewClient(
	host string,
	key string,
	version *string,
) Client {
	cl := azureclient.New(key)
	
	v := defaultAPIVersion
	if version != nil {
		v = *version
	}
	return &_client{
		cl:      cl,
		host:    host,
		version: v,
	}
}

func NewClientWithLogger(
	host string,
	key string,
	version *string,
) Client {
	cl := azureclient.New(key)	

	v := defaultAPIVersion
	if version != nil {
		v = *version
	}
	return &_client{
		cl:      cl,
		host:    host,
		version: v,
	}
}

func (c *_client) SendEmail(
	ctx context.Context,
	payload Payload,
) (result EmailResult, err error) {
	res, err := c.cl.Request(
		ctx,
		azureclient.POST,
		c.host,
		"/emails:send",
		map[string][]string{
			"api-version": {c.version},
		},
		payload,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return
}
