package model

// https://ranger.apache.org/apidocs/json_RangerService.html
type RangerService struct {
	RangerBaseModelObject
	Type             RangerServiceType `json:"type"`
	TagUpdateTime    int               `json:"tagUpdateTime"`
	TagService       string            `json:"tagService"`
	Description      string            `json:"description"`
	Name             string            `json:"name"`
	PolicyVersion    int               `json:"policyVersion"`
	PolicyUpdateTime int               `json:"policyUpdateTime"`
	DisplayName      string            `json:"displayName"`
	TagVersion       int               `json:"tagVersion"`
	Configs          interface{}       `json:"configs"`
}

type RangerServiceType string

const (
	RangerServiceTypeHdfs  RangerServiceType = "hdfs"
	RangerServiceTypeHive  RangerServiceType = "hive"
	RangerServiceTypeHbase RangerServiceType = "hbase"
)

type HadoopSecurityAuthenticationType string

const (
	HadoopSecurityAuthenticationSimple   HadoopSecurityAuthenticationType = "simple"
	HadoopSecurityAuthenticationKerberos HadoopSecurityAuthenticationType = "kerberos"
)

type HadoopSecurityAuthorizationType string

const (
	HadoopSecurityAuthorizationYes HadoopSecurityAuthorizationType = "true"
	HadoopSecurityAuthorizationNo  HadoopSecurityAuthorizationType = "false"
)

type HadoopRpcProtectionType string

const (
	HadoopRpcProtectionAuthentication HadoopRpcProtectionType = "authentication"
	HadoopRpcProtectionIntegrity      HadoopRpcProtectionType = "integrity"
	HadoopRpcProtectionPrivacy        HadoopRpcProtectionType = "privacy"
)

type RangerServiceHdfsConfig struct {
	FsDefaultName                string                           `json:"fs.default.name"`
	HadoopRpcProtection          HadoopRpcProtectionType          `json:"hadoop.rpc.protection"`
	HadoopSecurityAuthentication HadoopSecurityAuthenticationType `json:"hadoop.security.authentication"`
	HadoopSecurityAuthorization  HadoopSecurityAuthorizationType  `json:"hadoop.security.authorization"`
	Username                     string                           `json:"username"`
	Password                     string                           `json:"password"`
}

type RangerServiceHiveConfig struct {
	JdbcDriverClassName string `json:"jdbc.driverClassName"`
	JdbcURL             string `json:"jdbc.url"`
	Username            string `json:"username"`
	Password            string `json:"password"`
}

// use this to get default value
func NewRangerServiceHdfsConfig() RangerServiceHdfsConfig {

	config := RangerServiceHdfsConfig{
		HadoopRpcProtection:          HadoopRpcProtectionAuthentication,
		HadoopSecurityAuthentication: HadoopSecurityAuthenticationSimple,
		HadoopSecurityAuthorization:  HadoopSecurityAuthorizationYes,
	}

	return config

}

type RangerServiceHbaseConfig struct {
	HadoopSecurityAuthentication HadoopSecurityAuthenticationType `json:"hadoop.security.authentication"`
	HbaseSecurityAuthorization   HadoopSecurityAuthorizationType  `json:"hbase.security.authorization"`
}
