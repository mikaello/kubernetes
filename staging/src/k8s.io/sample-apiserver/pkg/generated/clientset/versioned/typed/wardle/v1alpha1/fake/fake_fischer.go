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
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1"
	wardlev1alpha1 "k8s.io/sample-apiserver/pkg/generated/applyconfiguration/wardle/v1alpha1"
)

// FakeFischers implements FischerInterface
type FakeFischers struct {
	Fake *FakeWardleV1alpha1
}

var fischersResource = v1alpha1.SchemeGroupVersion.WithResource("fischers")

var fischersKind = v1alpha1.SchemeGroupVersion.WithKind("Fischer")

// Get takes name of the fischer, and returns the corresponding fischer object, and an error if there is any.
func (c *FakeFischers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Fischer, err error) {
	emptyResult := &v1alpha1.Fischer{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(fischersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Fischer), err
}

// List takes label and field selectors, and returns the list of Fischers that match those selectors.
func (c *FakeFischers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FischerList, err error) {
	emptyResult := &v1alpha1.FischerList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(fischersResource, fischersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FischerList{ListMeta: obj.(*v1alpha1.FischerList).ListMeta}
	for _, item := range obj.(*v1alpha1.FischerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fischers.
func (c *FakeFischers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(fischersResource, opts))
}

// Create takes the representation of a fischer and creates it.  Returns the server's representation of the fischer, and an error, if there is any.
func (c *FakeFischers) Create(ctx context.Context, fischer *v1alpha1.Fischer, opts v1.CreateOptions) (result *v1alpha1.Fischer, err error) {
	emptyResult := &v1alpha1.Fischer{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(fischersResource, fischer, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Fischer), err
}

// Update takes the representation of a fischer and updates it. Returns the server's representation of the fischer, and an error, if there is any.
func (c *FakeFischers) Update(ctx context.Context, fischer *v1alpha1.Fischer, opts v1.UpdateOptions) (result *v1alpha1.Fischer, err error) {
	emptyResult := &v1alpha1.Fischer{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(fischersResource, fischer, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Fischer), err
}

// Delete takes name of the fischer and deletes it. Returns an error if one occurs.
func (c *FakeFischers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(fischersResource, name, opts), &v1alpha1.Fischer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFischers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(fischersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FischerList{})
	return err
}

// Patch applies the patch and returns the patched fischer.
func (c *FakeFischers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Fischer, err error) {
	emptyResult := &v1alpha1.Fischer{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(fischersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Fischer), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fischer.
func (c *FakeFischers) Apply(ctx context.Context, fischer *wardlev1alpha1.FischerApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Fischer, err error) {
	if fischer == nil {
		return nil, fmt.Errorf("fischer provided to Apply must not be nil")
	}
	data, err := json.Marshal(fischer)
	if err != nil {
		return nil, err
	}
	name := fischer.Name
	if name == nil {
		return nil, fmt.Errorf("fischer.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.Fischer{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(fischersResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Fischer), err
}
