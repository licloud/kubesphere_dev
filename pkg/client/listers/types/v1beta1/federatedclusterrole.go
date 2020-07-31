/*
Copyright 2020 The KubeSphere Authors.

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
	v1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
)

// FederatedClusterRoleLister helps list FederatedClusterRoles.
type FederatedClusterRoleLister interface {
	// List lists all FederatedClusterRoles in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRole, err error)
	// FederatedClusterRoles returns an object that can list and get FederatedClusterRoles.
	FederatedClusterRoles(namespace string) FederatedClusterRoleNamespaceLister
	FederatedClusterRoleListerExpansion
}

// federatedClusterRoleLister implements the FederatedClusterRoleLister interface.
type federatedClusterRoleLister struct {
	indexer cache.Indexer
}

// NewFederatedClusterRoleLister returns a new FederatedClusterRoleLister.
func NewFederatedClusterRoleLister(indexer cache.Indexer) FederatedClusterRoleLister {
	return &federatedClusterRoleLister{indexer: indexer}
}

// List lists all FederatedClusterRoles in the indexer.
func (s *federatedClusterRoleLister) List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedClusterRole))
	})
	return ret, err
}

// FederatedClusterRoles returns an object that can list and get FederatedClusterRoles.
func (s *federatedClusterRoleLister) FederatedClusterRoles(namespace string) FederatedClusterRoleNamespaceLister {
	return federatedClusterRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedClusterRoleNamespaceLister helps list and get FederatedClusterRoles.
type FederatedClusterRoleNamespaceLister interface {
	// List lists all FederatedClusterRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRole, err error)
	// Get retrieves the FederatedClusterRole from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedClusterRole, error)
	FederatedClusterRoleNamespaceListerExpansion
}

// federatedClusterRoleNamespaceLister implements the FederatedClusterRoleNamespaceLister
// interface.
type federatedClusterRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedClusterRoles in the indexer for a given namespace.
func (s federatedClusterRoleNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedClusterRole))
	})
	return ret, err
}

// Get retrieves the FederatedClusterRole from the indexer for a given namespace and name.
func (s federatedClusterRoleNamespaceLister) Get(name string) (*v1beta1.FederatedClusterRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedclusterrole"), name)
	}
	return obj.(*v1beta1.FederatedClusterRole), nil
}
