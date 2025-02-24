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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/containerattached/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContainerAttachedClusters implements ContainerAttachedClusterInterface
type FakeContainerAttachedClusters struct {
	Fake *FakeContainerattachedV1beta1
	ns   string
}

var containerattachedclustersResource = schema.GroupVersionResource{Group: "containerattached.cnrm.cloud.google.com", Version: "v1beta1", Resource: "containerattachedclusters"}

var containerattachedclustersKind = schema.GroupVersionKind{Group: "containerattached.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ContainerAttachedCluster"}

// Get takes name of the containerAttachedCluster, and returns the corresponding containerAttachedCluster object, and an error if there is any.
func (c *FakeContainerAttachedClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ContainerAttachedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(containerattachedclustersResource, c.ns, name), &v1beta1.ContainerAttachedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerAttachedCluster), err
}

// List takes label and field selectors, and returns the list of ContainerAttachedClusters that match those selectors.
func (c *FakeContainerAttachedClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ContainerAttachedClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(containerattachedclustersResource, containerattachedclustersKind, c.ns, opts), &v1beta1.ContainerAttachedClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ContainerAttachedClusterList{ListMeta: obj.(*v1beta1.ContainerAttachedClusterList).ListMeta}
	for _, item := range obj.(*v1beta1.ContainerAttachedClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerAttachedClusters.
func (c *FakeContainerAttachedClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(containerattachedclustersResource, c.ns, opts))

}

// Create takes the representation of a containerAttachedCluster and creates it.  Returns the server's representation of the containerAttachedCluster, and an error, if there is any.
func (c *FakeContainerAttachedClusters) Create(ctx context.Context, containerAttachedCluster *v1beta1.ContainerAttachedCluster, opts v1.CreateOptions) (result *v1beta1.ContainerAttachedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(containerattachedclustersResource, c.ns, containerAttachedCluster), &v1beta1.ContainerAttachedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerAttachedCluster), err
}

// Update takes the representation of a containerAttachedCluster and updates it. Returns the server's representation of the containerAttachedCluster, and an error, if there is any.
func (c *FakeContainerAttachedClusters) Update(ctx context.Context, containerAttachedCluster *v1beta1.ContainerAttachedCluster, opts v1.UpdateOptions) (result *v1beta1.ContainerAttachedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(containerattachedclustersResource, c.ns, containerAttachedCluster), &v1beta1.ContainerAttachedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerAttachedCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerAttachedClusters) UpdateStatus(ctx context.Context, containerAttachedCluster *v1beta1.ContainerAttachedCluster, opts v1.UpdateOptions) (*v1beta1.ContainerAttachedCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(containerattachedclustersResource, "status", c.ns, containerAttachedCluster), &v1beta1.ContainerAttachedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerAttachedCluster), err
}

// Delete takes name of the containerAttachedCluster and deletes it. Returns an error if one occurs.
func (c *FakeContainerAttachedClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(containerattachedclustersResource, c.ns, name, opts), &v1beta1.ContainerAttachedCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerAttachedClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(containerattachedclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ContainerAttachedClusterList{})
	return err
}

// Patch applies the patch and returns the patched containerAttachedCluster.
func (c *FakeContainerAttachedClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ContainerAttachedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(containerattachedclustersResource, c.ns, name, pt, data, subresources...), &v1beta1.ContainerAttachedCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerAttachedCluster), err
}
