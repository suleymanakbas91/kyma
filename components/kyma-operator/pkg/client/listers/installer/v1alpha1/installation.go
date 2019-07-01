// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/kyma-operator/pkg/apis/installer/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstallationLister helps list Installations.
type InstallationLister interface {
	// List lists all Installations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Installation, err error)
	// Installations returns an object that can list and get Installations.
	Installations(namespace string) InstallationNamespaceLister
	InstallationListerExpansion
}

// installationLister implements the InstallationLister interface.
type installationLister struct {
	indexer cache.Indexer
}

// NewInstallationLister returns a new InstallationLister.
func NewInstallationLister(indexer cache.Indexer) InstallationLister {
	return &installationLister{indexer: indexer}
}

// List lists all Installations in the indexer.
func (s *installationLister) List(selector labels.Selector) (ret []*v1alpha1.Installation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Installation))
	})
	return ret, err
}

// Installations returns an object that can list and get Installations.
func (s *installationLister) Installations(namespace string) InstallationNamespaceLister {
	return installationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstallationNamespaceLister helps list and get Installations.
type InstallationNamespaceLister interface {
	// List lists all Installations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Installation, err error)
	// Get retrieves the Installation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Installation, error)
	InstallationNamespaceListerExpansion
}

// installationNamespaceLister implements the InstallationNamespaceLister
// interface.
type installationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Installations in the indexer for a given namespace.
func (s installationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Installation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Installation))
	})
	return ret, err
}

// Get retrieves the Installation from the indexer for a given namespace and name.
func (s installationNamespaceLister) Get(name string) (*v1alpha1.Installation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("installation"), name)
	}
	return obj.(*v1alpha1.Installation), nil
}
