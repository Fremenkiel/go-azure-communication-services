package azureclient

import (
	"context"
)

type client_ struct {
	key    string
}

type Client interface {
	Request(
		ctx context.Context,
		method Method,
		host string,
		resource string,
		queryParsms map[string][]string,
		reqbody any,
	) ([]byte, error)
}

func New(
	key string,
) Client {
	return &client_{key}
}
