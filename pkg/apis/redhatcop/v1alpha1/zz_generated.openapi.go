// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/redhat-cop/podpreset-webhook/pkg/apis/redhatcop/v1alpha1.PodPreset":     schema_pkg_apis_redhatcop_v1alpha1_PodPreset(ref),
		"github.com/redhat-cop/podpreset-webhook/pkg/apis/redhatcop/v1alpha1.PodPresetSpec": schema_pkg_apis_redhatcop_v1alpha1_PodPresetSpec(ref),
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_PodPreset(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodPreset is the Schema for the podpresets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/podpreset-webhook/pkg/apis/redhatcop/v1alpha1.PodPresetSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/podpreset-webhook/pkg/apis/redhatcop/v1alpha1.PodPresetSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_PodPresetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodPresetSpec defines the desired state of PodPreset",
				Properties: map[string]spec.Schema{
					"selector": {
						SchemaProps: spec.SchemaProps{
							Description: "Selector is a label query over a set of resources, in this case pods. Required.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Description: "Env defines the collection of EnvVar to inject into containers.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"envFrom": {
						SchemaProps: spec.SchemaProps{
							Description: "EnvFrom defines the collection of EnvFromSource to inject into containers.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvFromSource"),
									},
								},
							},
						},
					},
					"volumes": {
						SchemaProps: spec.SchemaProps{
							Description: "Volumes defines the collection of Volume to inject into the pod.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Volume"),
									},
								},
							},
						},
					},
					"volumeMounts": {
						SchemaProps: spec.SchemaProps{
							Description: "VolumeMounts defines the collection of VolumeMount to inject into containers.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.VolumeMount"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvFromSource", "k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.Volume", "k8s.io/api/core/v1.VolumeMount", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}
