/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

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

	v1beta2 "github.com/apache/camel-k/addons/strimzi/duck/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKafkas implements KafkaInterface
type FakeKafkas struct {
	Fake *FakeKafkaV1beta2
	ns   string
}

var kafkasResource = schema.GroupVersionResource{Group: "kafka.strimzi.io", Version: "v1beta2", Resource: "kafkas"}

var kafkasKind = schema.GroupVersionKind{Group: "kafka.strimzi.io", Version: "v1beta2", Kind: "Kafka"}

// Get takes name of the kafka, and returns the corresponding kafka object, and an error if there is any.
func (c *FakeKafkas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.Kafka, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kafkasResource, c.ns, name), &v1beta2.Kafka{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.Kafka), err
}

// List takes label and field selectors, and returns the list of Kafkas that match those selectors.
func (c *FakeKafkas) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.KafkaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kafkasResource, kafkasKind, c.ns, opts), &v1beta2.KafkaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.KafkaList{ListMeta: obj.(*v1beta2.KafkaList).ListMeta}
	for _, item := range obj.(*v1beta2.KafkaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kafkas.
func (c *FakeKafkas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kafkasResource, c.ns, opts))

}