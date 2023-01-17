package v1

import (
	fmt "fmt"

	corev1 "k8s.io/api/core/v1"
)

// PodNameFromIndex returns the pod name for the given index e.g kafka-0
func PodNameFromIndex(name string, index uint32) string {
	return fmt.Sprintf("%s-%d", name, index)
}

//ExternalPodServiceName returns the name of the external k8s service for a pod
func ExternalPodServiceName(podName string) string {
	return fmt.Sprintf("%s-external", podName)
}

//InternalPodServiceName returns the name of the internal k8s service for a pod
func InternalPodServiceName(podName string) string {
	return fmt.Sprintf("%s-internal", podName)
}

//CommonPorts returns common ports for all services
func CommonPorts() []*Port {
	return []*Port{
		{
			Name:     JMXServiceName,
			Protocol: string(corev1.ProtocolTCP),
			Port:     JMXServicePort,
		},
		{
			Name:     JolokiaServiceName,
			Protocol: string(corev1.ProtocolTCP),
			Port:     JolokiaServicePort,
		},
	}
}
