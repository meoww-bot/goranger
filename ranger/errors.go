package ranger

import (
	"fmt"

	"github.com/ahuigo/requests"
)

type RangerServiceError struct {
	Method         HttpMethod //api.method.name
	Path           string     //api.path
	ExpectedStatus int        //api.expected_status
	RespStatusCode int
	RespContent    string
}

func (e *RangerServiceError) Error() error {
	return fmt.Errorf("%s %s failed: expected_status=\"%d\", status=\"%d\", msg=\"%s\"", e.Method, e.Path, e.ExpectedStatus, e.RespStatusCode, e.RespContent)
}

func NewRangerServiceError(api API, resp *requests.Response) RangerServiceError {
	return RangerServiceError{
		Method:         api.Method,
		Path:           api.Path,
		ExpectedStatus: api.ExpectedStatus,
		RespStatusCode: resp.StatusCode(),
		RespContent:    resp.Text(),
	}
}
