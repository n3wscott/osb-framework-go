package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/n3wscott/osb-framework-go/pkg/server"
	"github.com/n3wscott/osb-framework-go/pkg/controller"
	"fmt"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	broker := controller.BrokerController{}
	s := server.CreateServer(&broker)

	glog.Infof("Starting Broker, http://localhost%s", *addr)
	glog.Fatal(http.ListenAndServe(fmt.Sprintf("%s", *addr), s.Router))
}
