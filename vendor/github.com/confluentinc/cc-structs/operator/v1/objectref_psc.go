package v1

import apiv1 "k8s.io/api/core/v1"

func (r PhysicalStatefulCluster) ObjectReference() *apiv1.ObjectReference {
	return &apiv1.ObjectReference{
		APIVersion:      SchemeGroupVersion.String(),
		Kind:            ResourceKindPhysicalStatefulCluster,
		Namespace:       r.Namespace,
		Name:            r.Name,
		UID:             r.UID,
		ResourceVersion: r.ResourceVersion,
	}
}