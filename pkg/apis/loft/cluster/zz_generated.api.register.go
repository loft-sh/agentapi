// Code generated by generator. DO NOT EDIT.

package cluster

import (
	"context"
	"fmt"

	storagev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	"github.com/loft-sh/apiserver/pkg/builders"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	configrest "k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NewRESTFunc func(config *configrest.Config, cachedClient client.Client, uncachedClient client.Client) rest.Storage

var (
	ClusterChartInfoStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalChartInfo,
		func() runtime.Object { return &ChartInfo{} },     // Register versioned resource
		func() runtime.Object { return &ChartInfoList{} }, // Register versioned resource list
		NewChartInfoREST,
	)
	NewChartInfoREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewChartInfoRESTFunc(Config, CachedClient, UncachedClient)
	}
	NewChartInfoRESTFunc       NewRESTFunc
	ClusterClusterQuotaStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalClusterQuota,
		func() runtime.Object { return &ClusterQuota{} },     // Register versioned resource
		func() runtime.Object { return &ClusterQuotaList{} }, // Register versioned resource list
		NewClusterQuotaREST,
	)
	NewClusterQuotaREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewClusterQuotaRESTFunc(Config, CachedClient, UncachedClient)
	}
	NewClusterQuotaRESTFunc   NewRESTFunc
	ClusterHelmReleaseStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalHelmRelease,
		func() runtime.Object { return &HelmRelease{} },     // Register versioned resource
		func() runtime.Object { return &HelmReleaseList{} }, // Register versioned resource list
		NewHelmReleaseREST,
	)
	NewHelmReleaseREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewHelmReleaseRESTFunc(Config, CachedClient, UncachedClient)
	}
	NewHelmReleaseRESTFunc           NewRESTFunc
	ClusterLocalClusterAccessStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalLocalClusterAccess,
		func() runtime.Object { return &LocalClusterAccess{} },     // Register versioned resource
		func() runtime.Object { return &LocalClusterAccessList{} }, // Register versioned resource list
		NewLocalClusterAccessREST,
	)
	NewLocalClusterAccessREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewLocalClusterAccessRESTFunc(Config, CachedClient, UncachedClient)
	}
	NewLocalClusterAccessRESTFunc NewRESTFunc
	ClusterSleepModeConfigStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalSleepModeConfig,
		func() runtime.Object { return &SleepModeConfig{} },     // Register versioned resource
		func() runtime.Object { return &SleepModeConfigList{} }, // Register versioned resource list
		NewSleepModeConfigREST,
	)
	NewSleepModeConfigREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewSleepModeConfigRESTFunc(Config, CachedClient, UncachedClient)
	}
	NewSleepModeConfigRESTFunc NewRESTFunc
	ClusterSpaceStorage        = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalSpace,
		func() runtime.Object { return &Space{} },     // Register versioned resource
		func() runtime.Object { return &SpaceList{} }, // Register versioned resource list
		NewSpaceREST,
	)
	NewSpaceREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewSpaceRESTFunc(Config, CachedClient, UncachedClient)
	}
	NewSpaceRESTFunc             NewRESTFunc
	ClusterVirtualClusterStorage = builders.NewApiResourceWithStorage( // Resource status endpoint
		InternalVirtualCluster,
		func() runtime.Object { return &VirtualCluster{} },     // Register versioned resource
		func() runtime.Object { return &VirtualClusterList{} }, // Register versioned resource list
		NewVirtualClusterREST,
	)
	NewVirtualClusterREST = func(getter generic.RESTOptionsGetter) rest.Storage {
		return NewVirtualClusterRESTFunc(Config, CachedClient, UncachedClient)
	}
	NewVirtualClusterRESTFunc NewRESTFunc
	InternalChartInfo         = builders.NewInternalResource(
		"chartinfos",
		"ChartInfo",
		func() runtime.Object { return &ChartInfo{} },
		func() runtime.Object { return &ChartInfoList{} },
	)
	InternalChartInfoStatus = builders.NewInternalResourceStatus(
		"chartinfos",
		"ChartInfoStatus",
		func() runtime.Object { return &ChartInfo{} },
		func() runtime.Object { return &ChartInfoList{} },
	)
	InternalClusterQuota = builders.NewInternalResource(
		"clusterquotas",
		"ClusterQuota",
		func() runtime.Object { return &ClusterQuota{} },
		func() runtime.Object { return &ClusterQuotaList{} },
	)
	InternalClusterQuotaStatus = builders.NewInternalResourceStatus(
		"clusterquotas",
		"ClusterQuotaStatus",
		func() runtime.Object { return &ClusterQuota{} },
		func() runtime.Object { return &ClusterQuotaList{} },
	)
	InternalHelmRelease = builders.NewInternalResource(
		"helmreleases",
		"HelmRelease",
		func() runtime.Object { return &HelmRelease{} },
		func() runtime.Object { return &HelmReleaseList{} },
	)
	InternalHelmReleaseStatus = builders.NewInternalResourceStatus(
		"helmreleases",
		"HelmReleaseStatus",
		func() runtime.Object { return &HelmRelease{} },
		func() runtime.Object { return &HelmReleaseList{} },
	)
	InternalLocalClusterAccess = builders.NewInternalResource(
		"localclusteraccesses",
		"LocalClusterAccess",
		func() runtime.Object { return &LocalClusterAccess{} },
		func() runtime.Object { return &LocalClusterAccessList{} },
	)
	InternalLocalClusterAccessStatus = builders.NewInternalResourceStatus(
		"localclusteraccesses",
		"LocalClusterAccessStatus",
		func() runtime.Object { return &LocalClusterAccess{} },
		func() runtime.Object { return &LocalClusterAccessList{} },
	)
	InternalSleepModeConfig = builders.NewInternalResource(
		"sleepmodeconfigs",
		"SleepModeConfig",
		func() runtime.Object { return &SleepModeConfig{} },
		func() runtime.Object { return &SleepModeConfigList{} },
	)
	InternalSleepModeConfigStatus = builders.NewInternalResourceStatus(
		"sleepmodeconfigs",
		"SleepModeConfigStatus",
		func() runtime.Object { return &SleepModeConfig{} },
		func() runtime.Object { return &SleepModeConfigList{} },
	)
	InternalSpace = builders.NewInternalResource(
		"spaces",
		"Space",
		func() runtime.Object { return &Space{} },
		func() runtime.Object { return &SpaceList{} },
	)
	InternalSpaceStatus = builders.NewInternalResourceStatus(
		"spaces",
		"SpaceStatus",
		func() runtime.Object { return &Space{} },
		func() runtime.Object { return &SpaceList{} },
	)
	InternalVirtualCluster = builders.NewInternalResource(
		"virtualclusters",
		"VirtualCluster",
		func() runtime.Object { return &VirtualCluster{} },
		func() runtime.Object { return &VirtualClusterList{} },
	)
	InternalVirtualClusterStatus = builders.NewInternalResourceStatus(
		"virtualclusters",
		"VirtualClusterStatus",
		func() runtime.Object { return &VirtualCluster{} },
		func() runtime.Object { return &VirtualClusterList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("cluster.loft.sh").WithKinds(
		InternalChartInfo,
		InternalChartInfoStatus,
		InternalClusterQuota,
		InternalClusterQuotaStatus,
		InternalHelmRelease,
		InternalHelmReleaseStatus,
		InternalLocalClusterAccess,
		InternalLocalClusterAccessStatus,
		InternalSleepModeConfig,
		InternalSleepModeConfigStatus,
		InternalSpace,
		InternalSpaceStatus,
		InternalVirtualCluster,
		InternalVirtualClusterStatus,
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
		RegisterDefaults,
	}).AddToScheme
	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

type FinalizerName string
type NamespacePhase string
type Status string

type AppliedMetadata struct {
	Annotations map[string]string
	Labels      map[string]string
}

type AppliedObject struct {
	APIVersion string
	Kind       string
	Name       string
}

type Bash struct {
	Script      string
	Image       string
	ClusterRole string
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartInfo struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ChartInfoSpec
	Status ChartInfoStatus
}

type ChartInfoSpec struct {
	Chart storagev1.Chart
}

type ChartInfoStatus struct {
	Metadata *Metadata
	Readme   string
	Values   string
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterQuota struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ClusterQuotaSpec
	Status ClusterQuotaStatus
}

type ClusterQuotaSpec struct {
	storagev1.ClusterQuotaSpec
}

type ClusterQuotaStatus struct {
	storagev1.ClusterQuotaStatus
	Owner *UserOrTeam
}

type EntityInfo struct {
	Name        string
	DisplayName string
	Icon        string
	Username    string
	Email       string
	Subject     string
}

type EpochInfo struct {
	Start int64
	Slept int64
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmRelease struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   HelmReleaseSpec
	Status HelmReleaseStatus
}

type HelmReleaseConfig struct {
	Chart       storagev1.Chart
	Manifests   string
	Bash        *Bash
	Values      string
	Parameters  string
	Annotations map[string]string
}

type HelmReleaseSpec struct {
	HelmReleaseConfig
}

type HelmReleaseStatus struct {
	Revision int
	Info     *Info
	Metadata *Metadata
}

type Info struct {
	FirstDeployed metav1.Time
	LastDeployed  metav1.Time
	Deleted       metav1.Time
	Description   string
	Status        Status
	Notes         string
}

type LastActivityInfo struct {
	Subject        string
	Host           string
	Verb           string
	APIGroup       string
	Resource       string
	Subresource    string
	Name           string
	VirtualCluster string
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LocalClusterAccess struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   LocalClusterAccessSpec
	Status LocalClusterAccessStatus
}

type LocalClusterAccessSpec struct {
	storagev1.LocalClusterAccessSpec
}

type LocalClusterAccessStatus struct {
	storagev1.LocalClusterAccessStatus
	Users           []*UserOrTeam
	Teams           []*EntityInfo
	SpaceConstraint *EntityInfo
}

type Maintainer struct {
	Name  string
	Email string
	URL   string
}

type Metadata struct {
	Name        string
	Home        string
	Sources     []string
	Version     string
	Description string
	Keywords    []string
	Maintainers []*Maintainer
	Icon        string
	APIVersion  string
	Condition   string
	Tags        string
	AppVersion  string
	Deprecated  bool
	Annotations map[string]string
	KubeVersion string
	Type        string
	Urls        []string
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SleepModeConfig struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   SleepModeConfigSpec
	Status SleepModeConfigStatus
}

type SleepModeConfigSpec struct {
	ForceSleep              bool
	ForceSleepDuration      *int64
	DeleteAllPods           bool
	DeleteAfter             int64
	SleepAfter              int64
	SleepSchedule           string
	WakeupSchedule          string
	Timezone                string
	IgnoreActiveConnections bool
	IgnoreAll               bool
	IgnoreVClusters         bool
	IgnoreGroups            string
	IgnoreVerbs             string
	IgnoreResources         string
	IgnoreResourceVerbs     string
	IgnoreResourceNames     string
}

type SleepModeConfigStatus struct {
	LastActivity        int64
	LastActivityInfo    *LastActivityInfo
	SleepingSince       int64
	CurrentEpoch        *EpochInfo
	LastEpoch           *EpochInfo
	SleptLastThirtyDays *float64
	SleptLastSevenDays  *float64
	ScheduledSleep      *int64
	ScheduledWakeup     *int64
	SleepType           string
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Space struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   SpaceSpec
	Status SpaceStatus
}

type SpaceConstraintNamespaceStatus struct {
	SpaceConstraint    string
	User               string
	Team               string
	ObservedGeneration int64
	Phase              string
	Reason             string
	Message            string
	AppliedClusterRole *string
	AppliedMetadata    AppliedMetadata
	AppliedObjects     []AppliedObject
}

type SpaceObjectsNamespaceStatus struct {
	Phase          string
	Reason         string
	Message        string
	AppliedObjects []AppliedObject
}

type SpaceSpec struct {
	User       string
	Team       string
	Objects    string
	Finalizers []corev1.FinalizerName
}

type SpaceStatus struct {
	Phase                 corev1.NamespacePhase
	SleepModeConfig       *SleepModeConfig
	Owner                 *UserOrTeam
	SpaceConstraint       *EntityInfo
	SpaceConstraintStatus *SpaceConstraintNamespaceStatus
	SpaceObjectsStatus    *SpaceObjectsNamespaceStatus
	TemplateSyncStatus    *TemplateSyncStatus
}

type TemplateSyncStatus struct {
	Template string
	Phase    string
}

type UserOrTeam struct {
	User *EntityInfo
	Team *EntityInfo
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualCluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   VirtualClusterSpec
	Status VirtualClusterStatus
}

type VirtualClusterSpec struct {
	storagev1.VirtualClusterSpec
}

type VirtualClusterStatus struct {
	storagev1.VirtualClusterStatus
	SyncerPod          *corev1.Pod
	ClusterPod         *corev1.Pod
	SleepModeConfig    *SleepModeConfig
	TemplateSyncStatus *TemplateSyncStatus
}

//
// ChartInfo Functions and Structs
//
// +k8s:deepcopy-gen=false
type ChartInfoStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type ChartInfoStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartInfoList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ChartInfo
}

func (ChartInfo) NewStatus() interface{} {
	return ChartInfoStatus{}
}

func (pc *ChartInfo) GetStatus() interface{} {
	return pc.Status
}

func (pc *ChartInfo) SetStatus(s interface{}) {
	pc.Status = s.(ChartInfoStatus)
}

func (pc *ChartInfo) GetSpec() interface{} {
	return pc.Spec
}

func (pc *ChartInfo) SetSpec(s interface{}) {
	pc.Spec = s.(ChartInfoSpec)
}

func (pc *ChartInfo) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *ChartInfo) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc ChartInfo) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store ChartInfo.
// +k8s:deepcopy-gen=false
type ChartInfoRegistry interface {
	ListChartInfos(ctx context.Context, options *internalversion.ListOptions) (*ChartInfoList, error)
	GetChartInfo(ctx context.Context, id string, options *metav1.GetOptions) (*ChartInfo, error)
	CreateChartInfo(ctx context.Context, id *ChartInfo) (*ChartInfo, error)
	UpdateChartInfo(ctx context.Context, id *ChartInfo) (*ChartInfo, error)
	DeleteChartInfo(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewChartInfoRegistry(sp builders.StandardStorageProvider) ChartInfoRegistry {
	return &storageChartInfo{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageChartInfo struct {
	builders.StandardStorageProvider
}

func (s *storageChartInfo) ListChartInfos(ctx context.Context, options *internalversion.ListOptions) (*ChartInfoList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ChartInfoList), err
}

func (s *storageChartInfo) GetChartInfo(ctx context.Context, id string, options *metav1.GetOptions) (*ChartInfo, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ChartInfo), nil
}

func (s *storageChartInfo) CreateChartInfo(ctx context.Context, object *ChartInfo) (*ChartInfo, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*ChartInfo), nil
}

func (s *storageChartInfo) UpdateChartInfo(ctx context.Context, object *ChartInfo) (*ChartInfo, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*ChartInfo), nil
}

func (s *storageChartInfo) DeleteChartInfo(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// ClusterQuota Functions and Structs
//
// +k8s:deepcopy-gen=false
type ClusterQuotaStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type ClusterQuotaStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterQuotaList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ClusterQuota
}

func (ClusterQuota) NewStatus() interface{} {
	return ClusterQuotaStatus{}
}

func (pc *ClusterQuota) GetStatus() interface{} {
	return pc.Status
}

func (pc *ClusterQuota) SetStatus(s interface{}) {
	pc.Status = s.(ClusterQuotaStatus)
}

func (pc *ClusterQuota) GetSpec() interface{} {
	return pc.Spec
}

func (pc *ClusterQuota) SetSpec(s interface{}) {
	pc.Spec = s.(ClusterQuotaSpec)
}

func (pc *ClusterQuota) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *ClusterQuota) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc ClusterQuota) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store ClusterQuota.
// +k8s:deepcopy-gen=false
type ClusterQuotaRegistry interface {
	ListClusterQuotas(ctx context.Context, options *internalversion.ListOptions) (*ClusterQuotaList, error)
	GetClusterQuota(ctx context.Context, id string, options *metav1.GetOptions) (*ClusterQuota, error)
	CreateClusterQuota(ctx context.Context, id *ClusterQuota) (*ClusterQuota, error)
	UpdateClusterQuota(ctx context.Context, id *ClusterQuota) (*ClusterQuota, error)
	DeleteClusterQuota(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewClusterQuotaRegistry(sp builders.StandardStorageProvider) ClusterQuotaRegistry {
	return &storageClusterQuota{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageClusterQuota struct {
	builders.StandardStorageProvider
}

func (s *storageClusterQuota) ListClusterQuotas(ctx context.Context, options *internalversion.ListOptions) (*ClusterQuotaList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ClusterQuotaList), err
}

func (s *storageClusterQuota) GetClusterQuota(ctx context.Context, id string, options *metav1.GetOptions) (*ClusterQuota, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ClusterQuota), nil
}

func (s *storageClusterQuota) CreateClusterQuota(ctx context.Context, object *ClusterQuota) (*ClusterQuota, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*ClusterQuota), nil
}

func (s *storageClusterQuota) UpdateClusterQuota(ctx context.Context, object *ClusterQuota) (*ClusterQuota, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*ClusterQuota), nil
}

func (s *storageClusterQuota) DeleteClusterQuota(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// HelmRelease Functions and Structs
//
// +k8s:deepcopy-gen=false
type HelmReleaseStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type HelmReleaseStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmReleaseList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []HelmRelease
}

func (HelmRelease) NewStatus() interface{} {
	return HelmReleaseStatus{}
}

func (pc *HelmRelease) GetStatus() interface{} {
	return pc.Status
}

func (pc *HelmRelease) SetStatus(s interface{}) {
	pc.Status = s.(HelmReleaseStatus)
}

func (pc *HelmRelease) GetSpec() interface{} {
	return pc.Spec
}

func (pc *HelmRelease) SetSpec(s interface{}) {
	pc.Spec = s.(HelmReleaseSpec)
}

func (pc *HelmRelease) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *HelmRelease) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc HelmRelease) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store HelmRelease.
// +k8s:deepcopy-gen=false
type HelmReleaseRegistry interface {
	ListHelmReleases(ctx context.Context, options *internalversion.ListOptions) (*HelmReleaseList, error)
	GetHelmRelease(ctx context.Context, id string, options *metav1.GetOptions) (*HelmRelease, error)
	CreateHelmRelease(ctx context.Context, id *HelmRelease) (*HelmRelease, error)
	UpdateHelmRelease(ctx context.Context, id *HelmRelease) (*HelmRelease, error)
	DeleteHelmRelease(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewHelmReleaseRegistry(sp builders.StandardStorageProvider) HelmReleaseRegistry {
	return &storageHelmRelease{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageHelmRelease struct {
	builders.StandardStorageProvider
}

func (s *storageHelmRelease) ListHelmReleases(ctx context.Context, options *internalversion.ListOptions) (*HelmReleaseList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*HelmReleaseList), err
}

func (s *storageHelmRelease) GetHelmRelease(ctx context.Context, id string, options *metav1.GetOptions) (*HelmRelease, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*HelmRelease), nil
}

func (s *storageHelmRelease) CreateHelmRelease(ctx context.Context, object *HelmRelease) (*HelmRelease, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*HelmRelease), nil
}

func (s *storageHelmRelease) UpdateHelmRelease(ctx context.Context, object *HelmRelease) (*HelmRelease, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*HelmRelease), nil
}

func (s *storageHelmRelease) DeleteHelmRelease(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// LocalClusterAccess Functions and Structs
//
// +k8s:deepcopy-gen=false
type LocalClusterAccessStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type LocalClusterAccessStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LocalClusterAccessList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []LocalClusterAccess
}

func (LocalClusterAccess) NewStatus() interface{} {
	return LocalClusterAccessStatus{}
}

func (pc *LocalClusterAccess) GetStatus() interface{} {
	return pc.Status
}

func (pc *LocalClusterAccess) SetStatus(s interface{}) {
	pc.Status = s.(LocalClusterAccessStatus)
}

func (pc *LocalClusterAccess) GetSpec() interface{} {
	return pc.Spec
}

func (pc *LocalClusterAccess) SetSpec(s interface{}) {
	pc.Spec = s.(LocalClusterAccessSpec)
}

func (pc *LocalClusterAccess) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *LocalClusterAccess) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc LocalClusterAccess) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store LocalClusterAccess.
// +k8s:deepcopy-gen=false
type LocalClusterAccessRegistry interface {
	ListLocalClusterAccesss(ctx context.Context, options *internalversion.ListOptions) (*LocalClusterAccessList, error)
	GetLocalClusterAccess(ctx context.Context, id string, options *metav1.GetOptions) (*LocalClusterAccess, error)
	CreateLocalClusterAccess(ctx context.Context, id *LocalClusterAccess) (*LocalClusterAccess, error)
	UpdateLocalClusterAccess(ctx context.Context, id *LocalClusterAccess) (*LocalClusterAccess, error)
	DeleteLocalClusterAccess(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewLocalClusterAccessRegistry(sp builders.StandardStorageProvider) LocalClusterAccessRegistry {
	return &storageLocalClusterAccess{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageLocalClusterAccess struct {
	builders.StandardStorageProvider
}

func (s *storageLocalClusterAccess) ListLocalClusterAccesss(ctx context.Context, options *internalversion.ListOptions) (*LocalClusterAccessList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*LocalClusterAccessList), err
}

func (s *storageLocalClusterAccess) GetLocalClusterAccess(ctx context.Context, id string, options *metav1.GetOptions) (*LocalClusterAccess, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*LocalClusterAccess), nil
}

func (s *storageLocalClusterAccess) CreateLocalClusterAccess(ctx context.Context, object *LocalClusterAccess) (*LocalClusterAccess, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*LocalClusterAccess), nil
}

func (s *storageLocalClusterAccess) UpdateLocalClusterAccess(ctx context.Context, object *LocalClusterAccess) (*LocalClusterAccess, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*LocalClusterAccess), nil
}

func (s *storageLocalClusterAccess) DeleteLocalClusterAccess(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// SleepModeConfig Functions and Structs
//
// +k8s:deepcopy-gen=false
type SleepModeConfigStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type SleepModeConfigStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SleepModeConfigList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []SleepModeConfig
}

func (SleepModeConfig) NewStatus() interface{} {
	return SleepModeConfigStatus{}
}

func (pc *SleepModeConfig) GetStatus() interface{} {
	return pc.Status
}

func (pc *SleepModeConfig) SetStatus(s interface{}) {
	pc.Status = s.(SleepModeConfigStatus)
}

func (pc *SleepModeConfig) GetSpec() interface{} {
	return pc.Spec
}

func (pc *SleepModeConfig) SetSpec(s interface{}) {
	pc.Spec = s.(SleepModeConfigSpec)
}

func (pc *SleepModeConfig) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *SleepModeConfig) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc SleepModeConfig) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store SleepModeConfig.
// +k8s:deepcopy-gen=false
type SleepModeConfigRegistry interface {
	ListSleepModeConfigs(ctx context.Context, options *internalversion.ListOptions) (*SleepModeConfigList, error)
	GetSleepModeConfig(ctx context.Context, id string, options *metav1.GetOptions) (*SleepModeConfig, error)
	CreateSleepModeConfig(ctx context.Context, id *SleepModeConfig) (*SleepModeConfig, error)
	UpdateSleepModeConfig(ctx context.Context, id *SleepModeConfig) (*SleepModeConfig, error)
	DeleteSleepModeConfig(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewSleepModeConfigRegistry(sp builders.StandardStorageProvider) SleepModeConfigRegistry {
	return &storageSleepModeConfig{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageSleepModeConfig struct {
	builders.StandardStorageProvider
}

func (s *storageSleepModeConfig) ListSleepModeConfigs(ctx context.Context, options *internalversion.ListOptions) (*SleepModeConfigList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfigList), err
}

func (s *storageSleepModeConfig) GetSleepModeConfig(ctx context.Context, id string, options *metav1.GetOptions) (*SleepModeConfig, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfig), nil
}

func (s *storageSleepModeConfig) CreateSleepModeConfig(ctx context.Context, object *SleepModeConfig) (*SleepModeConfig, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfig), nil
}

func (s *storageSleepModeConfig) UpdateSleepModeConfig(ctx context.Context, object *SleepModeConfig) (*SleepModeConfig, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*SleepModeConfig), nil
}

func (s *storageSleepModeConfig) DeleteSleepModeConfig(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// Space Functions and Structs
//
// +k8s:deepcopy-gen=false
type SpaceStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type SpaceStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SpaceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Space
}

func (Space) NewStatus() interface{} {
	return SpaceStatus{}
}

func (pc *Space) GetStatus() interface{} {
	return pc.Status
}

func (pc *Space) SetStatus(s interface{}) {
	pc.Status = s.(SpaceStatus)
}

func (pc *Space) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Space) SetSpec(s interface{}) {
	pc.Spec = s.(SpaceSpec)
}

func (pc *Space) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Space) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Space) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Space.
// +k8s:deepcopy-gen=false
type SpaceRegistry interface {
	ListSpaces(ctx context.Context, options *internalversion.ListOptions) (*SpaceList, error)
	GetSpace(ctx context.Context, id string, options *metav1.GetOptions) (*Space, error)
	CreateSpace(ctx context.Context, id *Space) (*Space, error)
	UpdateSpace(ctx context.Context, id *Space) (*Space, error)
	DeleteSpace(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewSpaceRegistry(sp builders.StandardStorageProvider) SpaceRegistry {
	return &storageSpace{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageSpace struct {
	builders.StandardStorageProvider
}

func (s *storageSpace) ListSpaces(ctx context.Context, options *internalversion.ListOptions) (*SpaceList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*SpaceList), err
}

func (s *storageSpace) GetSpace(ctx context.Context, id string, options *metav1.GetOptions) (*Space, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Space), nil
}

func (s *storageSpace) CreateSpace(ctx context.Context, object *Space) (*Space, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Space), nil
}

func (s *storageSpace) UpdateSpace(ctx context.Context, object *Space) (*Space, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Space), nil
}

func (s *storageSpace) DeleteSpace(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}

//
// VirtualCluster Functions and Structs
//
// +k8s:deepcopy-gen=false
type VirtualClusterStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type VirtualClusterStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualClusterList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []VirtualCluster
}

func (VirtualCluster) NewStatus() interface{} {
	return VirtualClusterStatus{}
}

func (pc *VirtualCluster) GetStatus() interface{} {
	return pc.Status
}

func (pc *VirtualCluster) SetStatus(s interface{}) {
	pc.Status = s.(VirtualClusterStatus)
}

func (pc *VirtualCluster) GetSpec() interface{} {
	return pc.Spec
}

func (pc *VirtualCluster) SetSpec(s interface{}) {
	pc.Spec = s.(VirtualClusterSpec)
}

func (pc *VirtualCluster) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *VirtualCluster) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc VirtualCluster) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store VirtualCluster.
// +k8s:deepcopy-gen=false
type VirtualClusterRegistry interface {
	ListVirtualClusters(ctx context.Context, options *internalversion.ListOptions) (*VirtualClusterList, error)
	GetVirtualCluster(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualCluster, error)
	CreateVirtualCluster(ctx context.Context, id *VirtualCluster) (*VirtualCluster, error)
	UpdateVirtualCluster(ctx context.Context, id *VirtualCluster) (*VirtualCluster, error)
	DeleteVirtualCluster(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewVirtualClusterRegistry(sp builders.StandardStorageProvider) VirtualClusterRegistry {
	return &storageVirtualCluster{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageVirtualCluster struct {
	builders.StandardStorageProvider
}

func (s *storageVirtualCluster) ListVirtualClusters(ctx context.Context, options *internalversion.ListOptions) (*VirtualClusterList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualClusterList), err
}

func (s *storageVirtualCluster) GetVirtualCluster(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualCluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualCluster), nil
}

func (s *storageVirtualCluster) CreateVirtualCluster(ctx context.Context, object *VirtualCluster) (*VirtualCluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualCluster), nil
}

func (s *storageVirtualCluster) UpdateVirtualCluster(ctx context.Context, object *VirtualCluster) (*VirtualCluster, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualCluster), nil
}

func (s *storageVirtualCluster) DeleteVirtualCluster(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}
