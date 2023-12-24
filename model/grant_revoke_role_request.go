package model

// https://ranger.apache.org/apidocs/json_GrantRevokeRoleRequest.html
type GrantRevokeRoleRequest struct {
	GrantOption     bool     `json:"grantOption"`
	ClientIPAddress string   `json:"clientIPAddress"`
	Users           []string `json:"users"`
	Grantor         string   `json:"grantor"`
	TargetRoles     []string `json:"targetRoles"`
	RequestData     string   `json:"requestData"`
	SessionId       string   `json:"sessionId"`
	GrantorGroups   []string `json:"grantorGroups"`
	ClusterName     string   `json:"clusterName"`
	Roles           []string `json:"roles"`
	ClientType      string   `json:"clientType"`
	Groups          []string `json:"groups"`
}

func NewGrantRevokeRoleRequest() GrantRevokeRoleRequest {
	return GrantRevokeRoleRequest{
		GrantOption: false,
	}

}
