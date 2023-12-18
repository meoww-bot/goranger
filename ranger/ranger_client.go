package ranger

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/ahuigo/requests"
)

const APPLICATION_JSON = "application/json"

// Client is a Ranger API client.
type RangerClient struct {
	BaseURL string
	Auth    requests.Auth
	Req     *requests.Session
	Header  requests.Header
}

// New creates a new Ranger client.
func NewClient(baseURL string, auth requests.Auth) (*RangerClient, error) {

	session := requests.R()

	// 2. fake CA certificate
	// session.SetCaCert("conf/rootCA.crt")

	// 3. skip ssl
	session = session.SkipSsl(true)

	req := session

	resp, err := req.Get(baseURL, auth)
	if resp.StatusCode() == http.StatusOK {
		slog.Info("login success")
	}

	client := new(RangerClient)

	client.Auth = auth

	client.BaseURL = baseURL

	client.Req = req

	client.Header = requests.Header{"Accept": APPLICATION_JSON, "Content-type": APPLICATION_JSON}

	return client, err

}

func (c *RangerClient) CallAPI(api API, query_params requests.Params, request_data interface{}) (*requests.Response, error) {

	// return type_coerce(resp, RangerPolicy)
	var resp *requests.Response
	var err error
	var API_URL string

	if api.FormattedPath != "" {
		API_URL, err = url.JoinPath(c.BaseURL, api.FormattedPath)

	} else {
		API_URL, err = url.JoinPath(c.BaseURL, api.Path)

	}

	slog.Info("call api ", "url", API_URL, "method", api.Method)

	var data requests.Jsoni
	var params requests.Params

	if request_data != nil {
		data = requests.Jsoni(request_data)
	}

	if query_params != nil {
		params = query_params
	}

	if api.Method == HttpMethodGET {
		resp, err = c.Req.Get(API_URL, c.Auth, c.Header)
	}

	if api.Method == HttpMethodPOST {
		resp, err = c.Req.Post(API_URL, c.Auth, c.Header, params, data)
	}

	if api.Method == HttpMethodPUT {
		resp, err = c.Req.Put(API_URL, c.Auth, c.Header, params, data)
	}

	if api.Method == HttpMethodDELETE {
		resp, err = c.Req.Delete(API_URL, c.Auth, c.Header, params, data)
	}

	var ret *requests.Response

	if len(resp.Text()) == 0 {
		ret = nil
	}

	if resp.StatusCode() == int(api.ExpectedStatus) {
		if resp.StatusCode() == http.StatusNoContent || len(resp.Text()) == 0 {
			ret = nil
		} else {
			ret = resp
		}
	} else if resp.StatusCode() == http.StatusServiceUnavailable {
		slog.Error("Ranger admin unavailable. ", "HTTP StatusCode", http.StatusServiceUnavailable)

		ret = nil

	} else {
		rserror := NewRangerServiceError(api, resp)
		err = rserror.Error()
	}
	// fmt.Println(string(resp.Body()))
	return ret, err
}
