/*
Copyright 2022.

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

const (
	GlobalPythonPackageCollectionStatusTypeCreated = "Created"
	GlobalPythonPackageCollectionStatusTypeError   = "Error"
)

// GlobalPythonPackageCollectionSpec defines the desired state of GlobalPythonPackageCollection
type GlobalPythonPackageCollectionSpec struct {
	PersistentVolumeClaimSpec v1.PersistentVolumeClaimSpec `json:"persistentVolumeClaimSpec"`
}

// GlobalPythonPackageCollectionStatus defines the observed state of GlobalPythonPackageCollection
type GlobalPythonPackageCollectionStatus struct {
	// Message shows log when the status changed in last
	Message string `json:"message,omitempty"`
	// Reason shows why the status changed in last
	Reason string `json:"reason,omitempty"`
	//+kubebuilder:validation:Enum=Created;Error;
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=gppc
//+kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// GlobalPythonPackageCollection is the Schema for the globalpythonpackagecollections API
type GlobalPythonPackageCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalPythonPackageCollectionSpec   `json:"spec"`
	Status GlobalPythonPackageCollectionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GlobalPythonPackageCollectionList contains a list of GlobalPythonPackageCollection
type GlobalPythonPackageCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalPythonPackageCollection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlobalPythonPackageCollection{}, &GlobalPythonPackageCollectionList{})
}
