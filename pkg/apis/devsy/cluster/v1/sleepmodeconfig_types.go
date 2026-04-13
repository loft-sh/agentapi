package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SleepModeForceAnnotation          = "sleepmode.devsy.sh/force"
	SleepModeForceDurationAnnotation  = "sleepmode.devsy.sh/force-duration"
	SleepModeSleepAfterAnnotation     = "sleepmode.devsy.sh/sleep-after"
	SleepModeDeleteAfterAnnotation    = "sleepmode.devsy.sh/delete-after"
	SleepModeSleepScheduleAnnotation  = "sleepmode.devsy.sh/sleep-schedule"
	SleepModeWakeupScheduleAnnotation = "sleepmode.devsy.sh/wakeup-schedule"
	SleepModeTimezoneAnnotation       = "sleepmode.devsy.sh/timezone"

	SleepModeLastActivityAnnotation      = "sleepmode.devsy.sh/last-activity"
	SleepModeLastActivityAnnotationInfo  = "sleepmode.devsy.sh/last-activity-info"
	SleepModeSleepingSinceAnnotation     = "sleepmode.devsy.sh/sleeping-since"
	SleepModeCurrentEpochStartAnnotation = "sleepmode.devsy.sh/current-epoch-start"
	SleepModeCurrentEpochSleptAnnotation = "sleepmode.devsy.sh/current-epoch-slept"
	SleepModeLastEpochStartAnnotation    = "sleepmode.devsy.sh/last-epoch-start"
	SleepModeLastEpochSleptAnnotation    = "sleepmode.devsy.sh/last-epoch-slept"
	SleepModeScheduledSleepAnnotation    = "sleepmode.devsy.sh/scheduled-sleep"
	SleepModeScheduledWakeupAnnotation   = "sleepmode.devsy.sh/scheduled-wakeup"
	SleepModeSleepTypeAnnotation         = "sleepmode.devsy.sh/sleep-type"
	SleepModeDisableIngressWakeup        = "sleepmode.devsy.sh/disable-ingress-wakeup"
	SleepModeDisableMetricsTracking      = "sleepmode.devsy.sh/disable-metrics-tracking"
	SleepScopeAnnotation                 = "sleepmode.devsy.sh/scope"

	// Not yet in spec annotations
	SleepModeIgnoreAll                     = "sleepmode.devsy.sh/ignore-all"
	SleepModeIgnoreIngresses               = "sleepmode.devsy.sh/ignore-ingresses"
	SleepModeIgnoreGroupsAnnotation        = "sleepmode.devsy.sh/ignore-groups"
	SleepModeIgnoreDevsysAnnotation        = "sleepmode.devsy.sh/ignore-devsys"
	SleepModeIgnoreResourcesAnnotation     = "sleepmode.devsy.sh/ignore-resources"
	SleepModeIgnoreVerbsAnnotation         = "sleepmode.devsy.sh/ignore-verbs"
	SleepModeIgnoreResourceVerbsAnnotation = "sleepmode.devsy.sh/ignore-resource-verbs" // format: myresource.mygroup=create update delete, myresource2.mygroup=create update
	SleepModeIgnoreResourceNamesAnnotation = "sleepmode.devsy.sh/ignore-resource-names" // format: myresource.mygroup=name1 name2
	SleepModeIgnoreActiveConnections       = "sleepmode.devsy.sh/ignore-active-connections"
	SleepModeIgnoreUserAgents              = "sleepmode.devsy.sh/ignore-user-agents" // format: useragent1,useragentprefix/*,*

	SleepTypeInactivity     = "inactivitySleep"
	SleepTypeForced         = "forcedSleep"
	SleepTypeForcedDuration = "forcedDurationSleep"
	SleepTypeScheduled      = "scheduledSleep"
)

// SleepModeConfigAnnotationKeys returns relevant annotation keys that are evaluated in
// apimachinery/v2/pkg/sleepmode/extractSleepModeConfigs function. This is primarily used for
// knowing which annotations to copy from spaces to devsy instances when importing devsys
// to a project.
func SleepModeConfigAnnotationKeys() []string {
	return []string{
		SleepModeForceDurationAnnotation,
		SleepModeForceAnnotation,
		SleepModeScheduledSleepAnnotation,
		SleepModeScheduledWakeupAnnotation,
	}
}

func SleepModeStatusAnnotationKeys() []string {
	return []string{
		SleepModeForceDurationAnnotation,
		SleepModeForceAnnotation,
		SleepModeScheduledSleepAnnotation,
		SleepModeScheduledWakeupAnnotation,
		SleepModeSleepingSinceAnnotation,
		SleepModeSleepTypeAnnotation,
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SleepModeConfig holds the sleepmode information
type SleepModeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SleepModeConfigSpec   `json:"spec,omitempty"`
	Status SleepModeConfigStatus `json:"status,omitempty"`
}

type SleepModeConfigSpec struct {
	// If force sleep is true the space will sleep
	// +optional
	ForceSleep bool `json:"forceSleep,omitempty"`

	// If force sleep duration is set, this will force the space to sleep
	// for the given duration. It also implies that forceSleep is true.
	// During this period devsy will also block certain requests to that space.
	// If this is set to 0, it means the space will sleep until it is manually
	// woken up via the cli or ui.
	// +optional
	ForceSleepDuration *int64 `json:"forceSleepDuration,omitempty"`

	// DeleteAfter specifies after how many seconds of inactivity the space should be deleted
	// +optional
	DeleteAfter int64 `json:"deleteAfter,omitempty"`

	// SleepAfter specifies after how many seconds of inactivity the space should sleep
	// +optional
	SleepAfter int64 `json:"sleepAfter,omitempty"`

	// SleepSchedule specifies scheduled space sleep in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// Note: timezone defined in the schedule string will be ignored. Use ".Spec.Timezone" field instead.
	// +optional
	SleepSchedule string `json:"sleepSchedule,omitempty"`

	// WakeupSchedule specifies scheduled wakeup from sleep in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// Note: timezone defined in the schedule string will be ignored. Use ".Spec.Timezone" field instead.
	// +optional
	WakeupSchedule string `json:"wakeupSchedule,omitempty"`

	// Timezone specifies time zone used for scheduled space operations. Defaults to UTC.
	// Accepts the same format as time.LoadLocation() in Go (https://pkg.go.dev/time#LoadLocation).
	// The value should be a location name corresponding to a file in the IANA Time Zone database, such as "America/New_York".
	// See also: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	// +optional
	Timezone string `json:"timezone,omitempty"`

	// IgnoreActiveConnections ignores active connections on the namespace
	// +optional
	IgnoreActiveConnections bool `json:"ignoreActiveConnections,omitempty"`

	// IgnoreAll ignores all requests
	// +optional
	IgnoreAll bool `json:"ignoreAll,omitempty"`

	// IgnoreIngresses ignores all ingresses
	// +optional
	IgnoreIngresses bool `json:"ignoreIngresses,omitempty"`

	// IgnoreDevsys ignores devsy requests
	// +optional
	IgnoreDevsys bool `json:"ignoreDevsys,omitempty"`

	// IgnoreGroups are ignored user groups
	// +optional
	IgnoreGroups string `json:"ignoreGroups,omitempty"`

	// IgnoreVerbs are ignored request verbs
	// +optional
	IgnoreVerbs string `json:"ignoreVerbs,omitempty"`

	// IgnoreResources are ignored request resources
	// +optional
	IgnoreResources string `json:"ignoreResources,omitempty"`

	// IgnoreResourceVerbs are ignored resource verbs
	// +optional
	IgnoreResourceVerbs string `json:"ignoreResourceVerbs,omitempty"`

	// IgnoreResourceNames are ignored resources and names
	// +optional
	IgnoreResourceNames string `json:"ignoreResourceNames,omitempty"`

	// IgnoreUseragents are ignored user agents with trailing wildcards '*' allowed.
	// comma separated
	// +optional
	IgnoreUseragents string `json:"ignoreUserAgents,omitempty"`
}

type SleepModeConfigStatus struct {
	// LastActivity indicates the last activity in the namespace
	// +optional
	LastActivity int64 `json:"lastActivity,omitempty"`

	// LastActivityInfo holds information about the last activity within this space
	// +optional
	LastActivityInfo *LastActivityInfo `json:"lastActivityInfo,omitempty"`

	// SleepingSince specifies since when the space is sleeping (if this is not specified, devsy assumes the space is not sleeping)
	// +optional
	SleepingSince int64 `json:"sleepingSince,omitempty"`

	// Optional info that indicates how long the space was sleeping in the current epoch
	// +optional
	CurrentEpoch *EpochInfo `json:"currentEpoch,omitempty"`

	// Optional info that indicates how long the space was sleeping in the current epoch
	// +optional
	LastEpoch *EpochInfo `json:"lastEpoch,omitempty"`

	// This is a calculated field that will be returned but not saved and describes the percentage since the space
	// was created or the last 30 days the space has slept
	// +optional
	SleptLastThirtyDays *float64 `json:"sleptLastThirtyDays,omitempty"`

	// This is a calculated field that will be returned but not saved and describes the percentage since the space
	// was created or the last 7 days the space has slept
	// +optional
	SleptLastSevenDays *float64 `json:"sleptLastSevenDays,omitempty"`

	// Indicates time of the next scheduled sleep based on .Spec.SleepSchedule and .Spec.ScheduleTimeZone
	// The time is a Unix time, the number of seconds elapsed since January 1, 1970 UTC
	// +optional
	ScheduledSleep *int64 `json:"scheduledSleep,omitempty"`

	// Indicates time of the next scheduled wakeup based on .Spec.WakeupSchedule and .Spec.ScheduleTimeZone
	// The time is a Unix time, the number of seconds elapsed since January 1, 1970 UTC
	// +optional
	ScheduledWakeup *int64 `json:"scheduledWakeup,omitempty"`

	// SleepType specifies a type of sleep, which has effect on which actions will cause the space to wake up.
	// +optional
	SleepType string `json:"sleepType,omitempty"`
}

// EpochInfo holds information about how long the space was sleeping in the epoch
type EpochInfo struct {
	// Timestamp when the epoch has started
	// +optional
	Start int64 `json:"start,omitempty"`
	// Amount of milliseconds the space has slept in the epoch
	// +optional
	Slept int64 `json:"slept,omitempty"`
}

// LastActivityInfo holds information about the last activity
type LastActivityInfo struct {
	// Subject is the user or team where this activity was recorded
	// +optional
	Subject string `json:"subject,omitempty"`

	// Host is the host where this activity was recorded
	// +optional
	Host string `json:"host,omitempty"`

	// Verb is the verb that was used for the request
	// +optional
	Verb string `json:"verb,omitempty"`

	// APIGroup is the api group that was used for the request
	// +optional
	APIGroup string `json:"apiGroup,omitempty"`

	// Resource is the resource of the request
	// +optional
	Resource string `json:"resource,omitempty"`

	// Subresource is the subresource of the request
	// +optional
	Subresource string `json:"subresource,omitempty"`

	// Name is the name of the resource
	// +optional
	Name string `json:"name,omitempty"`

	// DevsyCluster is the devsy cluster this activity happened in
	// +optional
	DevsyCluster string `json:"devsyCluster,omitempty"`

	// MetricsRefreshInterval is the activity refresh interval. This is used to prevent sleeping instances if the last
	// activity metrics have not been refreshed within the interval. Useful for metrics based activty tracking.
	MetricsRefreshInterval int64 `json:"metricsRefreshInterval,omitempty"`
}
