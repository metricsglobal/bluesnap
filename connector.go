package bluesnap

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
		Client:      client,
		url:         url,
	}
}

func (c Connector) do(method, endpoint string, input Serializer, output Deserializer, opts Opts) error {
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

	req.Header.Add("Authorization", "Basic "+opts.Credentials.Parse())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.Client.Do(req)
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

func (c Credentials) Parse() string {
	return base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
}
