package v1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindPhysicalStatefulCluster = "PhysicalStatefulCluster"
	ResourceNamePhysicalStatefulCluster = "physicalstatefulcluster"
	ResourceTypePhysicalStatefulCluster = "physicalstatefulclusters"

	ResourceKindLogicalKafkaCluster = "LogicalKafkaCluster"
	ResourceNameLogicalKafkaCluster = "logicalkafkacluster"
	ResourceTypeLogicalKafkaCluster = "logicalkafkaclusters"
)

func (psc PhysicalStatefulCluster) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: ResourceTypePhysicalStatefulCluster + "." + SchemeGroupVersion.Group,
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   GroupName,
			Version: SchemeGroupVersion.Version,
			Scope:   apiextensions.NamespaceScoped,
			Names: apiextensions.CustomResourceDefinitionNames{
				Singular:   ResourceNamePhysicalStatefulCluster,
				Plural:     ResourceTypePhysicalStatefulCluster,
				Kind:       ResourceKindPhysicalStatefulCluster,
				ShortNames: []string{"psc"},
			},
		},
	}
}

func (lkc LogicalKafkaCluster) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: ResourceKindLogicalKafkaCluster + "." + SchemeGroupVersion.Group,
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   GroupName,
			Version: SchemeGroupVersion.Version,
			Scope:   apiextensions.NamespaceScoped,
			Names: apiextensions.CustomResourceDefinitionNames{
				Singular:   ResourceNameLogicalKafkaCluster,
				Plural:     ResourceTypeLogicalKafkaCluster,
				Kind:       ResourceKindLogicalKafkaCluster,
				ShortNames: []string{"lkc"},
			},
		},
	}
}
