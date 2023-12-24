package model

// https://ranger.apache.org/apidocs/json_RoleMember.html
type RoleMember struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}

func NewRoleMember() RoleMember {
	return RoleMember{
		Name:    "",
		IsAdmin: false,
	}
}

// https://ranger.apache.org/apidocs/json_RangerRole.html
type RangerRole struct {
	RangerBaseModelObject
	Groups        []RoleMember `json:"groups"`
	Users         []RoleMember `json:"users"`
	Options       Options      `json:"options"` // TODO: map of object
	CreatedByUser string       `json:"createdByUser"`
	Roles         []RoleMember `json:"roles"`
	Description   string       `json:"description"`
	Name          string       `json:"name"`
}
