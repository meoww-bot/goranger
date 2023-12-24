package model

// https://ranger.apache.org/apidocs/json_RangerSecurityZoneService.html
type RangerSecurityZoneService struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Property1 []string `json:"property1"`
	Property2 []string `json:"property2"`
}

// https://ranger.apache.org/apidocs/json_RangerSecurityZone.html
type RangerSecurityZone struct {
	RangerBaseModelObject
	AdminUsers      []string                             `json:"adminUsers"`
	AuditUserGroups []string                             `json:"auditUserGroups"`
	AuditUsers      []string                             `json:"auditUsers"`
	Description     string                               `json:"description"`
	Services        map[string]RangerSecurityZoneService `json:"services"`
	Name            string                               `json:"name"`
	TagServices     []string                             `json:"tagServices"`
	AdminUserGroups []string                             `json:"adminUserGroups"`
}
