package azureclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Method string

const (
    GET Method = "GET"
    POST Method = "POST"
		PUT Method = "PUT"
    PATCH Method = "PATCH"
    DELETE Method = "DELETE"
)

func (c *client_) buildRequest(
	ctx context.Context,
	method Method,
	host string,
	resource string,
	queryParsms map[string][]string,
	reqbody any,
) (*http.Request, error) {
	body := []byte("{}")

	var err error
	if reqbody != nil {
		body, err = json.Marshal(reqbody)

		if err != nil {
			return nil, err
		}
	}

	resourceStringBuilder := strings.Builder{}

	resourceStringBuilder.WriteString(resource)

	if len(queryParsms) > 0 {
		resourceStringBuilder.WriteString("?")
	}

	for key, values := range queryParsms {
		for _, value := range values {
			resourceStringBuilder.WriteString(key)
			resourceStringBuilder.WriteString("=")
			resourceStringBuilder.WriteString(value)
			resourceStringBuilder.WriteString("&")
		}
	}

	resour := resourceStringBuilder.String()

	if len(resour) > 0 {
		resour = resour[:len(resour)-1]
	}

	date := time.Now().UTC().Format(http.TimeFormat)

	content, authHeader := createAuthHeader(
		string(method),
		host,
		resour,
		date,
		c.key,
		body,
	)

	fmt.Println(authHeader)
	fmt.Println(host)
	fmt.Println(resour)
	res, err := http.NewRequestWithContext(ctx, string(method), "https://"+host+resour, bytes.NewReader(body))
	res.Header.Add("x-ms-date", date)
	res.Header.Add("x-ms-content-sha256", content)
	res.Header.Add("Authorization", authHeader)
	res.Header.Add("Content-Type", "application/json")

	return res, nil
}

