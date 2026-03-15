package azureclient

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
)

func (c *client_) Request(
	ctx context.Context,
	method Method,
	host string,
	resource string,
	queryParsms map[string][]string,
	reqbody any,
) ([]byte, error) {
	req, err := c.buildRequest(ctx, POST, host, resource, queryParsms, reqbody)
	if err != nil {
		log.Println(method + " Request failed")
		log.Println("error", err.Error())
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Println(method + " Request failed")
		log.Println("response: " + string(resBody))
		log.Println("status_code: " + res.Status)

		return nil, errors.New(string(resBody))
	}

	if len(resBody) == 0 {
		return nil, nil
	}

	return resBody, nil
}
