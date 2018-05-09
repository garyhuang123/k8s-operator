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

package v1

import (
	v1 "k8s-operator/operator/patch/pkg/apis/patch/v1"
	scheme "k8s-operator/operator/patch/pkg/client/clientset/versioned/scheme"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PatchOperatorsGetter has a method to return a PatchOperatorInterface.
// A group's client should implement this interface.
type PatchOperatorsGetter interface {
	PatchOperators(namespace string) PatchOperatorInterface
}

// PatchOperatorInterface has methods to work with PatchOperator resources.
type PatchOperatorInterface interface {
	Create(*v1.PatchOperator) (*v1.PatchOperator, error)
	Update(*v1.PatchOperator) (*v1.PatchOperator, error)
	UpdateStatus(*v1.PatchOperator) (*v1.PatchOperator, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.PatchOperator, error)
	List(opts meta_v1.ListOptions) (*v1.PatchOperatorList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PatchOperator, err error)
	PatchOperatorExpansion
}

// patchOperators implements PatchOperatorInterface
type patchOperators struct {
	client rest.Interface
	ns     string
}

// newPatchOperators returns a PatchOperators
func newPatchOperators(c *PatchV1Client, namespace string) *patchOperators {
	return &patchOperators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the patchOperator, and returns the corresponding patchOperator object, and an error if there is any.
func (c *patchOperators) Get(name string, options meta_v1.GetOptions) (result *v1.PatchOperator, err error) {
	result = &v1.PatchOperator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("patchoperators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PatchOperators that match those selectors.
func (c *patchOperators) List(opts meta_v1.ListOptions) (result *v1.PatchOperatorList, err error) {
	result = &v1.PatchOperatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("patchoperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested patchOperators.
func (c *patchOperators) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("patchoperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a patchOperator and creates it.  Returns the server's representation of the patchOperator, and an error, if there is any.
func (c *patchOperators) Create(patchOperator *v1.PatchOperator) (result *v1.PatchOperator, err error) {
	result = &v1.PatchOperator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("patchoperators").
		Body(patchOperator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a patchOperator and updates it. Returns the server's representation of the patchOperator, and an error, if there is any.
func (c *patchOperators) Update(patchOperator *v1.PatchOperator) (result *v1.PatchOperator, err error) {
	result = &v1.PatchOperator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("patchoperators").
		Name(patchOperator.Name).
		Body(patchOperator).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *patchOperators) UpdateStatus(patchOperator *v1.PatchOperator) (result *v1.PatchOperator, err error) {
	result = &v1.PatchOperator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("patchoperators").
		Name(patchOperator.Name).
		SubResource("status").
		Body(patchOperator).
		Do().
		Into(result)
	return
}

// Delete takes name of the patchOperator and deletes it. Returns an error if one occurs.
func (c *patchOperators) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("patchoperators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *patchOperators) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("patchoperators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched patchOperator.
func (c *patchOperators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PatchOperator, err error) {
	result = &v1.PatchOperator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("patchoperators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
