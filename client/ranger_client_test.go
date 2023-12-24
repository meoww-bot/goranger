package client

import (
	"encoding/json"
	"testing"

	"github.com/meoww-bot/goranger/model"

	"github.com/ahuigo/requests"
)

const (
	ranger_url      = "http://192.168.233.80:6080"
	ranger_username = "admin"
	ranger_password = "rangeradmin1"
)

var (
	PolicyToDelete model.RangerPolicy
	Client         *RangerClient
	err            error
)

var ranger_auth = requests.Auth{ranger_username, ranger_password}

func PrettyPrint(t *testing.T, data interface{}) {

	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%s \n", p)
}

func TestNewClient(t *testing.T) {
	Client, err = NewClient(ranger_url, ranger_auth)
	if err != nil {
		t.Error(err)
	}

}

// func TestGetPolicyByID1(t *testing.T) {

// 	PolicyId := 1

// 	policy, err := Client.GetPolicyById(PolicyId)

// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("Get Policy success")
// 		PrettyPrint(t, policy)
// 	}

// }

// func TestGetPolicyByID2(t *testing.T) {

// 	PolicyId := 2

// 	policy, err := Client.GetPolicyById(PolicyId)

// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Log("Get Policy success")
// 		PrettyPrint(t, policy)
// 	}

// }

// func TestGetPolicyByNotFoundID(t *testing.T) {

// 	PolicyId := 0

// 	policy, err := Client.GetPolicyById(PolicyId)

// 	if err == nil {

// 		t.Log("Expected error not received, but marking test as passed")
// 		t.Fail()
// 	} else {
// 		t.Log(err)
// 		PrettyPrint(t, policy)
// 	}

// }

// func TestDeletePolicyByName(t *testing.T) {

// 	err := Client.DeletePolicy("hdfs", "test_policy")

// 	if err != nil {
// 		t.Error(err)
// 	} else {

// 		t.Logf("Delete policy success")
// 	}

// }

// func TestCreatePolicyWithDunpName(t *testing.T) {

// 	policy := model.NewRangerPolicy()

// 	policy.Name = "all - path"
// 	policy.Service = "hdfs"
// 	policy.Resources = model.Resources{
// 		Path: model.RangerPolicyResource{
// 			Values:      []string{"/"},
// 			IsRecursive: true,
// 		}}

// 	allowItem := model.RangerPolicyItem{}
// 	allowItem.Users = []string{"admin1"}

// 	allowItem.Accesses = []model.RangerPolicyItemAccess{
// 		model.NewRangerPolicyItemAccess(model.RangerPolicyItemAccessTypeRead),
// 		model.NewRangerPolicyItemAccess(model.RangerPolicyItemAccessTypeExecute),
// 	}

// 	policy.PolicyItems = []model.RangerPolicyItem{allowItem}

// 	PolicyToDelete, err = Client.CreatePolicy(policy)

// 	if err == nil {
// 		t.Log("Expected error not received, but marking test as passed")
// 		t.Fail()
// 		t.Logf("Create policy success. Name: %s, PolicyId: %d", PolicyToDelete.Name, PolicyToDelete.Id)
// 		return
// 	} else {
// 		t.Log(err)

// 	}

// }

// func TestCreatePolicy(t *testing.T) {

// 	policy := model.NewRangerPolicy()

// 	policy.Name = "test_policy"
// 	policy.Service = "hdfs"
// 	policy.Resources = model.Resources{
// 		Path: model.RangerPolicyResource{
// 			Values:      []string{"/"},
// 			IsRecursive: true,
// 		}}

// 	allowItem := model.RangerPolicyItem{}
// 	allowItem.Users = []string{"admin1"}

// 	allowItem.Accesses = []model.RangerPolicyItemAccess{
// 		model.NewRangerPolicyItemAccess(model.RangerPolicyItemAccessTypeRead),
// 		model.NewRangerPolicyItemAccess(model.RangerPolicyItemAccessTypeExecute),
// 	}

// 	policy.PolicyItems = []model.RangerPolicyItem{allowItem}

// 	PolicyToDelete, err = Client.CreatePolicy(policy)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Logf("Create policy success. Name: %s, PolicyId: %d", PolicyToDelete.Name, PolicyToDelete.Id)

// }

// func TestDeletePolicyById(t *testing.T) {

// 	err := Client.DeletePolicyById(PolicyToDelete.Id)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Logf("Delete policy success. PolicyId: %d , PolicyName: %s", PolicyToDelete.Id, PolicyToDelete.Name)

// }

// func TestCreatePolicy2(t *testing.T) {

// 	policy := model.NewRangerPolicy()

// 	policy.Name = "test_policy"
// 	policy.Service = "hdfs"
// 	policy.Resources = model.Resources{
// 		Path: model.RangerPolicyResource{
// 			Values:      []string{"/"},
// 			IsRecursive: true,
// 		}}

// 	allowItem := model.RangerPolicyItem{}
// 	allowItem.Users = []string{"admin1"}

// 	allowItem.Accesses = []model.RangerPolicyItemAccess{
// 		model.NewRangerPolicyItemAccess(model.RangerPolicyItemAccessTypeRead),
// 		model.NewRangerPolicyItemAccess(model.RangerPolicyItemAccessTypeExecute),
// 	}

// 	policy.PolicyItems = []model.RangerPolicyItem{allowItem}

// 	PolicyToDelete, err = Client.CreatePolicy(policy)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Logf("Create policy success. Name: %s, PolicyId: %d", PolicyToDelete.Name, PolicyToDelete.Id)

// }

// func TestFindPolicies(t *testing.T) {

// 	filter := model.P{}

// 	var policies []model.RangerPolicy

// 	policies, err := Client.FindPolicies(filter)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Logf("Find policy success. Policy Count: %d", len(policies))

// 	// PrettyPrint(t, policies)

// }

// func TestGetPolicy(t *testing.T) {

// 	var policy model.RangerPolicy

// 	serviceName := "hdfs"
// 	policyName := "all - path"

// 	policy, err := Client.GetPolicy(serviceName, policyName)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	PrettyPrint(t, policy)

// }

// func TestGetPolicyInService(t *testing.T) {

// 	var policies []model.RangerPolicy

// 	serviceName := "hdfs"

// 	policies, err := Client.GetPolicyInService(serviceName)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Logf("Found %d policies in service [%s] :", len(policies), serviceName)

// 	// PrettyPrint(t, policies)

// 	for index, policy := range policies {
// 		t.Logf("%d. %s", index, policy.Name)
// 	}

// }

func TestFindServices(t *testing.T) {
	var services []model.RangerService

	filter := model.P{}

	services, err := Client.FindServices(filter)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Found %d service", len(services))

	PrettyPrint(t, services)

	for index, service := range services {
		t.Logf("%d. %s", index, service.Name)

		if service.Type == "hdfs" {
			if value, ok := service.Configs.(model.RangerServiceHdfsConfig); ok {
				t.Log(value.FsDefaultName)
			} else {
				t.Log(ok)
			}
		}

	}

}

// func TestCreateHdfsServices(t *testing.T) {

// 	service := model.RangerService{
// 		Type:        model.RangerServiceTypeHdfs,
// 		Name:        "hdfs1",
// 		DisplayName: "hdfs1",
// 	}

// 	service.IsEnabled = true

// 	serviceConfig := model.NewRangerServiceHdfsConfig()
// 	serviceConfig.FsDefaultName = "http://namenode:9000"
// 	serviceConfig.Username = "admin"
// 	serviceConfig.Password = "rangeradmin1"

// 	service.Configs = serviceConfig

// 	serviceResult, err := Client.CreateService(service)

// 	if err != nil {
// 		t.Error(err)
// 	} else {
// 		t.Logf("Create service [%s] success", serviceResult.Name)
// 		PrettyPrint(t, serviceResult)

// 	}

// }
