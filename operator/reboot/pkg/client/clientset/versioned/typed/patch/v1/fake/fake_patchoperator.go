/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	patch_v1 "k8s-operator/operator/patch/pkg/apis/patch/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePatchOperators implements PatchOperatorInterface
type FakePatchOperators struct {
	Fake *FakePatchV1
	ns   string
}

var patchoperatorsResource = schema.GroupVersionResource{Group: "patch", Version: "v1", Resource: "patchoperators"}

var patchoperatorsKind = schema.GroupVersionKind{Group: "patch", Version: "v1", Kind: "PatchOperator"}

// Get takes name of the patchOperator, and returns the corresponding patchOperator object, and an error if there is any.
func (c *FakePatchOperators) Get(name string, options v1.GetOptions) (result *patch_v1.PatchOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(patchoperatorsResource, c.ns, name), &patch_v1.PatchOperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*patch_v1.PatchOperator), err
}

// List takes label and field selectors, and returns the list of PatchOperators that match those selectors.
func (c *FakePatchOperators) List(opts v1.ListOptions) (result *patch_v1.PatchOperatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(patchoperatorsResource, patchoperatorsKind, c.ns, opts), &patch_v1.PatchOperatorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &patch_v1.PatchOperatorList{}
	for _, item := range obj.(*patch_v1.PatchOperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested patchOperators.
func (c *FakePatchOperators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(patchoperatorsResource, c.ns, opts))

}

// Create takes the representation of a patchOperator and creates it.  Returns the server's representation of the patchOperator, and an error, if there is any.
func (c *FakePatchOperators) Create(patchOperator *patch_v1.PatchOperator) (result *patch_v1.PatchOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(patchoperatorsResource, c.ns, patchOperator), &patch_v1.PatchOperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*patch_v1.PatchOperator), err
}

// Update takes the representation of a patchOperator and updates it. Returns the server's representation of the patchOperator, and an error, if there is any.
func (c *FakePatchOperators) Update(patchOperator *patch_v1.PatchOperator) (result *patch_v1.PatchOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(patchoperatorsResource, c.ns, patchOperator), &patch_v1.PatchOperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*patch_v1.PatchOperator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePatchOperators) UpdateStatus(patchOperator *patch_v1.PatchOperator) (*patch_v1.PatchOperator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(patchoperatorsResource, "status", c.ns, patchOperator), &patch_v1.PatchOperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*patch_v1.PatchOperator), err
}

// Delete takes name of the patchOperator and deletes it. Returns an error if one occurs.
func (c *FakePatchOperators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(patchoperatorsResource, c.ns, name), &patch_v1.PatchOperator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePatchOperators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(patchoperatorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &patch_v1.PatchOperatorList{})
	return err
}

// Patch applies the patch and returns the patched patchOperator.
func (c *FakePatchOperators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *patch_v1.PatchOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(patchoperatorsResource, c.ns, name, data, subresources...), &patch_v1.PatchOperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*patch_v1.PatchOperator), err
}
