// Copyright 2018 NetApp, Inc. All Rights Reserved.
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	netappv1 "github.com/netapp/trident/persistent_store/kubernetes/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageClasses implements StorageClassInterface
type FakeStorageClasses struct {
	Fake *FakeTridentV1
}

var storageclassesResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "storageclasses"}

var storageclassesKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "StorageClass"}

// Get takes name of the storageClass, and returns the corresponding storageClass object, and an error if there is any.
func (c *FakeStorageClasses) Get(name string, options v1.GetOptions) (result *netappv1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storageclassesResource, name), &netappv1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.StorageClass), err
}

// List takes label and field selectors, and returns the list of StorageClasses that match those selectors.
func (c *FakeStorageClasses) List(opts v1.ListOptions) (result *netappv1.StorageClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storageclassesResource, storageclassesKind, opts), &netappv1.StorageClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.StorageClassList{}
	for _, item := range obj.(*netappv1.StorageClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageClasses.
func (c *FakeStorageClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storageclassesResource, opts))
}

// Create takes the representation of a storageClass and creates it.  Returns the server's representation of the storageClass, and an error, if there is any.
func (c *FakeStorageClasses) Create(storageClass *netappv1.StorageClass) (result *netappv1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storageclassesResource, storageClass), &netappv1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.StorageClass), err
}

// Update takes the representation of a storageClass and updates it. Returns the server's representation of the storageClass, and an error, if there is any.
func (c *FakeStorageClasses) Update(storageClass *netappv1.StorageClass) (result *netappv1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storageclassesResource, storageClass), &netappv1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.StorageClass), err
}

// Delete takes name of the storageClass and deletes it. Returns an error if one occurs.
func (c *FakeStorageClasses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storageclassesResource, name), &netappv1.StorageClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storageclassesResource, listOptions)

	_, err := c.Fake.Invokes(action, &netappv1.StorageClassList{})
	return err
}

// Patch applies the patch and returns the patched storageClass.
func (c *FakeStorageClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *netappv1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storageclassesResource, name, data, subresources...), &netappv1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.StorageClass), err
}