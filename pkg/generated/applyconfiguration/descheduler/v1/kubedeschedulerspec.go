// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apioperatorv1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	deschedulerv1 "github.com/openshift/cluster-kube-descheduler-operator/pkg/apis/descheduler/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// KubeDeschedulerSpecApplyConfiguration represents a declarative configuration of the KubeDeschedulerSpec type for use
// with apply.
type KubeDeschedulerSpecApplyConfiguration struct {
	operatorv1.OperatorSpecApplyConfiguration `json:",inline"`
	Profiles                                  []deschedulerv1.DeschedulerProfile       `json:"profiles,omitempty"`
	DeschedulingIntervalSeconds               *int32                                   `json:"deschedulingIntervalSeconds,omitempty"`
	EvictionLimits                            *EvictionLimitsApplyConfiguration        `json:"evictionLimits,omitempty"`
	ProfileCustomizations                     *ProfileCustomizationsApplyConfiguration `json:"profileCustomizations,omitempty"`
	Mode                                      *deschedulerv1.Mode                      `json:"mode,omitempty"`
}

// KubeDeschedulerSpecApplyConfiguration constructs a declarative configuration of the KubeDeschedulerSpec type for use with
// apply.
func KubeDeschedulerSpec() *KubeDeschedulerSpecApplyConfiguration {
	return &KubeDeschedulerSpecApplyConfiguration{}
}

// WithManagementState sets the ManagementState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagementState field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithManagementState(value apioperatorv1.ManagementState) *KubeDeschedulerSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.ManagementState = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithLogLevel(value apioperatorv1.LogLevel) *KubeDeschedulerSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.LogLevel = &value
	return b
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithOperatorLogLevel(value apioperatorv1.LogLevel) *KubeDeschedulerSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.OperatorLogLevel = &value
	return b
}

// WithUnsupportedConfigOverrides sets the UnsupportedConfigOverrides field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnsupportedConfigOverrides field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithUnsupportedConfigOverrides(value runtime.RawExtension) *KubeDeschedulerSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.UnsupportedConfigOverrides = &value
	return b
}

// WithObservedConfig sets the ObservedConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedConfig field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithObservedConfig(value runtime.RawExtension) *KubeDeschedulerSpecApplyConfiguration {
	b.OperatorSpecApplyConfiguration.ObservedConfig = &value
	return b
}

// WithProfiles adds the given value to the Profiles field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Profiles field.
func (b *KubeDeschedulerSpecApplyConfiguration) WithProfiles(values ...deschedulerv1.DeschedulerProfile) *KubeDeschedulerSpecApplyConfiguration {
	for i := range values {
		b.Profiles = append(b.Profiles, values[i])
	}
	return b
}

// WithDeschedulingIntervalSeconds sets the DeschedulingIntervalSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeschedulingIntervalSeconds field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithDeschedulingIntervalSeconds(value int32) *KubeDeschedulerSpecApplyConfiguration {
	b.DeschedulingIntervalSeconds = &value
	return b
}

// WithEvictionLimits sets the EvictionLimits field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EvictionLimits field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithEvictionLimits(value *EvictionLimitsApplyConfiguration) *KubeDeschedulerSpecApplyConfiguration {
	b.EvictionLimits = value
	return b
}

// WithProfileCustomizations sets the ProfileCustomizations field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProfileCustomizations field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithProfileCustomizations(value *ProfileCustomizationsApplyConfiguration) *KubeDeschedulerSpecApplyConfiguration {
	b.ProfileCustomizations = value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *KubeDeschedulerSpecApplyConfiguration) WithMode(value deschedulerv1.Mode) *KubeDeschedulerSpecApplyConfiguration {
	b.Mode = &value
	return b
}
