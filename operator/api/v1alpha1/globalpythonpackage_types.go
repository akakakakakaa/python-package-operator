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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GlobalPythonPackageStatusTypeDownloading = "Downloading"
	GlobalPythonPackageStatusTypeDownloaded  = "Downloaded"
	GlobalPythonPackageStatusTypeError       = "Error"
)

type GlobalPythonPackagePythonSpec struct {
	Version  string `json:"version"`
	Platform string `json:"platform"`
}

type GlobalPythonPackagePythonLibrarySpec struct {
	PythonSpec GlobalPythonPackagePythonSpec `json:"pythonSpec"`
	Name       string                        `json:"name"`
	Version    string                        `json:"version,omitempty"`
}

// GlobalPythonPackageSpec defines the desired state of GlobalPythonPackage
type GlobalPythonPackageSpec struct {
	GlobalPythonPackageCollectionName string                               `json:"globalPythonPackageCollectionName"`
	PythonLibrarySpec                 GlobalPythonPackagePythonLibrarySpec `json:"pythonLibrarySpec"`
}

// GlobalPythonPackageStatus defines the observed state of GlobalPythonPackage
type GlobalPythonPackageStatus struct {
	// Message shows log when the status changed in last
	Message string `json:"message,omitempty"`
	// Reason shows why the status changed in last
	Reason string `json:"reason,omitempty"`
	//+kubebuilder:validation:Enum=Downloading;Downloaded;Error;
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=gpp
//+kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// GlobalPythonPackage is the Schema for the globalpythonpackages API
type GlobalPythonPackage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalPythonPackageSpec   `json:"spec"`
	Status GlobalPythonPackageStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GlobalPythonPackageList contains a list of GlobalPythonPackage
type GlobalPythonPackageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalPythonPackage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlobalPythonPackage{}, &GlobalPythonPackageList{})
}
