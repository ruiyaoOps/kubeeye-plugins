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
	"github.com/aquasecurity/kube-bench/check"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubeBenchSpec defines the desired state of KubeBench
type KubeBenchSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of KubeBench. Edit kubebench_types.go to remove/update
	AuditPeriod string `json:"auditPeriod"`
}

// KubeBenchStatus defines the observed state of KubeBench
type KubeBenchStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Controls []Controls `json:"controls,omitempty"`
}

type Controls struct {
	Version       string  `json:"version"`
	Text          string  `json:"text"`
	Groups        []Group `json:"groups"`
	check.Summary `json:"summary"`
}

type Group struct {
	ID     string  `json:"groupid"`
	Text   string  `json:"type"`
	Checks []Check `json:"results"`
}

type Check struct {
	ID          string `yaml:"id" json:"id"`
	Text        string `json:"desc"`
	Remediation string `json:"remediation"`
	check.State `json:"status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KubeBench is the Schema for the kubebenches API
type KubeBench struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeBenchSpec   `json:"spec,omitempty"`
	Status KubeBenchStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KubeBenchList contains a list of KubeBench
type KubeBenchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeBench `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeBench{}, &KubeBenchList{})
}
