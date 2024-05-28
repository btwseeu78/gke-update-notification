/*
Copyright 2024.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Metric struct {
	Name string `json:"name"`
	// +optional
	Labels      map[string]string `json:"labels,omitempty"`
	Type        string            `json:"type"`
	ExporterUrl string            `json:"exporterUrl"`
}

// NotificationConfigSpec defines the desired state of NotificationConfig
type NotificationConfigSpec struct {
	Metrics          Metric                `json:"metrics"`
	BackendSecretRef *v1.SecretKeySelector `json:"backendSecretRef"`
	CheckInterval    *metav1.Duration      `json:"checkInterval,omitempty"`
}

// NotificationConfigStatus defines the observed state of NotificationConfig
type NotificationConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LastUpdateReason             string      `json:"lastUpdateReason,omitempty"`
	LastUpdateTime               metav1.Time `json:"lastUpdateTime,omitempty"`
	LastNotificationStatus       string      `json:"lastNotificationStatus,omitempty"`
	LastNotificationFailedReason string      `json:"lastNotificationFailedReason,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="LastUpdateReason",type=string,JSONPath=`.status.lastUpdateReason`
// +kubebuilder:printcolumn:name="LastUpdateTime",type=date,JSONPath=`.status.lastUpdateTime`
// +kubebuilder:printcolumn:name="LastNotificationStatus",type=string,JSONPath=`.status.lastNotificationStatus`
// NotificationConfig is the Schema for the notificationconfigs API
type NotificationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotificationConfigSpec   `json:"spec,omitempty"`
	Status NotificationConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NotificationConfigList contains a list of NotificationConfig
type NotificationConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotificationConfig{}, &NotificationConfigList{})
}
