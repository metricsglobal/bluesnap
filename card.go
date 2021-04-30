package bluesnap

func (c Connector) CardSale(input Serializer, output Deserializer) error {
	return c.do("POST", "/services/2/transactions", input, output)
}

func (c Connector) CardAuth(input Serializer, output Deserializer) error {
	return c.do("POST", "/services/2/transactions", input, output)
}

func (c Connector) CardCapture(input Serializer, output Deserializer) error {
	return c.do("POST", "/services/2/transactions", input, output)
}

func (c Connector) CardAuthReversal(input Serializer, output Deserializer) error {
	return c.do("POST", "/services/2/transactions", input, output)
}

func (c Connector) CardRetrieve(transactionID string, output Deserializer) error {
	return c.do("POST", "/services/2/transactions/"+transactionID, nil, output)
}
