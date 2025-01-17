/*
Copyright 2022 The Knative Authors

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-rabbitmq/pkg/apis/eventing/v1alpha1"
)

// FakeRabbitmqBrokerConfigs implements RabbitmqBrokerConfigInterface
type FakeRabbitmqBrokerConfigs struct {
	Fake *FakeEventingV1alpha1
	ns   string
}

var rabbitmqbrokerconfigsResource = v1alpha1.SchemeGroupVersion.WithResource("rabbitmqbrokerconfigs")

var rabbitmqbrokerconfigsKind = v1alpha1.SchemeGroupVersion.WithKind("RabbitmqBrokerConfig")

// Get takes name of the rabbitmqBrokerConfig, and returns the corresponding rabbitmqBrokerConfig object, and an error if there is any.
func (c *FakeRabbitmqBrokerConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RabbitmqBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rabbitmqbrokerconfigsResource, c.ns, name), &v1alpha1.RabbitmqBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqBrokerConfig), err
}

// List takes label and field selectors, and returns the list of RabbitmqBrokerConfigs that match those selectors.
func (c *FakeRabbitmqBrokerConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RabbitmqBrokerConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rabbitmqbrokerconfigsResource, rabbitmqbrokerconfigsKind, c.ns, opts), &v1alpha1.RabbitmqBrokerConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RabbitmqBrokerConfigList{ListMeta: obj.(*v1alpha1.RabbitmqBrokerConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.RabbitmqBrokerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rabbitmqBrokerConfigs.
func (c *FakeRabbitmqBrokerConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rabbitmqbrokerconfigsResource, c.ns, opts))

}

// Create takes the representation of a rabbitmqBrokerConfig and creates it.  Returns the server's representation of the rabbitmqBrokerConfig, and an error, if there is any.
func (c *FakeRabbitmqBrokerConfigs) Create(ctx context.Context, rabbitmqBrokerConfig *v1alpha1.RabbitmqBrokerConfig, opts v1.CreateOptions) (result *v1alpha1.RabbitmqBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rabbitmqbrokerconfigsResource, c.ns, rabbitmqBrokerConfig), &v1alpha1.RabbitmqBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqBrokerConfig), err
}

// Update takes the representation of a rabbitmqBrokerConfig and updates it. Returns the server's representation of the rabbitmqBrokerConfig, and an error, if there is any.
func (c *FakeRabbitmqBrokerConfigs) Update(ctx context.Context, rabbitmqBrokerConfig *v1alpha1.RabbitmqBrokerConfig, opts v1.UpdateOptions) (result *v1alpha1.RabbitmqBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rabbitmqbrokerconfigsResource, c.ns, rabbitmqBrokerConfig), &v1alpha1.RabbitmqBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqBrokerConfig), err
}

// Delete takes name of the rabbitmqBrokerConfig and deletes it. Returns an error if one occurs.
func (c *FakeRabbitmqBrokerConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rabbitmqbrokerconfigsResource, c.ns, name, opts), &v1alpha1.RabbitmqBrokerConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRabbitmqBrokerConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rabbitmqbrokerconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RabbitmqBrokerConfigList{})
	return err
}

// Patch applies the patch and returns the patched rabbitmqBrokerConfig.
func (c *FakeRabbitmqBrokerConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RabbitmqBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rabbitmqbrokerconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RabbitmqBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RabbitmqBrokerConfig), err
}
