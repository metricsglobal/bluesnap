package bluesnap

import (
	"errors"

	"github.com/metricsglobal/bluesnap/card"
)

func (c Connector) Sale(input Serializer, output Deserializer, opts Opts) (Errors, error) {
	if input.Method() != output.Method() {
		return emptyErrors(), errors.New("input method differs from output method")
	}

	switch input.Method() {
	case card.Method:
		return c.do("POST", "/services/2/transactions", input, output, opts)
	}

	return emptyErrors(), errors.New("invalid method passed")
}

func (c Connector) Auth(input Serializer, output Deserializer, opts Opts) (Errors, error) {
	if input.Method() != output.Method() {
		return emptyErrors(), errors.New("input method differs from output method")
	}

	switch input.Method() {
	case card.Method:
		return c.do("POST", "/services/2/transactions", input, output, opts)
	}

	return emptyErrors(), errors.New("invalid method passed")
}

func (c Connector) Capture(input Serializer, output Deserializer, opts Opts) (Errors, error) {
	if input.Method() != output.Method() {
		return emptyErrors(), errors.New("input method differs from output method")
	}

	switch input.Method() {
	case card.Method:
		return c.do("POST", "/services/2/transactions", input, output, opts)
	}

	return emptyErrors(), errors.New("invalid method passed")
}

func (c Connector) AuthReversal(input Serializer, output Deserializer, opts Opts) (Errors, error) {
	if input.Method() != output.Method() {
		return emptyErrors(), errors.New("input method differs from output method")
	}

	switch input.Method() {
	case card.Method:
		return c.do("POST", "/services/2/transactions", input, output, opts)
	}

	return emptyErrors(), errors.New("invalid method passed")
}

func (c Connector) Retrieve(transactionID string, output Deserializer, opts Opts) (Errors, error) {
	switch output.Method() {
	case card.Method:
		return c.do("POST", "/services/2/transactions/"+transactionID, nil, output, opts)
	}

	return emptyErrors(), errors.New("invalid method passed")
}
