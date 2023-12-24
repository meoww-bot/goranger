package client

import (
	"encoding/json"

	"github.com/meoww-bot/goranger/model"
)

func (c *RangerClient) CreateService(service model.RangerService) (model.RangerService, error) {

	resp, err := c.CallAPI(CREATE_SERVICE, nil, service)

	if err != nil {
		return service, err
	}

	err = json.Unmarshal(resp.Body(), &service)

	return service, err

}

func (c *RangerClient) FindServices(filter model.P) (services []model.RangerService, err error) {

	resp, err := c.CallAPI(FIND_SERVICES, nil, nil)

	if err != nil {
		return services, err
	}

	err = json.Unmarshal(resp.Body(), &services)

	for index, service := range services {

		switch service.Type {
		case model.RangerServiceTypeHdfs:
			var hdfsConfig model.RangerServiceHdfsConfig
			var raw []byte
			raw, err = json.Marshal(service.Configs)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(raw, &hdfsConfig)
			if err != nil {
				panic(err)
			}

			services[index].Configs = hdfsConfig

		}
	}

	return services, err

}

// def get_service_by_id(self, serviceId):
// resp = self.__call_api(RangerClient.GET_SERVICE_BY_ID.format_path({ 'id': serviceId }))

// return type_coerce(resp, RangerService)

// def get_service(self, serviceName):
// resp = self.__call_api(RangerClient.GET_SERVICE_BY_NAME.format_path({ 'name': serviceName }))

// return type_coerce(resp, RangerService)

// def update_service_by_id(self, serviceId, service, params=None):
// resp = self.__call_api(RangerClient.UPDATE_SERVICE_BY_ID.format_path({ 'id': serviceId }), params, service)

// return type_coerce(resp, RangerService)

// def update_service(self, serviceName, service, params=None):
// resp = self.__call_api(RangerClient.UPDATE_SERVICE_BY_NAME.format_path({ 'name': serviceName }), params, service)

// return type_coerce(resp, RangerService)

// def delete_service_by_id(self, serviceId):
// self.__call_api(RangerClient.DELETE_SERVICE_BY_ID.format_path({ 'id': serviceId }))

// def delete_service(self, serviceName):
// self.__call_api(RangerClient.DELETE_SERVICE_BY_NAME.format_path({ 'name': serviceName }))
