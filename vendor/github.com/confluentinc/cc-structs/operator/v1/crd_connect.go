package v1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindConnector = "Connector"
	ResourceNameConnector = "connector"
	ResourceTypeConnector = "connectors"
)

func (m Connector) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: ResourceTypeConnector + "." + SchemeGroupVersion.Group,
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   GroupName,
			Version: SchemeGroupVersion.Version,
			Scope:   apiextensions.NamespaceScoped,
			Names: apiextensions.CustomResourceDefinitionNames{
				Singular:   ResourceNameConnector,
				Plural:     ResourceTypeConnector,
				Kind:       ResourceKindConnector,
				ShortNames: []string{"connect"},
			},
		},
	}
}
