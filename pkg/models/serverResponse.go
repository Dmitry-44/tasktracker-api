package models

const (
	StatusSuccess = "ok"
	StatusError   = "error"
)

type ServerResponse struct {
	Status       string `json:"status"`
	Data         string `json:"data"`
	ErrorMessage string `json:"error"`
}

func (r *ServerResponse) Success() *ServerResponse {
	response := &ServerResponse{
		Status: StatusSuccess,
	}
	return response
}

func (r *ServerResponse) Error() *ServerResponse {
	response := &ServerResponse{
		Status: StatusError,
	}
	return response
}

var ErrorServerResponse = ServerResponse{
	Status: "error",
}
