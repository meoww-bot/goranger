package model

// https://ranger.apache.org/apidocs/json_RangerServiceDef.html
type RangerServiceDef struct {
	RangerBaseModelObject
	RbKeyLabel       string                     `json:"rbKeyLabel"`
	Resources        []RangerResourceDef        `json:"resources"`
	ContextEnrichers []RangerContextEnricherDef `json:"contextEnrichers"`
	Description      string                     `json:"description"`
	ImplClass        string                     `json:"implClass"`
	Enums            []RangerEnumDef            `json:"enums"`
	Label            string                     `json:"label"`
	Configs          []RangerServiceConfigDef   `json:"configs"`
	Options          map[string]string          `json:"options"`
	AccessTypes      []RangerAccessTypeDef      `json:"accessTypes"`
	DataMaskDef      RangerDataMaskDef          `json:"dataMaskDef"`
	PolicyConditions []RangerPolicyConditionDef `json:"policyConditions"`
	DisplayName      string                     `json:"displayName"`
	Name             string                     `json:"name"`
	RowFilterDef     RangerRowFilterDef         `json:"rowFilterDef"`
	RbKeyDescription string                     `json:"rbKeyDescription"`
}

// https://ranger.apache.org/apidocs/json_RangerServiceConfigDef.html
type RangerServiceConfigDef struct {
	ItemId                 int    `json:"itemId"`
	SubType                string `json:"subType"`
	Type                   string `json:"type"`
	Label                  string `json:"label"`
	RbKeyDescription       string `json:"rbKeyDescription"`
	RbKeyLabel             string `json:"rbKeyLabel"`
	DefaultValue           string `json:"defaultValue"`
	ValidationRegEx        string `json:"validationRegEx"`
	UiHint                 string `json:"uiHint"`
	RbKeyValidationMessage string `json:"rbKeyValidationMessage"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Mandatory              bool   `json:"mandatory"`
	ValidationMessage      string `json:"validationMessage"`
}

// https://ranger.apache.org/apidocs/json_RangerResourceDef.html
type RangerResourceDef struct {
	Type                   string            `json:"type"`
	MatcherOptions         map[string]string `json:"matcherOptions"`
	ExcludesSupported      bool              `json:"excludesSupported"`
	Parent                 string            `json:"parent"`
	Name                   string            `json:"name"`
	ValidationRegEx        string            `json:"validationRegEx"`
	Level                  int               `json:"level"`
	LookupSupported        bool              `json:"lookupSupported"`
	ItemId                 int               `json:"itemId"`
	ValidationMessage      string            `json:"validationMessage"`
	Mandatory              bool              `json:"mandatory"`
	AccessTypeRestrictions []string          `json:"accessTypeRestrictions"`
	UiHint                 string            `json:"uiHint"`
	Description            string            `json:"description"`
	Matcher                string            `json:"matcher"`
	RbKeyValidationMessage string            `json:"rbKeyValidationMessage"`
	Label                  string            `json:"label"`
	RbKeyDescription       string            `json:"rbKeyDescription"`
	IsValidLeaf            bool              `json:"isValidLeaf"`
	RbKeyLabel             string            `json:"rbKeyLabel"`
	RecursiveSupported     bool              `json:"recursiveSupported"`
}

// https://ranger.apache.org/apidocs/json_RangerAccessTypeDef.html
type RangerAccessTypeDef struct {
	Name          string   `json:"name"`
	Label         string   `json:"label"`
	ItemId        int      `json:"itemId"`
	ImpliedGrants []string `json:"impliedGrants"`
	RbKeyLabel    string   `json:"rbKeyLabel"`
}

// https://ranger.apache.org/apidocs/json_RangerPolicyConditionDef.html
type RangerPolicyConditionDef struct {
	EvaluatorOptions       map[string]string `json:"evaluatorOptions"`
	ItemId                 int               `json:"itemId"`
	Description            string            `json:"description"`
	ValidationMessage      string            `json:"validationMessage"`
	Label                  string            `json:"label"`
	Name                   string            `json:"name"`
	RbKeyValidationMessage string            `json:"rbKeyValidationMessage"`
	RbKeyLabel             string            `json:"rbKeyLabel"`
	ValidationRegEx        string            `json:"validationRegEx"`
	UiHint                 string            `json:"uiHint"`
	RbKeyDescription       string            `json:"rbKeyDescription"`
	Evaluator              string            `json:"evaluator"`
}

// https://ranger.apache.org/apidocs/json_RangerContextEnricherDef.html
type RangerContextEnricherDef struct {
	Enricher        string            `json:"enricher"`
	Name            string            `json:"name"`
	ItemId          int               `json:"itemId"`
	EnricherOptions map[string]string `json:"enricherOptions"`
}

// https://ranger.apache.org/apidocs/json_RangerEnumDef.html
type RangerEnumDef struct {
	DefaultIndex int                    `json:"defaultIndex"`
	Name         string                 `json:"name"`
	Elements     []RangerEnumElementDef `json:"elements"`
	ItemId       int                    `json:"itemId"`
}

// https://ranger.apache.org/apidocs/json_RangerDataMaskDef.html
type RangerDataMaskDef struct {
	MaskTypes   []RangerDataMaskTypeDef `json:"maskTypes"`
	AccessTypes []RangerAccessTypeDef   `json:"accessTypes"`
	Resources   []RangerResourceDef     `json:"resources"`
}

// https://ranger.apache.org/apidocs/json_RangerRowFilterDef.html
type RangerRowFilterDef struct {
	AccessTypes []RangerAccessTypeDef `json:"accessTypes"`
	Resources   []RangerResourceDef   `json:"resources"`
}

// https://ranger.apache.org/apidocs/json_RangerEnumElementDef.html
type RangerEnumElementDef struct {
	Name       string `json:"name"`
	Label      string `json:"label"`
	RbKeyLabel string `json:"rbKeyLabel"`
	ItemId     int    `json:"itemId"`
}

// https://ranger.apache.org/apidocs/json_RangerDataMaskTypeDef.html
type RangerDataMaskTypeDef struct {
	Transformer      string            `json:"transformer"`
	DataMaskOptions  map[string]string `json:"dataMaskOptions"`
	RbKeyLabel       string            `json:"rbKeyLabel"`
	ItemId           int               `json:"itemId"`
	Label            string            `json:"label"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	RbKeyDescription string            `json:"rbKeyDescription"`
}
