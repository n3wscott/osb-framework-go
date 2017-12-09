package server

import (
	"github.com/gorilla/mux"
	"github.com/n3wscott/osb-framework-go/pkg/apis/broker/v2"
)

type server struct {
	Router     *mux.Router
	Controller v2.BrokerController
}

func CreateServer(controller v2.BrokerController) *server {

	s := server{
		Router:     mux.NewRouter(),
		Controller: controller,
	}

	s.Router.HandleFunc("/v2/catalog", s.GetCatalog).Methods("GET")

	s.Router.HandleFunc("/v2/service_instances/{instance_id}", s.CreateServiceInstance).
		Queries("accepts_incomplete", "{accepts_incomplete}").
			Methods("PUT")

	s.Router.HandleFunc("/v2/service_instances/{instance_id}", s.UpdateServiceInstance).
		Queries("accepts_incomplete", "{accepts_incomplete}").
			Methods("PATCH")

	s.Router.HandleFunc("/v2/service_instances/{instance_id}", s.DeleteServiceInstance).
		Queries("service_id", "{service_id}", "plan_id", "{plan_id}", "accepts_incomplete", "{accepts_incomplete}").
			Methods("DELETE")

	s.Router.HandleFunc("/v2/service_instances/{instance_id}/last_operation", s.PollServiceInstance).
		Queries("service_id", "{service_id}", "plan_id", "{plan_id}", "operation", "{operation}").
			Methods("GET")

	s.Router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", s.CreateServiceBinding).Methods("PUT")

	s.Router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", s.DeleteServiceBinding).
		Queries("service_id", "{service_id}", "plan_id", "{plan_id}").
			Methods("DELETE")

	return &s
}
