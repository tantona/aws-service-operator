/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "github.com/awslabs/aws-service-operator/pkg/apis/service-operator.aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSNSSubscriptions implements SNSSubscriptionInterface
type FakeSNSSubscriptions struct {
	Fake *FakeServiceoperatorV1alpha1
	ns   string
}

var snssubscriptionsResource = schema.GroupVersionResource{Group: "serviceoperator.aws", Version: "v1alpha1", Resource: "snssubscriptions"}

var snssubscriptionsKind = schema.GroupVersionKind{Group: "serviceoperator.aws", Version: "v1alpha1", Kind: "SNSSubscription"}

// Get takes name of the sNSSubscription, and returns the corresponding sNSSubscription object, and an error if there is any.
func (c *FakeSNSSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.SNSSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snssubscriptionsResource, c.ns, name), &v1alpha1.SNSSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSSubscription), err
}

// List takes label and field selectors, and returns the list of SNSSubscriptions that match those selectors.
func (c *FakeSNSSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.SNSSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snssubscriptionsResource, snssubscriptionsKind, c.ns, opts), &v1alpha1.SNSSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SNSSubscriptionList{}
	for _, item := range obj.(*v1alpha1.SNSSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sNSSubscriptions.
func (c *FakeSNSSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snssubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a sNSSubscription and creates it.  Returns the server's representation of the sNSSubscription, and an error, if there is any.
func (c *FakeSNSSubscriptions) Create(sNSSubscription *v1alpha1.SNSSubscription) (result *v1alpha1.SNSSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snssubscriptionsResource, c.ns, sNSSubscription), &v1alpha1.SNSSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSSubscription), err
}

// Update takes the representation of a sNSSubscription and updates it. Returns the server's representation of the sNSSubscription, and an error, if there is any.
func (c *FakeSNSSubscriptions) Update(sNSSubscription *v1alpha1.SNSSubscription) (result *v1alpha1.SNSSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snssubscriptionsResource, c.ns, sNSSubscription), &v1alpha1.SNSSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSSubscription), err
}

// Delete takes name of the sNSSubscription and deletes it. Returns an error if one occurs.
func (c *FakeSNSSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snssubscriptionsResource, c.ns, name), &v1alpha1.SNSSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSNSSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snssubscriptionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SNSSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched sNSSubscription.
func (c *FakeSNSSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SNSSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snssubscriptionsResource, c.ns, name, data, subresources...), &v1alpha1.SNSSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSSubscription), err
}
