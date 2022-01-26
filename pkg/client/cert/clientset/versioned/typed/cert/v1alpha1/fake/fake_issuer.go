/*
Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gardener/cert-management/pkg/apis/cert/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIssuers implements IssuerInterface
type FakeIssuers struct {
	Fake *FakeCertV1alpha1
	ns   string
}

var issuersResource = schema.GroupVersionResource{Group: "cert.gardener.cloud", Version: "v1alpha1", Resource: "issuers"}

var issuersKind = schema.GroupVersionKind{Group: "cert.gardener.cloud", Version: "v1alpha1", Kind: "Issuer"}

// Get takes name of the issuer, and returns the corresponding issuer object, and an error if there is any.
func (c *FakeIssuers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Issuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(issuersResource, c.ns, name), &v1alpha1.Issuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Issuer), err
}

// List takes label and field selectors, and returns the list of Issuers that match those selectors.
func (c *FakeIssuers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IssuerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(issuersResource, issuersKind, c.ns, opts), &v1alpha1.IssuerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IssuerList{ListMeta: obj.(*v1alpha1.IssuerList).ListMeta}
	for _, item := range obj.(*v1alpha1.IssuerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested issuers.
func (c *FakeIssuers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(issuersResource, c.ns, opts))

}

// Create takes the representation of a issuer and creates it.  Returns the server's representation of the issuer, and an error, if there is any.
func (c *FakeIssuers) Create(ctx context.Context, issuer *v1alpha1.Issuer, opts v1.CreateOptions) (result *v1alpha1.Issuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(issuersResource, c.ns, issuer), &v1alpha1.Issuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Issuer), err
}

// Update takes the representation of a issuer and updates it. Returns the server's representation of the issuer, and an error, if there is any.
func (c *FakeIssuers) Update(ctx context.Context, issuer *v1alpha1.Issuer, opts v1.UpdateOptions) (result *v1alpha1.Issuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(issuersResource, c.ns, issuer), &v1alpha1.Issuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Issuer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIssuers) UpdateStatus(ctx context.Context, issuer *v1alpha1.Issuer, opts v1.UpdateOptions) (*v1alpha1.Issuer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(issuersResource, "status", c.ns, issuer), &v1alpha1.Issuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Issuer), err
}

// Delete takes name of the issuer and deletes it. Returns an error if one occurs.
func (c *FakeIssuers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(issuersResource, c.ns, name), &v1alpha1.Issuer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIssuers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(issuersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IssuerList{})
	return err
}

// Patch applies the patch and returns the patched issuer.
func (c *FakeIssuers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Issuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(issuersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Issuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Issuer), err
}
