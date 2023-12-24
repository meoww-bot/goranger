package client

import (
	"net/http"

	"github.com/meoww-bot/goranger/model"

	"github.com/slongfield/pyfmt"
)

type API struct {
	Path           string
	FormattedPath  string
	Method         HttpMethod
	ExpectedStatus int
	// Consumes       string
	// Produces       string
}

func (a *API) FormatPath(p model.P) {

	a.FormattedPath = pyfmt.Must(a.Path, p)
}

type HttpMethod string

const (
	HttpMethodGET    HttpMethod = "GET"
	HttpMethodPUT    HttpMethod = "PUT"
	HttpMethodPOST   HttpMethod = "POST"
	HttpMethodDELETE HttpMethod = "DELETE"
)

// URIs
const (
	URI_BASE = "service/public/v2/api"

	URI_SERVICEDEF         = URI_BASE + "/servicedef"
	URI_SERVICEDEF_BY_ID   = URI_SERVICEDEF + "/{id}"
	URI_SERVICEDEF_BY_NAME = URI_SERVICEDEF + "/name/{name}"

	URI_SERVICE             = URI_BASE + "/service"
	URI_SERVICE_BY_ID       = URI_SERVICE + "/{id}"
	URI_SERVICE_BY_NAME     = URI_SERVICE + "/name/{name}"
	URI_POLICIES_IN_SERVICE = URI_SERVICE + "/{serviceName}/policy"

	URI_POLICY         = URI_BASE + "/policy"
	URI_APPLY_POLICY   = URI_POLICY + "/apply"
	URI_POLICY_BY_ID   = URI_POLICY + "/{id}"
	URI_POLICY_BY_NAME = URI_SERVICE + "/{serviceName}/policy/{policyName}"

	URI_ROLE         = URI_BASE + "/roles"
	URI_ROLE_NAMES   = URI_ROLE + "/names"
	URI_ROLE_BY_ID   = URI_ROLE + "/{id}"
	URI_ROLE_BY_NAME = URI_ROLE + "/name/{name}"
	URI_USER_ROLES   = URI_ROLE + "/user/{name}"
	URI_GRANT_ROLE   = URI_ROLE + "/grant/{name}"
	URI_REVOKE_ROLE  = URI_ROLE + "/revoke/{name}"

	URI_ZONE         = URI_BASE + "/zones"
	URI_ZONE_BY_ID   = URI_ZONE + "/{id}"
	URI_ZONE_BY_NAME = URI_ZONE + "/name/{name}"

	URI_PLUGIN_INFO   = URI_BASE + "/plugins/info"
	URI_POLICY_DELTAS = URI_BASE + "/server/policydeltas"
)

// APIs
var (
	CREATE_SERVICEDEF         = API{Path: URI_SERVICEDEF, Method: HttpMethodPOST, ExpectedStatus: http.StatusOK}
	UPDATE_SERVICEDEF_BY_ID   = API{Path: URI_SERVICEDEF_BY_ID, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	UPDATE_SERVICEDEF_BY_NAME = API{Path: URI_SERVICEDEF_BY_NAME, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	DELETE_SERVICEDEF_BY_ID   = API{Path: URI_SERVICEDEF_BY_ID, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	DELETE_SERVICEDEF_BY_NAME = API{Path: URI_SERVICEDEF_BY_NAME, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	GET_SERVICEDEF_BY_ID      = API{Path: URI_SERVICEDEF_BY_ID, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_SERVICEDEF_BY_NAME    = API{Path: URI_SERVICEDEF_BY_NAME, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	FIND_SERVICEDEFS          = API{Path: URI_SERVICEDEF, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}

	CREATE_SERVICE         = API{Path: URI_SERVICE, Method: HttpMethodPOST, ExpectedStatus: http.StatusOK}
	UPDATE_SERVICE_BY_ID   = API{Path: URI_SERVICE_BY_ID, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	UPDATE_SERVICE_BY_NAME = API{Path: URI_SERVICE_BY_NAME, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	DELETE_SERVICE_BY_ID   = API{Path: URI_SERVICE_BY_ID, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	DELETE_SERVICE_BY_NAME = API{Path: URI_SERVICE_BY_NAME, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	GET_SERVICE_BY_ID      = API{Path: URI_SERVICE_BY_ID, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_SERVICE_BY_NAME    = API{Path: URI_SERVICE_BY_NAME, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	FIND_SERVICES          = API{Path: URI_SERVICE, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}

	CREATE_POLICY           = API{Path: URI_POLICY, Method: HttpMethodPOST, ExpectedStatus: http.StatusOK}
	UPDATE_POLICY_BY_ID     = API{Path: URI_POLICY_BY_ID, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	UPDATE_POLICY_BY_NAME   = API{Path: URI_POLICY_BY_NAME, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	APPLY_POLICY            = API{Path: URI_APPLY_POLICY, Method: HttpMethodPOST, ExpectedStatus: http.StatusOK}
	DELETE_POLICY_BY_ID     = API{Path: URI_POLICY_BY_ID, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	DELETE_POLICY_BY_NAME   = API{Path: URI_POLICY, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	GET_POLICY_BY_ID        = API{Path: URI_POLICY_BY_ID, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_POLICY_BY_NAME      = API{Path: URI_POLICY_BY_NAME, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_POLICIES_IN_SERVICE = API{Path: URI_POLICIES_IN_SERVICE, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	FIND_POLICIES           = API{Path: URI_POLICY, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}

	CREATE_ZONE         = API{Path: URI_ZONE, Method: HttpMethodPOST, ExpectedStatus: http.StatusOK}
	UPDATE_ZONE_BY_ID   = API{Path: URI_ZONE_BY_ID, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	UPDATE_ZONE_BY_NAME = API{Path: URI_ZONE_BY_NAME, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	DELETE_ZONE_BY_ID   = API{Path: URI_ZONE_BY_ID, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	DELETE_ZONE_BY_NAME = API{Path: URI_ZONE_BY_NAME, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	GET_ZONE_BY_ID      = API{Path: URI_ZONE_BY_ID, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_ZONE_BY_NAME    = API{Path: URI_ZONE_BY_NAME, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	FIND_ZONES          = API{Path: URI_ZONE, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}

	CREATE_ROLE         = API{Path: URI_ROLE, Method: HttpMethodPOST, ExpectedStatus: http.StatusOK}
	UPDATE_ROLE_BY_ID   = API{Path: URI_ROLE_BY_ID, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	DELETE_ROLE_BY_ID   = API{Path: URI_ROLE_BY_ID, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	DELETE_ROLE_BY_NAME = API{Path: URI_ROLE_BY_NAME, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
	GET_ROLE_BY_ID      = API{Path: URI_ROLE_BY_ID, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_ROLE_BY_NAME    = API{Path: URI_ROLE_BY_NAME, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_ALL_ROLE_NAMES  = API{Path: URI_ROLE_NAMES, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GET_USER_ROLES      = API{Path: URI_USER_ROLES, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	GRANT_ROLE          = API{Path: URI_GRANT_ROLE, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	REVOKE_ROLE         = API{Path: URI_REVOKE_ROLE, Method: HttpMethodPUT, ExpectedStatus: http.StatusOK}
	FIND_ROLES          = API{Path: URI_ROLE, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}

	GET_PLUGIN_INFO      = API{Path: URI_PLUGIN_INFO, Method: HttpMethodGET, ExpectedStatus: http.StatusOK}
	DELETE_POLICY_DELTAS = API{Path: URI_POLICY_DELTAS, Method: HttpMethodDELETE, ExpectedStatus: http.StatusNoContent}
)
