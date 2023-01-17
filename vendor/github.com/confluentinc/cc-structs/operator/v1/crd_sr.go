package v1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindSchemaRegistryCluster = "SchemaRegistryCluster"
	ResourceNameSchemaRegistryCluster = "schemaregistrycluster"
	ResourceTypeSchemaRegistryCluster = "schemaregistryclusters"
)

func (m SchemaRegistryCluster) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: ResourceTypeSchemaRegistryCluster + "." + SchemeGroupVersion.Group,
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   GroupName,
			Version: SchemeGroupVersion.Version,
			Scope:   apiextensions.NamespaceScoped,
			Names: apiextensions.CustomResourceDefinitionNames{
				Singular:   ResourceNameSchemaRegistryCluster,
				Plural:     ResourceTypeSchemaRegistryCluster,
				Kind:       ResourceKindSchemaRegistryCluster,
				ShortNames: []string{"sr"},
			},
		},
	}
}
