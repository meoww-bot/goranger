package client

import (
	"fmt"

	"github.com/ahuigo/requests"
)

type RangerServiceError struct {
	Method         HttpMethod //api.method.name
	URL            string     //
	ExpectedStatus int        //api.expected_status
	RespStatusCode int
	RespContent    string
}

func (e *RangerServiceError) Error() error {
	return fmt.Errorf("%s %s failed: expected_status=\"%d\", status=\"%d\", msg=\"%s\"", e.Method, e.URL, e.ExpectedStatus, e.RespStatusCode, e.RespContent)
}

func NewRangerServiceError(URL string, api API, resp *requests.Response) RangerServiceError {
	return RangerServiceError{
		Method:         api.Method,
		URL:            URL,
		ExpectedStatus: api.ExpectedStatus,
		RespStatusCode: resp.StatusCode(),
		RespContent:    resp.Text(),
	}
}
