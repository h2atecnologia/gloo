// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type VirtualServiceWatcher interface {
	// watch namespace-scoped VirtualServices
	Watch(namespace string, opts clients.WatchOpts) (<-chan VirtualServiceList, <-chan error, error)
}

type VirtualServiceClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*VirtualService, error)
	Write(resource *VirtualService, opts clients.WriteOpts) (*VirtualService, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (VirtualServiceList, error)
	VirtualServiceWatcher
}

type virtualServiceClient struct {
	rc clients.ResourceClient
}

func NewVirtualServiceClient(rcFactory factory.ResourceClientFactory) (VirtualServiceClient, error) {
	return NewVirtualServiceClientWithToken(rcFactory, "")
}

func NewVirtualServiceClientWithToken(rcFactory factory.ResourceClientFactory, token string) (VirtualServiceClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &VirtualService{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base VirtualService resource client")
	}
	return NewVirtualServiceClientWithBase(rc), nil
}

func NewVirtualServiceClientWithBase(rc clients.ResourceClient) VirtualServiceClient {
	return &virtualServiceClient{
		rc: rc,
	}
}

func (client *virtualServiceClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *virtualServiceClient) Register() error {
	return client.rc.Register()
}

func (client *virtualServiceClient) Read(namespace, name string, opts clients.ReadOpts) (*VirtualService, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*VirtualService), nil
}

func (client *virtualServiceClient) Write(virtualService *VirtualService, opts clients.WriteOpts) (*VirtualService, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(virtualService, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*VirtualService), nil
}

func (client *virtualServiceClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *virtualServiceClient) List(namespace string, opts clients.ListOpts) (VirtualServiceList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToVirtualService(resourceList), nil
}

func (client *virtualServiceClient) Watch(namespace string, opts clients.WatchOpts) (<-chan VirtualServiceList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	virtualServicesChan := make(chan VirtualServiceList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				virtualServicesChan <- convertToVirtualService(resourceList)
			case <-opts.Ctx.Done():
				close(virtualServicesChan)
				return
			}
		}
	}()
	return virtualServicesChan, errs, nil
}

func convertToVirtualService(resources resources.ResourceList) VirtualServiceList {
	var virtualServiceList VirtualServiceList
	for _, resource := range resources {
		virtualServiceList = append(virtualServiceList, resource.(*VirtualService))
	}
	return virtualServiceList
}
