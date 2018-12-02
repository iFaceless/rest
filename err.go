package rest

var (
	HTTPBadRequestError       = NewAPIError(400, 400, "HTTPBadRequestError", "bad request")
	HTTPMethodNotAllowedError = NewAPIError(405, 405, "HTTPMethodNotAllowedError", "http method not allowed")
	HTTPResourceNotFoundError = NewAPIError(404, 404, "HTTPResourceNotFoundError","resource not found")
)

type DetailError struct {
	Code int `json:"code"`
	Name string `json:"name"`
	Message string `json:"message"`
}

type APIError struct {
	HTTPStatusCode int `json:"-"`
	DetailError    DetailError `json:"error"`
}

func (e *APIError) Error()string  {
	return e.DetailError.Message
}

func NewAPIError(httpStatusCode, serviceCode int, name, message string) *APIError {
	return &APIError{
		HTTPStatusCode: httpStatusCode,
		DetailError:DetailError{
			Code:serviceCode,
			Name:name,
			Message:message,
		},
	}
}