package models

// ActionParameter는 여러 메시지에서 공통으로 사용됩니다.
type ActionParameter struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// Action은 여러 메시지에서 공통으로 사용됩니다.
type Action struct {
	ActionType        string            `json:"actionType"`
	ActionID          string            `json:"actionId"`
	ActionDescription string            `json:"actionDescription,omitempty"`
	BlockingType      string            `json:"blockingType"`
	ActionParameters  []ActionParameter `json:"actionParameters,omitempty"`
}
