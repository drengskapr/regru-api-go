package client

import (
	"io"
	"net/url"

	log "github.com/sirupsen/logrus"

	"github.com/drengskapr/regru-api-go/connector"
)

// ApiRequest make any POST request with default API settings and return response body.
func ApiRequest(reqUrl string, postFields map[string]string) (body []byte) {
	postData := url.Values{}
	for k, v := range postFields {
		postData.Add(k, v)
	}

	c := connector.NewConnection()
	res, err := c.PostForm(reqUrl, postData)
	if err != nil {
		log.Errorf("Connection error: %v", err)

		return
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Error(err)
	}
	return body
}
