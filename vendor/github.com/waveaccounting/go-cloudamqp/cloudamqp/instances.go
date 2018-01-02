package cloudamqp

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

// Instance represents a CloudAMQP instance.
// Based on https://customer.cloudamqp.com/team/api.
type Instance struct {
	ID     int    `json:"id"`
	Plan   string `json:"plan"`
	Region string `json:"region"`
	Name   string `json:"name"`
	URL    string `json:"url,omitempty"`
	APIKey string `json:"apikey,omitempty"`
}

// InstanceService provides methods for accessing CloudAMQP instance API endpoints.
// https://customer.cloudamqp.com/team/api
type InstanceService struct {
	sling *sling.Sling
}

func newInstanceService(sling *sling.Sling) *InstanceService {
	return &InstanceService{
		sling: sling.Path("instances"),
	}
}

// List instances available to the authenticated session.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) List() ([]Instance, *http.Response, error) {
	instances := new([]Instance)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").Receive(instances, apiError)
	apiError.Status = resp.StatusCode
	return *instances, resp, relevantError(err, *apiError)
}

// Get a CloudAMQP instance.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) Get(id int) (*Instance, *http.Response, error) {
	instance := new(Instance)
	apiError := new(APIError)
	resp, err := s.sling.New().Path("instances/").Get(strconv.Itoa(id)).Receive(instance, apiError)
	apiError.Status = resp.StatusCode
	return instance, resp, relevantError(err, *apiError)
}

// CreateInstanceParams are the parameters for OrganizationService.Create.
type CreateInstanceParams struct {
	Name       string `url:"name"`
	Plan       string `url:"plan"`
	Region     string `url:"region"`
	VpcSubnet  string `url:"vpc_subnet,omitempty"`
	Nodes      int    `url:"nodes,omitempty"`
	RmqVersion string `url:"rmq_version,omitempty"`
}

// Create a new CloudAMP instance.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) Create(params *CreateInstanceParams) (*Instance, *http.Response, error) {
	instance := new(Instance)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("instances").BodyForm(params).Receive(instance, apiError)
	apiError.Status = resp.StatusCode
	return instance, resp, relevantError(err, *apiError)
}

// UpdateInstanceParams are the parameters for OrganizationService.Create.
type UpdateInstanceParams struct {
	Name  string `url:"name,omitempty"`
	Plan  string `url:"plan,omitempty"`
	Nodes int    `url:"nodes,omitempty"`
}

// Update a CloudAMQP instance.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) Update(id int, params *UpdateInstanceParams) (*Instance, *http.Response, error) {
	instance := new(Instance)
	apiError := new(APIError)
	// TODO: Figure out if this is actually returning errors correctly
	resp, err := s.sling.New().Path("instances/").Put(strconv.Itoa(id)).BodyForm(params).Receive(instance, apiError)
	log.Printf("%+v", resp)
	log.Printf("%+v", err)
	apiError.Status = resp.StatusCode
	log.Printf("apiError = %+v", apiError)
	return instance, resp, relevantError(err, *apiError)
}

// Delete a CloudAMQP instance.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) Delete(id int) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Path("instances/").Delete(strconv.Itoa(id)).Receive(nil, apiError)
	apiError.Status = resp.StatusCode
	return resp, relevantError(err, *apiError)
}
