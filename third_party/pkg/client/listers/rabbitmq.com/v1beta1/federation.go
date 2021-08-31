/*
Copyright 2020 The Knative Authors

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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/eventing-rabbitmq/third_party/pkg/apis/rabbitmq.com/v1beta1"
)

// FederationLister helps list Federations.
// All objects returned here must be treated as read-only.
type FederationLister interface {
	// List lists all Federations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Federation, err error)
	// Federations returns an object that can list and get Federations.
	Federations(namespace string) FederationNamespaceLister
	FederationListerExpansion
}

// federationLister implements the FederationLister interface.
type federationLister struct {
	indexer cache.Indexer
}

// NewFederationLister returns a new FederationLister.
func NewFederationLister(indexer cache.Indexer) FederationLister {
	return &federationLister{indexer: indexer}
}

// List lists all Federations in the indexer.
func (s *federationLister) List(selector labels.Selector) (ret []*v1beta1.Federation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Federation))
	})
	return ret, err
}

// Federations returns an object that can list and get Federations.
func (s *federationLister) Federations(namespace string) FederationNamespaceLister {
	return federationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederationNamespaceLister helps list and get Federations.
// All objects returned here must be treated as read-only.
type FederationNamespaceLister interface {
	// List lists all Federations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Federation, err error)
	// Get retrieves the Federation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Federation, error)
	FederationNamespaceListerExpansion
}

// federationNamespaceLister implements the FederationNamespaceLister
// interface.
type federationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Federations in the indexer for a given namespace.
func (s federationNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Federation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Federation))
	})
	return ret, err
}

// Get retrieves the Federation from the indexer for a given namespace and name.
func (s federationNamespaceLister) Get(name string) (*v1beta1.Federation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federation"), name)
	}
	return obj.(*v1beta1.Federation), nil
}
