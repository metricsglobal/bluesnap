package bluesnap

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Serializer interface {
	ToJSON() ([]byte, error)
}

type Deserializer interface {
	FromJSON(data []byte) error
}

type Connector struct {
	client      *http.Client
	url         string
	credentials string
}

func New(url, credentials string) Connector {
	return Connector{
		client:      http.DefaultClient, // TODO
		url:         url,
		credentials: credentials,
	}
}

func (c Connector) do(method, endpoint string, input Serializer, output Deserializer) error {
	// TODO check output is a pointer
	body, err := input.ToJSON()
	if err != nil {
		return err
	}
	fmt.Println("Input: " + string(body))

	// TODO handle empty body
	req, err := http.NewRequest(method, c.getURL(endpoint), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+c.credentials)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("Status Code: " + strconv.FormatInt(int64(resp.StatusCode), 10))

	// TODO handle status_code

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Output: " + string(respBody))
	fmt.Println("----------------")
	//return nil
	if output != nil {
		return output.FromJSON(respBody)
	}
	return nil // TODO error
}

func (c Connector) getURL(endpoint string) string {
	return c.url + endpoint
}
