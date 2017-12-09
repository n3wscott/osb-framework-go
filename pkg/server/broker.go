package server

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/n3wscott/osb-framework-go/pkg/apis/broker/v2"
	"github.com/golang/glog"
)

// TODO look for the service broker headers.
// TODO the query params are not in 100% yet. fix dis.
// swagger: http://petstore.swagger.io/?config=https://raw.githubusercontent.com/n3wscott/servicebroker/9bc40f42ecaa3ab89fe632404a3eb56f15fb03b8/openapi.yaml

func (s *server) GetCatalog(w http.ResponseWriter, req *http.Request) {
	response, err := s.Controller.GetCatalog()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *server) CreateServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instanceID := vars["instance_id"]
	acceptsIncomplete := vars["accepts_incomplete"]

	if acceptsIncomplete == "true" {
		glog.Info("accepts_incomplete: true")
	}

	var request v2.CreateServiceInstanceRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	response, code, err := s.Controller.CreateServiceInstance(instanceID, &request)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *server) UpdateServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instanceID := vars["instance_id"]
	acceptsIncomplete := vars["accepts_incomplete"]

	if acceptsIncomplete == "true" {
		glog.Info("accepts_incomplete: true")
	}

	var request v2.CreateServiceInstanceRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	response, code, err := s.Controller.UpdateServiceInstance(instanceID, &request)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *server) DeleteServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instanceID := vars["instance_id"]
	serviceID := vars["service_id"]
	planID := vars["plan_id"]
	acceptsIncomplete := vars["accepts_incomplete"]

	if acceptsIncomplete == "true" {
		glog.Info("accepts_incomplete: true")
	}

	var request v2.DeleteServiceInstanceRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	request.ServiceID = serviceID
	request.PlanID = planID

	response, code, err := s.Controller.DeleteServiceInstance(instanceID, &request)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *server) PollServiceInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instanceID := vars["instance_id"]
	serviceID := vars["service_id"]
	planID := vars["plan_id"]
	operation := vars["operation"]

	request := v2.LastOperationRequest{
		ServiceID: serviceID,
		PlanID: planID,
		Operation:operation,
	}

	request.ServiceID = serviceID
	request.PlanID = planID
	request.Operation = operation

	response, code, err := s.Controller.PollServiceInstance(instanceID, &request)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *server) CreateServiceBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instanceID := vars["instance_id"]
	bindingID := vars["binding_id"]

	var request v2.BindingRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	response, err := s.Controller.CreateServiceBinding(instanceID, bindingID, &request)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (s *server) DeleteServiceBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instanceID := vars["instance_id"]
	bindingID := vars["binding_id"]
	serviceID := vars["service_id"]
	planID := vars["plan_id"]

	err := s.Controller.DeleteServiceBinding(instanceID, bindingID, serviceID, planID)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode("{}")
}
