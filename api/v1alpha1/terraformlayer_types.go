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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TerraformLayerSpec defines the desired state of TerraformLayer
// +kubebuilder:validation:XValidation:rule="!(has(self.terraform) && has(self.opentofu) && has(self.terraform.enabled) && has(self.opentofu.enabled) && self.terraform.enabled == true && self.opentofu.enabled == true)",message="Both terraform.enabled and opentofu.enabled cannot be true at the same time"
// +kubebuilder:validation:XValidation:rule="!(has(self.terraform) && has(self.opentofu) && has(self.terraform.enabled) && has(self.opentofu.enabled) && self.terraform.enabled == false && self.opentofu.enabled == false)",message="Both terraform.enabled and opentofu.enabled cannot be false at the same time"
type TerraformLayerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Path                string                   `json:"path,omitempty"`
	Branch              string                   `json:"branch,omitempty"`
	AllowPrOnTags       *bool                    `json:"allowPrOnTags,omitempty"`
	TerraformConfig     TerraformConfig          `json:"terraform,omitempty"`
	OpenTofuConfig      OpenTofuConfig           `json:"opentofu,omitempty"`
	TerragruntConfig    TerragruntConfig         `json:"terragrunt,omitempty"`
	Repository          TerraformLayerRepository `json:"repository,omitempty"`
	RemediationStrategy RemediationStrategy      `json:"remediationStrategy,omitempty"`
	OverrideRunnerSpec  OverrideRunnerSpec       `json:"overrideRunnerSpec,omitempty"`
	RunHistoryPolicy    RunHistoryPolicy         `json:"runHistoryPolicy,omitempty"`
}

type TerraformLayerRepository struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// TerraformLayerStatus defines the observed state of TerraformLayer
type TerraformLayerStatus struct {
	Conditions []metav1.Condition  `json:"conditions,omitempty"`
	State      string              `json:"state,omitempty"`
	LastResult string              `json:"lastResult,omitempty"`
	LastRun    TerraformLayerRun   `json:"lastRun,omitempty"`
	LatestRuns []TerraformLayerRun `json:"latestRuns,omitempty"`
}

type TerraformLayerRun struct {
	Name   string      `json:"name,omitempty"`
	Commit string      `json:"commit,omitempty"`
	Date   metav1.Time `json:"date,omitempty"`
	Action string      `json:"action,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=layers;layer;tfls;tfl;
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Repository",type=string,JSONPath=`.spec.repository.name`
// +kubebuilder:printcolumn:name="Branch",type=string,JSONPath=`.spec.branch`
// +kubebuilder:printcolumn:name="Path",type=string,JSONPath=`.spec.path`
// +kubebuilder:printcolumn:name="Last Result",type=string,JSONPath=`.status.lastResult`
// TerraformLayer is the Schema for the terraformlayers API
type TerraformLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TerraformLayerSpec   `json:"spec,omitempty"`
	Status TerraformLayerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TerraformLayerList contains a list of TerraformLayer
type TerraformLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TerraformLayer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TerraformLayer{}, &TerraformLayerList{})
}

// Associated functions

// Workaround needed for envtest which does not populate the TypeMeta structure
// See https://github.com/kubernetes-sigs/controller-runtime/issues/1870
func (layer *TerraformLayer) GetAPIVersion() string {
	if layer.APIVersion == "" {
		return "config.terraform.padok.cloud/v1alpha1"
	}
	return layer.APIVersion
}

// Same as above
func (layer *TerraformLayer) GetKind() string {
	if layer.Kind == "" {
		return "TerraformLayer"
	}
	return layer.Kind
}
