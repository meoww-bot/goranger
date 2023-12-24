package model

type RangerBase struct {
}

// https://ranger.apache.org/apidocs/json_RangerBaseModelObject.html
type RangerBaseModelObject struct {
	GUID       string `json:"guid"`
	IsEnabled  bool   `json:"isEnabled"`
	CreateTime int    `json:"createTime"`
	CreatedBy  string `json:"createdBy"`
	UpdatedBy  string `json:"updatedBy"`
	Id         int    `json:"id"`
	UpdateTime int    `json:"updateTime"`
	Version    int    `json:"version"`
}

func NewRangerBaseModelObject() RangerBaseModelObject {
	return RangerBaseModelObject{
		IsEnabled: true,
	}
}
