// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIAMServiceAccounts implements IAMServiceAccountInterface
type FakeIAMServiceAccounts struct {
	Fake *FakeIamV1beta1
	ns   string
}

var iamserviceaccountsResource = schema.GroupVersionResource{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iamserviceaccounts"}

var iamserviceaccountsKind = schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMServiceAccount"}

// Get takes name of the iAMServiceAccount, and returns the corresponding iAMServiceAccount object, and an error if there is any.
func (c *FakeIAMServiceAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamserviceaccountsResource, c.ns, name), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}

// List takes label and field selectors, and returns the list of IAMServiceAccounts that match those selectors.
func (c *FakeIAMServiceAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IAMServiceAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamserviceaccountsResource, iamserviceaccountsKind, c.ns, opts), &v1beta1.IAMServiceAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IAMServiceAccountList{ListMeta: obj.(*v1beta1.IAMServiceAccountList).ListMeta}
	for _, item := range obj.(*v1beta1.IAMServiceAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iAMServiceAccounts.
func (c *FakeIAMServiceAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamserviceaccountsResource, c.ns, opts))

}

// Create takes the representation of a iAMServiceAccount and creates it.  Returns the server's representation of the iAMServiceAccount, and an error, if there is any.
func (c *FakeIAMServiceAccounts) Create(ctx context.Context, iAMServiceAccount *v1beta1.IAMServiceAccount, opts v1.CreateOptions) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamserviceaccountsResource, c.ns, iAMServiceAccount), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}

// Update takes the representation of a iAMServiceAccount and updates it. Returns the server's representation of the iAMServiceAccount, and an error, if there is any.
func (c *FakeIAMServiceAccounts) Update(ctx context.Context, iAMServiceAccount *v1beta1.IAMServiceAccount, opts v1.UpdateOptions) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamserviceaccountsResource, c.ns, iAMServiceAccount), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIAMServiceAccounts) UpdateStatus(ctx context.Context, iAMServiceAccount *v1beta1.IAMServiceAccount, opts v1.UpdateOptions) (*v1beta1.IAMServiceAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamserviceaccountsResource, "status", c.ns, iAMServiceAccount), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}

// Delete takes name of the iAMServiceAccount and deletes it. Returns an error if one occurs.
func (c *FakeIAMServiceAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(iamserviceaccountsResource, c.ns, name, opts), &v1beta1.IAMServiceAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIAMServiceAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamserviceaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IAMServiceAccountList{})
	return err
}

// Patch applies the patch and returns the patched iAMServiceAccount.
func (c *FakeIAMServiceAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IAMServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamserviceaccountsResource, c.ns, name, pt, data, subresources...), &v1beta1.IAMServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMServiceAccount), err
}
