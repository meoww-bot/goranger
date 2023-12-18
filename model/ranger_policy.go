package model

// https://ranger.apache.org/apidocs/json_RangerPolicy.html
type RangerPolicy struct {
	RangerBaseModelObject
	Service              string                      `json:"service"`
	Name                 string                      `json:"name"`
	PolicyType           PolicyType                  `json:"policyType"`
	PolicyPriority       int                         `json:"policyPriority"`
	Description          string                      `json:"description"`
	ResourceSignature    string                      `json:"resourceSignature"`
	IsAuditEnabled       bool                        `json:"isAuditEnabled"`
	Resources            Resources                   `json:"resources"`
	AdditionalResources  []RangerPolicyResource      `json:"additionalResources"`
	PolicyItems          []RangerPolicyItem          `json:"policyItems"`
	DenyPolicyItems      []RangerPolicyItem          `json:"denyPolicyItems"`
	AllowExceptions      []RangerPolicyItem          `json:"allowExceptions"`
	DenyExceptions       []RangerPolicyItem          `json:"denyExceptions"`
	DataMaskPolicyItems  []RangerDataMaskPolicyItem  `json:"dataMaskPolicyItems"`
	RowFilterPolicyItems []RangerRowFilterPolicyItem `json:"rowFilterPolicyItems"`
	ServiceType          string                      `json:"serviceType"`
	Options              Options                     `json:"options"` // TODO: map of object
	ValiditySchedules    []RangerValiditySchedule    `json:"validitySchedules"`
	PolicyLabels         []string                    `json:"policyLabels"`
	ZoneName             string                      `json:"zoneName"`
	Conditions           []RangerPolicyItemCondition `json:"conditions"`
	IsDenyAllElse        bool                        `json:"isDenyAllElse"`
}

func NewRangerPolicy() RangerPolicy {

	p := RangerPolicy{}
	p.IsEnabled = true
	p.IsAuditEnabled = true
	return p
}

type Options struct {
}

type Resources struct {
	Path RangerPolicyResource `json:"path"`
}

// type Resources struct {
// 	Path struct {
// 		RangerPolicyResource
// 	} `json:"path"`
// }

// https://ranger.apache.org/apidocs/json_RangerPolicyItem.html
type RangerPolicyItem struct {
	Accesses      []RangerPolicyItemAccess    `json:"accesses"`
	Users         []string                    `json:"users"`
	Groups        []string                    `json:"groups"`
	Roles         []string                    `json:"roles"`
	Conditions    []RangerPolicyItemCondition `json:"conditions"`
	DelegateAdmin bool                        `json:"delegateAdmin"`
}

// https://ranger.apache.org/apidocs/json_RangerPolicyItemCondition.html
type RangerPolicyItemCondition struct {
	Values []string `json:"values"`
	Type   string   `json:"type"`
}

// https://ranger.apache.org/apidocs/json_RangerPolicyItemAccess.html
type RangerPolicyItemAccess struct {
	IsAllowed bool                       `json:"isAllowed"`
	Type      RangerPolicyItemAccessType `json:"type"`
}

type RangerPolicyItemAccessType string

var (

	// hdfs
	RangerPolicyItemAccessTypeRead    RangerPolicyItemAccessType = "read"
	RangerPolicyItemAccessTypeWrite   RangerPolicyItemAccessType = "write"
	RangerPolicyItemAccessTypeExecute RangerPolicyItemAccessType = "execute"

	// hive
	RangerPolicyItemAccessTypeSelect       RangerPolicyItemAccessType = "select"
	RangerPolicyItemAccessTypeUpdate       RangerPolicyItemAccessType = "update"
	RangerPolicyItemAccessTypeCreate       RangerPolicyItemAccessType = "create"
	RangerPolicyItemAccessTypeDrop         RangerPolicyItemAccessType = "drop"
	RangerPolicyItemAccessTypeAlter        RangerPolicyItemAccessType = "alter"
	RangerPolicyItemAccessTypeIndex        RangerPolicyItemAccessType = "index"
	RangerPolicyItemAccessTypeLock         RangerPolicyItemAccessType = "lock"
	RangerPolicyItemAccessTypeAll          RangerPolicyItemAccessType = "all"
	RangerPolicyItemAccessTypeReplAdmin    RangerPolicyItemAccessType = "repladmin"
	RangerPolicyItemAccessTypeServiceAdmin RangerPolicyItemAccessType = "serviceadmin"
	RangerPolicyItemAccessTypeTempUDFAdmin RangerPolicyItemAccessType = "tempudfadmin"
	RangerPolicyItemAccessTypeRefresh      RangerPolicyItemAccessType = "refresh"
)

func NewRangerPolicyItemAccess(t RangerPolicyItemAccessType) RangerPolicyItemAccess {
	p := RangerPolicyItemAccess{}
	p.IsAllowed = true
	p.Type = t
	return p

}

// https://ranger.apache.org/apidocs/json_RangerPolicyResource.html
type RangerPolicyResource struct {
	IsExcludes  bool     `json:"isExcludes"`
	Values      []string `json:"values"`
	IsRecursive bool     `json:"isRecursive"`
}

// https://ranger.apache.org/apidocs/json_RangerDataMaskPolicyItem.html
type RangerDataMaskPolicyItem struct {
	RangerPolicyItem
	DataMaskInfo RangerPolicyItemDataMaskInfo `json:"dataMaskInfo"`
}

// https://ranger.apache.org/apidocs/json_RangerPolicyItemDataMaskInfo.html
type RangerPolicyItemDataMaskInfo struct {
	ConditionExpr string `json:"conditionExpr"`
	DataMaskType  string `json:"dataMaskType"`
	ValueExpr     string `json:"valueExpr"`
}

// https://ranger.apache.org/apidocs/json_RangerRowFilterPolicyItem.html
type RangerRowFilterPolicyItem struct {
	RangerPolicyItem
	RowFilterInfo RangerPolicyItemRowFilterInfo `json:"rowFilterInfo"`
}

// https://ranger.apache.org/apidocs/json_RangerPolicyItemRowFilterInfo.html
type RangerPolicyItemRowFilterInfo struct {
	FilterExpr string `json:"filterExpr"`
}

// https://ranger.apache.org/apidocs/json_RangerValiditySchedule.html
type RangerValiditySchedule struct {
	StartTime   string                     `json:"startTime"`
	EndTime     string                     `json:"endTime"`
	TimeZone    string                     `json:"timeZone"`
	Recurrences []RangerValidityRecurrence `json:"recurrences"`
}

// https://ranger.apache.org/apidocs/json_RangerValidityRecurrence.html
type RangerValidityRecurrence struct {
	Interval ValidityInterval   `json:"interval"`
	Schedule RecurrenceSchedule `json:"schedule"`
}

type ValidityInterval struct {
	Minutes int `json:"minutes"`
	Hours   int `json:"hours"`
	Days    int `json:"days"`
}

// https://ranger.apache.org/apidocs/json_RecurrenceSchedule.html
type RecurrenceSchedule struct {
	DayOfMonth string `json:"dayOfMonth"`
	Month      string `json:"month"`
	Hour       string `json:"hour"`
	DayOfWeek  string `json:"dayOfWeek"`
	Minute     string `json:"minute"`
	Year       string `json:"year"`
}
