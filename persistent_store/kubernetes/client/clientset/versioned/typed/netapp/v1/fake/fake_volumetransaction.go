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

// FakeVolumeTransactions implements VolumeTransactionInterface
type FakeVolumeTransactions struct {
	Fake *FakeTridentV1
}

var volumetransactionsResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "volumetransactions"}

var volumetransactionsKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "VolumeTransaction"}

// Get takes name of the volumeTransaction, and returns the corresponding volumeTransaction object, and an error if there is any.
func (c *FakeVolumeTransactions) Get(name string, options v1.GetOptions) (result *netappv1.VolumeTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumetransactionsResource, name), &netappv1.VolumeTransaction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.VolumeTransaction), err
}

// List takes label and field selectors, and returns the list of VolumeTransactions that match those selectors.
func (c *FakeVolumeTransactions) List(opts v1.ListOptions) (result *netappv1.VolumeTransactionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumetransactionsResource, volumetransactionsKind, opts), &netappv1.VolumeTransactionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.VolumeTransactionList{}
	for _, item := range obj.(*netappv1.VolumeTransactionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeTransactions.
func (c *FakeVolumeTransactions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumetransactionsResource, opts))
}

// Create takes the representation of a volumeTransaction and creates it.  Returns the server's representation of the volumeTransaction, and an error, if there is any.
func (c *FakeVolumeTransactions) Create(volumeTransaction *netappv1.VolumeTransaction) (result *netappv1.VolumeTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumetransactionsResource, volumeTransaction), &netappv1.VolumeTransaction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.VolumeTransaction), err
}

// Update takes the representation of a volumeTransaction and updates it. Returns the server's representation of the volumeTransaction, and an error, if there is any.
func (c *FakeVolumeTransactions) Update(volumeTransaction *netappv1.VolumeTransaction) (result *netappv1.VolumeTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumetransactionsResource, volumeTransaction), &netappv1.VolumeTransaction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.VolumeTransaction), err
}

// Delete takes name of the volumeTransaction and deletes it. Returns an error if one occurs.
func (c *FakeVolumeTransactions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(volumetransactionsResource, name), &netappv1.VolumeTransaction{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeTransactions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(volumetransactionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &netappv1.VolumeTransactionList{})
	return err
}

// Patch applies the patch and returns the patched volumeTransaction.
func (c *FakeVolumeTransactions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *netappv1.VolumeTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumetransactionsResource, name, data, subresources...), &netappv1.VolumeTransaction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.VolumeTransaction), err
}
