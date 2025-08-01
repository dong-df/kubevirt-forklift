package handler

import (
	api "github.com/kubev2v/forklift/pkg/apis/forklift/v1beta1"
	"github.com/kubev2v/forklift/pkg/controller/plan/handler/ocp"
	"github.com/kubev2v/forklift/pkg/controller/plan/handler/openstack"
	"github.com/kubev2v/forklift/pkg/controller/plan/handler/ova"
	"github.com/kubev2v/forklift/pkg/controller/plan/handler/ovirt"
	"github.com/kubev2v/forklift/pkg/controller/plan/handler/vsphere"
	"github.com/kubev2v/forklift/pkg/controller/watch/handler"
	liberr "github.com/kubev2v/forklift/pkg/lib/error"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type WatchManager = handler.WatchManager

// Inventory event handler.
type Handler interface {
	// Ensure watch started.
	Watch(m *handler.WatchManager) error
}

// Handler factory.
func New(
	client client.Client,
	channel chan event.GenericEvent,
	provider *api.Provider) (h Handler, err error) {
	//
	switch provider.Type() {
	case api.OpenShift:
		h, err = ocp.New(
			client,
			channel,
			provider)
	case api.VSphere:
		h, err = vsphere.New(
			client,
			channel,
			provider)
	case api.OVirt:
		h, err = ovirt.New(
			client,
			channel,
			provider)
	case api.OpenStack:
		h, err = openstack.New(
			client,
			channel,
			provider)
	case api.Ova:
		h, err = ova.New(
			client,
			channel,
			provider)
	default:
		err = liberr.New("provider not supported.")
	}

	return
}
