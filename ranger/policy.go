package ranger

import (
	"encoding/json"

	"github.com/meoww-bot/goranger/model"

	"github.com/ahuigo/requests"
)

func (c *RangerClient) GetPolicyById(PolicyId int) (policy model.RangerPolicy, err error) {

	GET_POLICY_BY_ID.FormatPath(model.P{"id": PolicyId})
	resp, err := c.CallAPI(GET_POLICY_BY_ID, nil, nil)

	if err != nil {
		return policy, err
	}

	err = json.Unmarshal(resp.Body(), &policy)

	return policy, err

}

func (c *RangerClient) CreatePolicy(policy model.RangerPolicy) (model.RangerPolicy, error) {
	resp, err := c.CallAPI(CREATE_POLICY, nil, policy)

	if err != nil {
		return policy, err
	}

	err = json.Unmarshal(resp.Body(), &policy)

	return policy, err

}

func (c *RangerClient) DeletePolicyById(PolicyId int) (err error) {

	DELETE_POLICY_BY_ID.FormatPath(model.P{"id": PolicyId})
	_, err = c.CallAPI(DELETE_POLICY_BY_ID, nil, nil)

	return err

}

func (c *RangerClient) DeletePolicy(serviceName, policyName string) (err error) {

	query_params := requests.Params{"servicename": serviceName, "policyname": policyName}

	_, err = c.CallAPI(DELETE_POLICY_BY_NAME, query_params, nil)

	return err

}

func (c *RangerClient) FindPolicies(filter model.P) (policies []model.RangerPolicy, err error) {

	resp, err := c.CallAPI(FIND_POLICIES, nil, filter)

	if err != nil {
		return policies, err
	}

	err = json.Unmarshal(resp.Body(), &policies)

	return policies, err

}

func (c *RangerClient) GetPolicy(serviceName, policyName string) (policy model.RangerPolicy, err error) {

	GET_POLICY_BY_NAME.FormatPath(model.P{"serviceName": serviceName, "policyName": policyName})
	resp, err := c.CallAPI(GET_POLICY_BY_NAME, nil, nil)

	if err != nil {
		return policy, err
	}

	err = json.Unmarshal(resp.Body(), &policy)

	return policy, err

}

func (c *RangerClient) GetPolicyInService(serviceName string) (policies []model.RangerPolicy, err error) {

	GET_POLICIES_IN_SERVICE.FormatPath(model.P{"serviceName": serviceName})
	resp, err := c.CallAPI(GET_POLICIES_IN_SERVICE, nil, nil)

	if err != nil {
		return policies, err
	}

	err = json.Unmarshal(resp.Body(), &policies)

	return policies, err

}

func (c *RangerClient) UpdatePolicyById(policyId int, policy model.RangerPolicy) (model.RangerPolicy, error) {

	UPDATE_POLICY_BY_ID.FormatPath(model.P{"id": policyId})
	resp, err := c.CallAPI(UPDATE_POLICY_BY_ID, nil, policy)

	if err != nil {
		return policy, err
	}

	err = json.Unmarshal(resp.Body(), &policy)

	return policy, err

}

func (c *RangerClient) UpdatePolicy(serviceName, policyName string, policy model.RangerPolicy) (model.RangerPolicy, error) {

	UPDATE_POLICY_BY_NAME.FormatPath(model.P{"serviceName": serviceName, "policyName": policyName})
	resp, err := c.CallAPI(UPDATE_POLICY_BY_NAME, nil, policy)

	if err != nil {
		return policy, err
	}

	err = json.Unmarshal(resp.Body(), &policy)

	return policy, err

}

func (c *RangerClient) ApplyPolicy(policy model.RangerPolicy) (model.RangerPolicy, error) {

	resp, err := c.CallAPI(APPLY_POLICY, nil, policy)

	if err != nil {
		return policy, err
	}

	err = json.Unmarshal(resp.Body(), &policy)

	return policy, err

}
