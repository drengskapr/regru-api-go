package client

import (
	"fmt"
	"io"
	"net/url"

	"github.com/drengskapr/regru-api-go/connector"
)

// ApiRequest makes a POST request with the given fields and returns the response body.
func ApiRequest(reqUrl string, postFields map[string]string) ([]byte, error) {
	postData := url.Values{}
	for k, v := range postFields {
		postData.Add(k, v)
	}

	c := connector.NewConnection()
	res, err := c.PostForm(reqUrl, postData)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status %d: %s", res.StatusCode, body)
	}
	return body, nil
}
