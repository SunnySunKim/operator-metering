package fake

import (
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageLocations implements StorageLocationInterface
type FakeStorageLocations struct {
	Fake *FakeChargebackV1alpha1
	ns   string
}

var storagelocationsResource = schema.GroupVersionResource{Group: "chargeback.coreos.com", Version: "v1alpha1", Resource: "storagelocations"}

var storagelocationsKind = schema.GroupVersionKind{Group: "chargeback.coreos.com", Version: "v1alpha1", Kind: "StorageLocation"}

// Get takes name of the storageLocation, and returns the corresponding storageLocation object, and an error if there is any.
func (c *FakeStorageLocations) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagelocationsResource, c.ns, name), &v1alpha1.StorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageLocation), err
}

// List takes label and field selectors, and returns the list of StorageLocations that match those selectors.
func (c *FakeStorageLocations) List(opts v1.ListOptions) (result *v1alpha1.StorageLocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagelocationsResource, storagelocationsKind, c.ns, opts), &v1alpha1.StorageLocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageLocationList{}
	for _, item := range obj.(*v1alpha1.StorageLocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageLocations.
func (c *FakeStorageLocations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagelocationsResource, c.ns, opts))

}

// Create takes the representation of a storageLocation and creates it.  Returns the server's representation of the storageLocation, and an error, if there is any.
func (c *FakeStorageLocations) Create(storageLocation *v1alpha1.StorageLocation) (result *v1alpha1.StorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagelocationsResource, c.ns, storageLocation), &v1alpha1.StorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageLocation), err
}

// Update takes the representation of a storageLocation and updates it. Returns the server's representation of the storageLocation, and an error, if there is any.
func (c *FakeStorageLocations) Update(storageLocation *v1alpha1.StorageLocation) (result *v1alpha1.StorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagelocationsResource, c.ns, storageLocation), &v1alpha1.StorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageLocation), err
}

// Delete takes name of the storageLocation and deletes it. Returns an error if one occurs.
func (c *FakeStorageLocations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagelocationsResource, c.ns, name), &v1alpha1.StorageLocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageLocations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagelocationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageLocationList{})
	return err
}

// Patch applies the patch and returns the patched storageLocation.
func (c *FakeStorageLocations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagelocationsResource, c.ns, name, data, subresources...), &v1alpha1.StorageLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageLocation), err
}