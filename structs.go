package bluesnap

type ErrorMessage struct {
	ErrorName       string `json:"errorName"`
	Code            int64  `json:"code"`
	Description     string `json:"description"`
	InvalidProperty string `json:"invalidProperty"`
	FraudEvents     string `json:"fraudEvents"`
}

type Errors struct {
	StatusCode int `json:"-"`
	Messages []ErrorMessage `json:"message"`
}

func (e Errors) IsEmpty() bool {
	return len(e.Messages) == 0
}

func emptyErrors() Errors {
	return Errors{}
}