package bluesnap

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Serializer interface {
	ToJSON() ([]byte, error)
	Method() string
}

type Deserializer interface {
	FromJSON(data []byte) error
	Method() string
}

type Connector struct {
	Client *http.Client
	url    string
}

type Opts struct {
	Credentials Credentials
}

type Credentials struct {
	Username string
	Password string
}

func New(client *http.Client, url string) Connector {
	return Connector{
		Client: client,
		url:    url,
	}
}

func (c Connector) do(method, endpoint string, input Serializer, output Deserializer, opts Opts) error {
	if reflect.ValueOf(output).Kind() != reflect.Ptr {
		return errors.New("output must be a pointer")
	}

	var buf *bytes.Buffer = nil
	if input != nil {
		body, err := input.ToJSON()
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, c.getURL(endpoint), buf)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+opts.Credentials.Parse())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("invalid status code [%d]", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if output != nil {
		return output.FromJSON(respBody)
	}
	return nil
}

func (c Connector) getURL(endpoint string) string {
	return c.url + endpoint
}

func (c Credentials) Parse() string {
	return base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
}
