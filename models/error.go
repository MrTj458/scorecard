package models

type Error struct {
	StatusCode int           `json:"status_code"`
	Detail     string        `json:"detail"`
	Fields     []*ErrorField `json:"fields"`
}

type ErrorField struct {
	Location string `json:"location"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
}

func NewError(statusCode int, detail string) *Error {
	return &Error{
		StatusCode: statusCode,
		Detail:     detail,
		Fields:     make([]*ErrorField, 0),
	}
}

func NewErrorWithFields(statusCode int, detail string, fields []*ErrorField) *Error {
	return &Error{
		StatusCode: statusCode,
		Detail:     detail,
		Fields:     fields,
	}
}
