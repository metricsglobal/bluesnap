package bluesnap

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
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

func (c Connector) do(method, endpoint string, input Serializer, output Deserializer, opts Opts) (Errors, error) {
	if reflect.ValueOf(output).Kind() != reflect.Ptr {
		return emptyErrors(), errors.New("output must be a pointer")
	}

	var buf *bytes.Buffer = nil
	if input != nil {
		body, err := input.ToJSON()
		if err != nil {
			return emptyErrors(), err
		}
		buf = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, c.getURL(endpoint), buf)
	if err != nil {
		return emptyErrors(), err
	}

	req.Header.Add("Authorization", "Basic "+opts.Credentials.Parse())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return emptyErrors(), err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyErrors(), err
	}

	if resp.StatusCode > 399 {
		var errs Errors
		if err := json.Unmarshal(respBody, &errs); err != nil {
			return emptyErrors(), err
		}
		errs.StatusCode = resp.StatusCode
		return errs, nil
	}

	if output != nil {
		return emptyErrors(), output.FromJSON(respBody)
	}
	return emptyErrors(), nil
}

func (c Connector) getURL(endpoint string) string {
	return c.url + endpoint
}

func (c Credentials) Parse() string {
	return base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
}
