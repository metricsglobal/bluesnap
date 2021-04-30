package card

import "encoding/json"

func (r Request) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Response) FromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}
