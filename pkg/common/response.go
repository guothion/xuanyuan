package common

var (
	StatusOk        = &Status{Code: 0, Message: "OK"}
	StatusForbidden = &Status{Code: 0, Message: "Forbidden"}
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

type CreateResponse struct {
	ID interface{} `json:"id"`
}
