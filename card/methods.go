package card

import "encoding/json"

const Method = "card"

func (r Request) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r Request) Method() string {
	return Method
}

func (r *Response) FromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r Response) Method() string {
	return Method
}
