package v1

const (
	physicalStatefulClusterCoreKey = "physicalstatefulcluster.core.confluent.cloud"
	//PhysicalStatefulClusterLastSyncedPodTemplateSpecDiff is an annotation set on PhysicalStatefulCluster used in the psc controller to check whether a pod template spec has changed when the psc remains the same
	PhysicalStatefulClusterLastSyncedPodTemplateSpecDiff = physicalStatefulClusterCoreKey + "/last-synced-pod-template-spec-diff"
	//PhysicalStatefulClusterLastSyncedPodTemplateSpecHash is an annotation set on PhysicalStatefulCluster used in the psc controller to determine if a roll is required or not. The hash is generated on the object constructed by the operator. The LastSyncedPodTemplateSpecDiff annotation is used to figure out if any pod template parameter has been changed out of band. We can't use the same because the diff between the object stored in k8s and the one we create will always have the difference because of the default assigned by k8s.
	PhysicalStatefulClusterLastSyncedPodTemplateSpecHash = physicalStatefulClusterCoreKey + "/last-synced-pod-template-spec-hash"
	//PhysicalStatefulClusterLastSyncedOperatorVersionAnnotation is an annotation set on PhysicalStatefulCluster CRs by the CC operator that stores the semantic version (Docker image tag) of last operator version to apply the PSC
	PhysicalStatefulClusterLastSyncedOperatorVersionAnnotation = physicalStatefulClusterCoreKey + "/last-synced-operator-version"
	//PhysicalStatefulClusterCoreLastSyncedVersionAnnotation is an annotation set on PhysicalStatefulCluster CRs by the "core" operator that stores the resource version of the last sync, used for debugging
	PhysicalStatefulClusterCoreLastSyncedVersionAnnotation = physicalStatefulClusterCoreKey + "/last-synced-version"
	//PhysicalStatefulClusterCoreLastSyncedTimestampAnnotation is an annotation set on PhysicalStatefulCluster CRs by the "core" operator that stores the timestamp of the last sync, used for debugging
	PhysicalStatefulClusterCoreLastSyncedTimestampAnnotation = physicalStatefulClusterCoreKey + "/last-synced-ts"
	//PhysicalStatefulClusterNameLabel is a label set on pods by the "core" operator in order to link the pods back to their parent CR
	PhysicalStatefulClusterNameLabel = physicalStatefulClusterCoreKey + "/name"
	//PhysicalStatefulClusterVersionLabel is a label set on pods by the "core" operator in order to verify which operator version last updated the pods
	PhysicalStatefulClusterVersionLabel = physicalStatefulClusterCoreKey + "/version"
	//PhysicalStatefulClusterFinalizer is the finalizer string for the "core" operator
	PhysicalStatefulClusterFinalizer = physicalStatefulClusterCoreKey
	//PhysicalStatefulClusterIgnoreAnnotation is an annotation set on PhysicalStatefulCluster CRs manually to indicate this CR must be ignored.
	PhysicalStatefulClusterIgnoreAnnotation = physicalStatefulClusterCoreKey + "/ignore"
	//PhysicalStatefulClusterIgnoreSyncAnnotation is an annotation set on PhysicalStatefulCluster CRs manually to indicate this CR must not be synced back by cc-sync service.
	PhysicalStatefulClusterIgnoreSyncAnnotation = physicalStatefulClusterCoreKey + "/ignore-sync"
	//PhysicalStatefulClusterRollDelayIntervalInSecondsAnnotation is an annotation set to control the delay between rolling two pods specified in seconds. Default is 0.
	PhysicalStatefulClusterRollDelayIntervalInSecondsAnnotation = physicalStatefulClusterCoreKey + "/roll-delay-interval-seconds"
	//PhysicalStatefulClusterRollPrecheckAnnotation is used to control any component specific checks that get invoked as part of a roll. The values for this key are "enable" / "disable". Default is "enable"
	PhysicalStatefulClusterRollPrecheckAnnotation = physicalStatefulClusterCoreKey + "/roll-precheck"
	//PhysicalStatefulClusterRollPostcheckAnnotation is used to control any component specific checks that get invoked as part of a roll. Valid values are json representation of `PostRollConfig` struct
	PhysicalStatefulClusterRollPostcheckAnnotation = physicalStatefulClusterCoreKey + "/roll-postcheck"
	//PhysicalStatefulClusterResizeComponentHookAnnotation is used to control any component specific hooks that get called as part of the cluster resize lifecycle. The values for this key are "enable" / "disable". Default is "enable"
	PhysicalStatefulClusterResizeComponentHookAnnotation = physicalStatefulClusterCoreKey + "/cluster-resize-component-hook"
	//PhysicalStatefulClusterStorageShrinkAnnotation is used to control any component specific checks that get invoked as part of a storage disk shrink. The values for this key are "enable" / "disable". Default is "disable"
	PhysicalStatefulClusterStorageShrinkAnnotation = physicalStatefulClusterCoreKey + "/storage-shrink"

	//PhysicalStatefulClusterStorageTypeChangeAnnotation is used to control any component specific checks that get invoked as part of storage type change. The values for this key are "enable" / "disable". Default is "disable"
	PhysicalStatefulClusterStorageTypeChangeAnnotation = physicalStatefulClusterCoreKey + "/storage-type-change"

	//PhysicalStatefulClusterCreateInternalServicesAnnotation is used to signal to Operator if each pod should get an internal service. The values for this key are "enable" / "disable". Default is "disable"
	PhysicalStatefulClusterCreateInternalServicesAnnotation = physicalStatefulClusterCoreKey + "/create-internal-services"

	physicalStatefulClusterProxyKey = "physicalstatefulcluster.proxy.confluent.cloud"
	//PhysicalStatefulClusterProxyLastSyncedSpecAnnotation is an annotation set on PhysicalStatefulCluster CRs by the "proxy" operator that stores last synced spec, used to evaluate if spec has changed
	PhysicalStatefulClusterProxyLastSyncedSpecAnnotation = physicalStatefulClusterProxyKey + "/last-synced-spec"
	//PhysicalStatefulClusterProxyLastSyncedVersionAnnotation is an annotation set on PhysicalStatefulCluster CRs by the "proxy" operator that stores the resource version of the last sync, used for debugging
	PhysicalStatefulClusterProxyLastSyncedVersionAnnotation = physicalStatefulClusterProxyKey + "/last-synced-version"
	//PhysicalStatefulClusterProxyLastSyncedTimestampAnnotation is an annotation set on PhysicalStatefulCluster CRs by the "proxy" operator that stores the timestamp of the last sync, used for debugging
	PhysicalStatefulClusterProxyLastSyncedTimestampAnnotation = physicalStatefulClusterProxyKey + "/last-synced-ts"
	//PhysicalStatefulClusterProxyFinalizer is the finalizer string for the "proxy" operator
	PhysicalStatefulClusterProxyFinalizer = physicalStatefulClusterProxyKey

	PhysicalStatefulClusterMetricsAnnotation = "confluent.cloud/metrics"

	//PhysicalStatefulClusterEnvironmentLabel is a label set on psc.ObjectMeta to indicate the cluster enviroment . e.g. "prod"
	PhysicalStatefulClusterEnvironmentLabel = "env"
	//PhysicalStatefulClusterRegionLabel is a label set on psc.ObjectMeta to indicate the cluster region . e.g. "us-west-2"
	PhysicalStatefulClusterRegionLabel = "region"
	//PhysicalStatefulClusterK8SNameLabel is a label set on psc.ObjectMeta to indicate the kubernete cluster long name . e.g. "k8s-sz-a1.us-west-2.aws.internal.stag.cddev.cloud"
	PhysicalStatefulClusterK8SNameLabel = "k8sname"
	//PhysicalStatefulClusterK8SNameLabel is a label set on psc.ObjectMeta to indicate the kubernete cluster id . e.g. "k8s1"
	PhysicalStatefulClusterK8SIdLabel = "k8s_cluster_id"
	//PhysicalStatefulClusterZoneLabel is a label set on AWS Tag to indicate the zone of the pod .
	PhysicalStatefulClusterZoneLabel = "zone"
	//PhysicalStatefulClusterK8SNamespaceLabel is a label set on AWS Tag to indicate the cluster namespace .
	PhysicalStatefulClusterK8SNamespaceLabel = "kube_namespace"
	//PhysicalStatefulClusterLBScopeLabel is a label set on AWS ELB tag to indicate to whether it covers a single pod or multiple pods .
	PhysicalStatefulClusterLBScopeLabel = "lb-scope"
	//PhysicalStatefulClusterLBScopeLabelPerPod is a value set on AWS ELB tag for pod level lb .
	PhysicalStatefulClusterLBScopeLabelPerPod = "per-pod"
	//PhysicalStatefulClusterLBScopeLabelAllPods is a value set on AWS ELB tag for bootstrap lb .
	PhysicalStatefulClusterLBScopeLabelAllPods = "all-pods"
	//PhysicalStatefulClusterIdLabel is a label set on psc.Spec to indicate the physical stateful cluster id.
	PhysicalStatefulClusterIdLabel = "clusterId"

	ConnectFinalizer = physicalStatefulClusterCoreKey + "/connect.finalizer"
)
