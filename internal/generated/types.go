package api

type GenericResponse struct {
	Code    int         `json:"code"`    // http status code
	Message interface{} `json:"message"` // message
}
