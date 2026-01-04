package twentygo

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

func checkForError(resp *resty.Response, err error, errMessage string) error {
	if err != nil {
		return &APIError{
			Code:    0,
			Message: WrapError(err, errMessage).Error(),
			Type:    ParseAPIErrType(err),
		}
	}

	if resp == nil {
		return &APIError{
			Message: "empty response",
			Type:    ParseAPIErrType(err),
		}
	}

	if resp.IsError() {
		var errorResponse ErrorResponse
		err2 := json.Unmarshal(resp.Body(), &errorResponse)
		if err2 != nil {
			return err2
		}

		return &APIError{
			Code:    resp.StatusCode(),
			Message: strings.Join(errorResponse.Messages, ", "),
			Type:    ParseAPIErrType(err),
		}
	}

	return nil
}

func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("...%w...", err, message)
}

// HTTPErrorResponse is a model of an error response
type HTTPErrorResponse struct {
	Error       string `json:"error,omitempty"`
	Message     string `json:"errorMessage,omitempty"`
	Description string `json:"error_description,omitempty"`
}

// String returns a string representation of an error
func (e HTTPErrorResponse) String() string {
	var res strings.Builder
	if len(e.Error) > 0 {
		res.WriteString(e.Error)
	}
	if len(e.Message) > 0 {
		if res.Len() > 0 {
			res.WriteString(": ")
		}
		res.WriteString(e.Message)
	}
	if len(e.Description) > 0 {
		if res.Len() > 0 {
			res.WriteString(": ")
		}
		res.WriteString(e.Description)
	}
	return res.String()
}

// NotEmpty validates that error is not emptyp
func (e HTTPErrorResponse) NotEmpty() bool {
	return len(e.Error) > 0 || len(e.Message) > 0 || len(e.Description) > 0
}
