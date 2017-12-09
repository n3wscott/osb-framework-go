package v2

type BrokerController interface {
	CatalogController
	InstanceController
	BindingController
}
