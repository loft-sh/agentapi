// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/cluster/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ChartInfoLister helps list ChartInfos.
// All objects returned here must be treated as read-only.
type ChartInfoLister interface {
	// List lists all ChartInfos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ChartInfo, err error)
	// Get retrieves the ChartInfo from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ChartInfo, error)
	ChartInfoListerExpansion
}

// chartInfoLister implements the ChartInfoLister interface.
type chartInfoLister struct {
	indexer cache.Indexer
}

// NewChartInfoLister returns a new ChartInfoLister.
func NewChartInfoLister(indexer cache.Indexer) ChartInfoLister {
	return &chartInfoLister{indexer: indexer}
}

// List lists all ChartInfos in the indexer.
func (s *chartInfoLister) List(selector labels.Selector) (ret []*v1.ChartInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ChartInfo))
	})
	return ret, err
}

// Get retrieves the ChartInfo from the index for a given name.
func (s *chartInfoLister) Get(name string) (*v1.ChartInfo, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("chartinfo"), name)
	}
	return obj.(*v1.ChartInfo), nil
}
