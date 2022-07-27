package load

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Nodes struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				Node_Alpha_Kubernetes_Io_Ttl                        string `json:"node.alpha.kubernetes.io/ttl"`
				Volumes_kubernetes_io_controllerManagedAttachDetach string `json:"volumes.kubernetes.io/controller-managed-attach-detach"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				Environment                             string `json:"Environment"`
				Team                                    string `json:"Team"`
				Terraform                               string `json:"Terraform"`
				Beta_Kubernetes_Io_Arch                 string `json:"beta.kubernetes.io/arch"`
				Beta_kubernetes_io_instanceType         string `json:"beta.kubernetes.io/instance-type"`
				Beta_Kubernetes_Io_OS                   string `json:"beta.kubernetes.io/os"`
				Eks_Amazonaws_Com_CapacityType          string `json:"eks.amazonaws.com/capacityType"`
				Eks_Amazonaws_Com_Nodegroup             string `json:"eks.amazonaws.com/nodegroup"`
				Eks_amazonaws_com_nodegroupImage        string `json:"eks.amazonaws.com/nodegroup-image"`
				FailureDomain_beta_kubernetes_io_region string `json:"failure-domain.beta.kubernetes.io/region"`
				FailureDomain_beta_kubernetes_io_zone   string `json:"failure-domain.beta.kubernetes.io/zone"`
				Kubernetes_Io_Arch                      string `json:"kubernetes.io/arch"`
				Kubernetes_Io_Hostname                  string `json:"kubernetes.io/hostname"`
				Kubernetes_Io_OS                        string `json:"kubernetes.io/os"`
				Node_kubernetes_io_instanceType         string `json:"node.kubernetes.io/instance-type"`
				Topology_Kubernetes_Io_Region           string `json:"topology.kubernetes.io/region"`
				Topology_Kubernetes_Io_Zone             string `json:"topology.kubernetes.io/zone"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata struct {
						F_Annotations struct {
							_                                                     *struct{} `json:".,omitempty"`
							F_Node_Alpha_Kubernetes_Io_Ttl                        *struct{} `json:"f:node.alpha.kubernetes.io/ttl,omitempty"`
							F_volumes_kubernetes_io_controllerManagedAttachDetach *struct{} `json:"f:volumes.kubernetes.io/controller-managed-attach-detach,omitempty"`
						} `json:"f:annotations"`
						F_Labels *struct {
							_                                         struct{} `json:"."`
							F_Environment                             struct{} `json:"f:Environment"`
							F_Team                                    struct{} `json:"f:Team"`
							F_Terraform                               struct{} `json:"f:Terraform"`
							F_Beta_Kubernetes_Io_Arch                 struct{} `json:"f:beta.kubernetes.io/arch"`
							F_beta_kubernetes_io_instanceType         struct{} `json:"f:beta.kubernetes.io/instance-type"`
							F_Beta_Kubernetes_Io_OS                   struct{} `json:"f:beta.kubernetes.io/os"`
							F_Eks_Amazonaws_Com_CapacityType          struct{} `json:"f:eks.amazonaws.com/capacityType"`
							F_Eks_Amazonaws_Com_Nodegroup             struct{} `json:"f:eks.amazonaws.com/nodegroup"`
							F_eks_amazonaws_com_nodegroupImage        struct{} `json:"f:eks.amazonaws.com/nodegroup-image"`
							F_failureDomain_beta_kubernetes_io_region struct{} `json:"f:failure-domain.beta.kubernetes.io/region"`
							F_failureDomain_beta_kubernetes_io_zone   struct{} `json:"f:failure-domain.beta.kubernetes.io/zone"`
							F_Kubernetes_Io_Arch                      struct{} `json:"f:kubernetes.io/arch"`
							F_Kubernetes_Io_Hostname                  struct{} `json:"f:kubernetes.io/hostname"`
							F_Kubernetes_Io_OS                        struct{} `json:"f:kubernetes.io/os"`
							F_node_kubernetes_io_instanceType         struct{} `json:"f:node.kubernetes.io/instance-type"`
							F_Topology_Kubernetes_Io_Region           struct{} `json:"f:topology.kubernetes.io/region"`
							F_Topology_Kubernetes_Io_Zone             struct{} `json:"f:topology.kubernetes.io/zone"`
						} `json:"f:labels,omitempty"`
					} `json:"f:metadata"`
					F_Spec *struct {
						F_ProviderID struct{} `json:"f:providerID"`
					} `json:"f:spec,omitempty"`
					F_Status struct {
						F_Allocatable *struct {
							F_ephemeralStorage struct{} `json:"f:ephemeral-storage"`
						} `json:"f:allocatable,omitempty"`
						F_Capacity *struct {
							F_ephemeralStorage struct{} `json:"f:ephemeral-storage"`
						} `json:"f:capacity,omitempty"`
						F_Conditions *struct {
							// "k:{\"type\":\"DiskPressure\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"type\":\"MemoryPressure\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"type\":\"PIDPressure\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"type\":\"Ready\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:conditions,omitempty"`
						F_Images          *struct{} `json:"f:images,omitempty"`
						F_VolumesAttached *struct{} `json:"f:volumesAttached,omitempty"`
						F_VolumesInUse    *struct{} `json:"f:volumesInUse,omitempty"`
					} `json:"f:status"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			ProviderID string `json:"providerID"`
		} `json:"spec"`
		Status struct {
			Addresses []struct {
				Address string `json:"address"`
				Type    string `json:"type"`
			} `json:"addresses"`
			Allocatable struct {
				AttachableVolumesAwsEbs string `json:"attachable-volumes-aws-ebs"`
				Cpu                     string `json:"cpu"`
				EphemeralStorage        string `json:"ephemeral-storage"`
				Hugepages1Gi            string `json:"hugepages-1Gi"`
				Hugepages2Mi            string `json:"hugepages-2Mi"`
				Memory                  string `json:"memory"`
				Pods                    string `json:"pods"`
			} `json:"allocatable"`
			Capacity struct {
				AttachableVolumesAwsEbs string `json:"attachable-volumes-aws-ebs"`
				Cpu                     string `json:"cpu"`
				EphemeralStorage        string `json:"ephemeral-storage"`
				Hugepages1Gi            string `json:"hugepages-1Gi"`
				Hugepages2Mi            string `json:"hugepages-2Mi"`
				Memory                  string `json:"memory"`
				Pods                    string `json:"pods"`
			} `json:"capacity"`
			Conditions []struct {
				LastHeartbeatTime  time.Time `json:"lastHeartbeatTime"`
				LastTransitionTime time.Time `json:"lastTransitionTime"`
				Message            string    `json:"message"`
				Reason             string    `json:"reason"`
				Status             string    `json:"status"`
				Type               string    `json:"type"`
			} `json:"conditions"`
			DaemonEndpoints struct {
				KubeletEndpoint struct {
					Port int `json:"Port"`
				} `json:"kubeletEndpoint"`
			} `json:"daemonEndpoints"`
			Images []struct {
				Names     []string `json:"names"`
				SizeBytes int      `json:"sizeBytes"`
			} `json:"images"`
			NodeInfo struct {
				Architecture            string `json:"architecture"`
				BootID                  string `json:"bootID"`
				ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
				KernelVersion           string `json:"kernelVersion"`
				KubeProxyVersion        string `json:"kubeProxyVersion"`
				KubeletVersion          string `json:"kubeletVersion"`
				MachineID               string `json:"machineID"`
				OperatingSystem         string `json:"operatingSystem"`
				OSImage                 string `json:"osImage"`
				SystemUuid              string `json:"systemUUID"`
			} `json:"nodeInfo"`
			VolumesAttached []struct {
				DevicePath string `json:"devicePath"`
				Name       string `json:"name"`
			} `json:"volumesAttached"`
			VolumesInUse []string `json:"volumesInUse"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type Pods struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				Co_Elastic_Logs_Module          string    `json:"co.elastic.logs/module"`
				Kubernetes_Io_Psp               string    `json:"kubernetes.io/psp"`
				Update_K8S_Elastic_Co_Timestamp time.Time `json:"update.k8s.elastic.co/timestamp,omitempty"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			GenerateName      string    `json:"generateName"`
			Labels            struct {
				Common_K8S_Elastic_Co_Type                             string `json:"common.k8s.elastic.co/type"`
				ControllerRevisionHash                                 string `json:"controller-revision-hash,omitempty"`
				Elasticsearch_k8s_elastic_co_clusterName               string `json:"elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
				Elasticsearch_k8s_elastic_co_configHash                string `json:"elasticsearch.k8s.elastic.co/config-hash,omitempty"`
				Elasticsearch_k8s_elastic_co_httpScheme                string `json:"elasticsearch.k8s.elastic.co/http-scheme,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeData                  string `json:"elasticsearch.k8s.elastic.co/node-data,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeData_cold             string `json:"elasticsearch.k8s.elastic.co/node-data_cold,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeData_content          string `json:"elasticsearch.k8s.elastic.co/node-data_content,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeData_hot              string `json:"elasticsearch.k8s.elastic.co/node-data_hot,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeData_warm             string `json:"elasticsearch.k8s.elastic.co/node-data_warm,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeIngest                string `json:"elasticsearch.k8s.elastic.co/node-ingest,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeMaster                string `json:"elasticsearch.k8s.elastic.co/node-master,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeMl                    string `json:"elasticsearch.k8s.elastic.co/node-ml,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeRemote_cluster_client string `json:"elasticsearch.k8s.elastic.co/node-remote_cluster_client,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeTransform             string `json:"elasticsearch.k8s.elastic.co/node-transform,omitempty"`
				Elasticsearch_k8s_elastic_co_nodeVoting_only           string `json:"elasticsearch.k8s.elastic.co/node-voting_only,omitempty"`
				Elasticsearch_k8s_elastic_co_secureSettingsHash        string `json:"elasticsearch.k8s.elastic.co/secure-settings-hash,omitempty"`
				Elasticsearch_k8s_elastic_co_statefulsetName           string `json:"elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
				Elasticsearch_K8S_Elastic_Co_Version                   string `json:"elasticsearch.k8s.elastic.co/version,omitempty"`
				Kibana_k8s_elastic_co_configChecksum                   string `json:"kibana.k8s.elastic.co/config-checksum,omitempty"`
				Kibana_K8S_Elastic_Co_Name                             string `json:"kibana.k8s.elastic.co/name,omitempty"`
				Kibana_K8S_Elastic_Co_Version                          string `json:"kibana.k8s.elastic.co/version,omitempty"`
				PodTemplateHash                                        string `json:"pod-template-hash,omitempty"`
				Statefulset_kubernetes_io_podName                      string `json:"statefulset.kubernetes.io/pod-name,omitempty"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata *struct {
						F_Annotations struct {
							_                                 *struct{} `json:".,omitempty"`
							F_Co_Elastic_Logs_Module          *struct{} `json:"f:co.elastic.logs/module,omitempty"`
							F_Update_K8S_Elastic_Co_Timestamp *struct{} `json:"f:update.k8s.elastic.co/timestamp,omitempty"`
						} `json:"f:annotations"`
						F_GenerateName *struct{} `json:"f:generateName,omitempty"`
						F_Labels       *struct {
							_                                                        struct{}  `json:"."`
							F_Common_K8S_Elastic_Co_Type                             struct{}  `json:"f:common.k8s.elastic.co/type"`
							F_controllerRevisionHash                                 *struct{} `json:"f:controller-revision-hash,omitempty"`
							F_elasticsearch_k8s_elastic_co_clusterName               *struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
							F_elasticsearch_k8s_elastic_co_configHash                *struct{} `json:"f:elasticsearch.k8s.elastic.co/config-hash,omitempty"`
							F_elasticsearch_k8s_elastic_co_httpScheme                *struct{} `json:"f:elasticsearch.k8s.elastic.co/http-scheme,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeData                  *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeData_cold             *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_cold,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeData_content          *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_content,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeData_hot              *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_hot,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeData_warm             *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_warm,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeIngest                *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-ingest,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeMaster                *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-master,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeMl                    *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-ml,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeRemote_cluster_client *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-remote_cluster_client,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeTransform             *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-transform,omitempty"`
							F_elasticsearch_k8s_elastic_co_nodeVoting_only           *struct{} `json:"f:elasticsearch.k8s.elastic.co/node-voting_only,omitempty"`
							F_elasticsearch_k8s_elastic_co_secureSettingsHash        *struct{} `json:"f:elasticsearch.k8s.elastic.co/secure-settings-hash,omitempty"`
							F_elasticsearch_k8s_elastic_co_statefulsetName           *struct{} `json:"f:elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
							F_Elasticsearch_K8S_Elastic_Co_Version                   *struct{} `json:"f:elasticsearch.k8s.elastic.co/version,omitempty"`
							F_kibana_k8s_elastic_co_configChecksum                   *struct{} `json:"f:kibana.k8s.elastic.co/config-checksum,omitempty"`
							F_Kibana_K8S_Elastic_Co_Name                             *struct{} `json:"f:kibana.k8s.elastic.co/name,omitempty"`
							F_Kibana_K8S_Elastic_Co_Version                          *struct{} `json:"f:kibana.k8s.elastic.co/version,omitempty"`
							F_podTemplateHash                                        *struct{} `json:"f:pod-template-hash,omitempty"`
							F_statefulset_kubernetes_io_podName                      *struct{} `json:"f:statefulset.kubernetes.io/pod-name,omitempty"`
						} `json:"f:labels,omitempty"`
						F_OwnerReferences *struct {
							_ struct{} `json:"."`
							// "k:{\"uid\":\"256808fd-92b4-4b0a-acd1-9989784bdd5d\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"uid\":\"2d752cc2-77ac-42b7-a12a-f44e075eb5e0\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"uid\":\"44d8a127-6a4c-4d2b-a2da-04ccf15a9be7\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"uid\":\"fa62d79b-23e1-4fa7-b27b-129902597736\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:ownerReferences,omitempty"`
					} `json:"f:metadata,omitempty"`
					F_Spec *struct {
						F_Affinity *struct {
							_                 struct{} `json:"."`
							F_PodAntiAffinity struct {
								_                                                 struct{} `json:"."`
								F_PreferredDuringSchedulingIgnoredDuringExecution struct{} `json:"f:preferredDuringSchedulingIgnoredDuringExecution"`
							} `json:"f:podAntiAffinity"`
						} `json:"f:affinity,omitempty"`
						F_AutomountServiceAccountToken struct{} `json:"f:automountServiceAccountToken"`
						F_Containers                   struct {
							// "k:{\"name\":\"elasticsearch\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"kibana\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:containers"`
						F_DnsPolicy          struct{}  `json:"f:dnsPolicy"`
						F_EnableServiceLinks struct{}  `json:"f:enableServiceLinks"`
						F_Hostname           *struct{} `json:"f:hostname,omitempty"`
						F_InitContainers     struct {
							_ struct{} `json:"."`
							// "k:{\"name\":\"elastic-internal-init-config\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-init-filesystem\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-init-keystore\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-suspend\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"install-plugins\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:initContainers"`
						F_RestartPolicy                 struct{}  `json:"f:restartPolicy"`
						F_SchedulerName                 struct{}  `json:"f:schedulerName"`
						F_SecurityContext               struct{}  `json:"f:securityContext"`
						F_Subdomain                     *struct{} `json:"f:subdomain,omitempty"`
						F_TerminationGracePeriodSeconds struct{}  `json:"f:terminationGracePeriodSeconds"`
						F_Volumes                       struct {
							_ struct{} `json:"."`
							// "k:{\"name\":\"downward-api\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-elasticsearch-bin-local\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-elasticsearch-config\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-elasticsearch-config-local\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-elasticsearch-plugins-local\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-http-certificates\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-kibana-config\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-kibana-config-local\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-probe-user\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-remote-certificate-authorities\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-scripts\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-secure-settings\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-transport-certificates\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-unicast-hosts\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elastic-internal-xpack-file-realm\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elasticsearch-certs\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elasticsearch-data\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"elasticsearch-logs\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"name\":\"kibana-data\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:volumes"`
					} `json:"f:spec,omitempty"`
					F_Status *struct {
						F_Conditions struct {
							// "k:{\"type\":\"ContainersReady\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"type\":\"Initialized\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"type\":\"Ready\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:conditions"`
						F_ContainerStatuses     struct{} `json:"f:containerStatuses"`
						F_HostIp                struct{} `json:"f:hostIP"`
						F_InitContainerStatuses struct{} `json:"f:initContainerStatuses"`
						F_Phase                 struct{} `json:"f:phase"`
						F_PodIp                 struct{} `json:"f:podIP"`
						F_PodIPs                struct {
							_ struct{} `json:"."`
							// "k:{\"ip\":\"10.52.1.120\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.1.155\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.1.19\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.1.81\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.2.147\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.2.218\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.2.231\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.3.240\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.3.245\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"ip\":\"10.52.3.95\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:podIPs"`
						F_StartTime struct{} `json:"f:startTime"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			OwnerReferences []struct {
				APIVersion         string `json:"apiVersion"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
				Controller         bool   `json:"controller"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				Uid                string `json:"uid"`
			} `json:"ownerReferences"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Affinity *struct {
				PodAntiAffinity struct {
					PreferredDuringSchedulingIgnoredDuringExecution []struct {
						PodAffinityTerm struct {
							LabelSelector struct {
								MatchLabels struct {
									Elasticsearch_k8s_elastic_co_clusterName string `json:"elasticsearch.k8s.elastic.co/cluster-name"`
								} `json:"matchLabels"`
							} `json:"labelSelector"`
							TopologyKey string `json:"topologyKey"`
						} `json:"podAffinityTerm"`
						Weight int `json:"weight"`
					} `json:"preferredDuringSchedulingIgnoredDuringExecution"`
				} `json:"podAntiAffinity"`
			} `json:"affinity,omitempty"`
			AutomountServiceAccountToken bool `json:"automountServiceAccountToken"`
			Containers                   []struct {
				Env []struct {
					Name      string `json:"name"`
					Value     string `json:"value,omitempty"`
					ValueFrom *struct {
						FieldRef struct {
							APIVersion string `json:"apiVersion"`
							FieldPath  string `json:"fieldPath"`
						} `json:"fieldRef"`
					} `json:"valueFrom,omitempty"`
				} `json:"env,omitempty"`
				Image           string `json:"image"`
				ImagePullPolicy string `json:"imagePullPolicy"`
				Lifecycle       *struct {
					PreStop struct {
						Exec struct {
							Command []string `json:"command"`
						} `json:"exec"`
					} `json:"preStop"`
				} `json:"lifecycle,omitempty"`
				Name  string `json:"name"`
				Ports []struct {
					ContainerPort int    `json:"containerPort"`
					Name          string `json:"name"`
					Protocol      string `json:"protocol"`
				} `json:"ports"`
				ReadinessProbe struct {
					Exec *struct {
						Command []string `json:"command"`
					} `json:"exec,omitempty"`
					FailureThreshold int `json:"failureThreshold"`
					HTTPGet          *struct {
						Path   string `json:"path"`
						Port   int    `json:"port"`
						Scheme string `json:"scheme"`
					} `json:"httpGet,omitempty"`
					InitialDelaySeconds int `json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       int `json:"periodSeconds"`
					SuccessThreshold    int `json:"successThreshold"`
					TimeoutSeconds      int `json:"timeoutSeconds"`
				} `json:"readinessProbe"`
				Resources struct {
					Limits struct {
						Cpu    string `json:"cpu,omitempty"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						Cpu    string `json:"cpu,omitempty"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
				TerminationMessagePath   string `json:"terminationMessagePath"`
				TerminationMessagePolicy string `json:"terminationMessagePolicy"`
				VolumeMounts             []struct {
					MountPath string `json:"mountPath"`
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly,omitempty"`
				} `json:"volumeMounts"`
			} `json:"containers"`
			DnsPolicy          string `json:"dnsPolicy"`
			EnableServiceLinks bool   `json:"enableServiceLinks"`
			Hostname           string `json:"hostname,omitempty"`
			InitContainers     []struct {
				Command []string `json:"command"`
				Env     []struct {
					Name      string `json:"name"`
					Value     string `json:"value,omitempty"`
					ValueFrom *struct {
						FieldRef struct {
							APIVersion string `json:"apiVersion"`
							FieldPath  string `json:"fieldPath"`
						} `json:"fieldRef"`
					} `json:"valueFrom,omitempty"`
				} `json:"env"`
				Image           string `json:"image"`
				ImagePullPolicy string `json:"imagePullPolicy"`
				Name            string `json:"name"`
				Resources       struct {
					Limits struct {
						Cpu    string `json:"cpu,omitempty"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						Cpu    string `json:"cpu,omitempty"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
				SecurityContext *struct {
					Privileged bool `json:"privileged"`
				} `json:"securityContext,omitempty"`
				TerminationMessagePath   string `json:"terminationMessagePath"`
				TerminationMessagePolicy string `json:"terminationMessagePolicy"`
				VolumeMounts             []struct {
					MountPath string `json:"mountPath"`
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly,omitempty"`
				} `json:"volumeMounts"`
			} `json:"initContainers"`
			NodeName                      string   `json:"nodeName"`
			PreemptionPolicy              string   `json:"preemptionPolicy"`
			Priority                      int      `json:"priority"`
			RestartPolicy                 string   `json:"restartPolicy"`
			SchedulerName                 string   `json:"schedulerName"`
			SecurityContext               struct{} `json:"securityContext"`
			ServiceAccount                string   `json:"serviceAccount"`
			ServiceAccountName            string   `json:"serviceAccountName"`
			Subdomain                     string   `json:"subdomain,omitempty"`
			TerminationGracePeriodSeconds int      `json:"terminationGracePeriodSeconds"`
			Tolerations                   []struct {
				Effect            string `json:"effect"`
				Key               string `json:"key"`
				Operator          string `json:"operator"`
				TolerationSeconds int    `json:"tolerationSeconds"`
			} `json:"tolerations"`
			Volumes []struct {
				ConfigMap *struct {
					DefaultMode int    `json:"defaultMode"`
					Name        string `json:"name"`
					Optional    bool   `json:"optional"`
				} `json:"configMap,omitempty"`
				DownwardAPI *struct {
					DefaultMode int `json:"defaultMode"`
					Items       []struct {
						FieldRef struct {
							APIVersion string `json:"apiVersion"`
							FieldPath  string `json:"fieldPath"`
						} `json:"fieldRef"`
						Path string `json:"path"`
					} `json:"items"`
				} `json:"downwardAPI,omitempty"`
				EmptyDir              *struct{} `json:"emptyDir,omitempty"`
				Name                  string    `json:"name"`
				PersistentVolumeClaim *struct {
					ClaimName string `json:"claimName"`
				} `json:"persistentVolumeClaim,omitempty"`
				Secret *struct {
					DefaultMode int `json:"defaultMode"`
					Items       []struct {
						Key  string `json:"key"`
						Path string `json:"path"`
					} `json:"items,omitempty"`
					Optional   bool   `json:"optional"`
					SecretName string `json:"secretName"`
				} `json:"secret,omitempty"`
			} `json:"volumes"`
		} `json:"spec"`
		Status struct {
			Conditions []struct {
				LastProbeTime      interface{} `json:"lastProbeTime"`
				LastTransitionTime time.Time   `json:"lastTransitionTime"`
				Status             string      `json:"status"`
				Type               string      `json:"type"`
			} `json:"conditions"`
			ContainerStatuses []struct {
				ContainerID  string   `json:"containerID"`
				Image        string   `json:"image"`
				ImageID      string   `json:"imageID"`
				LastState    struct{} `json:"lastState"`
				Name         string   `json:"name"`
				Ready        bool     `json:"ready"`
				RestartCount int      `json:"restartCount"`
				Started      bool     `json:"started"`
				State        struct {
					Running struct {
						StartedAt time.Time `json:"startedAt"`
					} `json:"running"`
				} `json:"state"`
			} `json:"containerStatuses"`
			HostIp                string `json:"hostIP"`
			InitContainerStatuses []struct {
				ContainerID  string   `json:"containerID"`
				Image        string   `json:"image"`
				ImageID      string   `json:"imageID"`
				LastState    struct{} `json:"lastState"`
				Name         string   `json:"name"`
				Ready        bool     `json:"ready"`
				RestartCount int      `json:"restartCount"`
				State        struct {
					Terminated struct {
						ContainerID string    `json:"containerID"`
						ExitCode    int       `json:"exitCode"`
						FinishedAt  time.Time `json:"finishedAt"`
						Reason      string    `json:"reason"`
						StartedAt   time.Time `json:"startedAt"`
					} `json:"terminated"`
				} `json:"state"`
			} `json:"initContainerStatuses"`
			Phase  string `json:"phase"`
			PodIp  string `json:"podIP"`
			PodIPs []struct {
				Ip string `json:"ip"`
			} `json:"podIPs"`
			QosClass  string    `json:"qosClass"`
			StartTime time.Time `json:"startTime"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type ES struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				Common_k8s_elastic_co_controllerVersion        string `json:"common.k8s.elastic.co/controller-version"`
				Elasticsearch_k8s_elastic_co_clusterUuid       string `json:"elasticsearch.k8s.elastic.co/cluster-uuid"`
				Kubectl_kubernetes_io_lastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Generation        int       `json:"generation"`
			ManagedFields     []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata struct {
						F_Annotations struct {
							_                                                *struct{} `json:".,omitempty"`
							F_common_k8s_elastic_co_controllerVersion        *struct{} `json:"f:common.k8s.elastic.co/controller-version,omitempty"`
							F_elasticsearch_k8s_elastic_co_clusterUuid       *struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-uuid,omitempty"`
							F_kubectl_kubernetes_io_lastAppliedConfiguration *struct{} `json:"f:kubectl.kubernetes.io/last-applied-configuration,omitempty"`
						} `json:"f:annotations"`
					} `json:"f:metadata"`
					F_Spec struct {
						_      *struct{} `json:".,omitempty"`
						F_Auth *struct{} `json:"f:auth,omitempty"`
						F_HTTP struct {
							_         *struct{} `json:".,omitempty"`
							F_Service *struct {
								_          struct{} `json:"."`
								F_Metadata struct {
									_             struct{} `json:"."`
									F_Annotations struct {
										_                                                    struct{} `json:"."`
										F_service_beta_kubernetes_io_awsLoadBalancerInternal struct{} `json:"f:service.beta.kubernetes.io/aws-load-balancer-internal"`
										F_service_beta_kubernetes_io_awsLoadBalancerType     struct{} `json:"f:service.beta.kubernetes.io/aws-load-balancer-type"`
									} `json:"f:annotations"`
								} `json:"f:metadata"`
								F_Spec struct {
									_      struct{} `json:"."`
									F_Type struct{} `json:"f:type"`
								} `json:"f:spec"`
							} `json:"f:service,omitempty"`
							F_Tls *struct {
								_             struct{} `json:"."`
								F_Certificate struct{} `json:"f:certificate"`
							} `json:"f:tls,omitempty"`
						} `json:"f:http"`
						F_Monitoring *struct {
							_         struct{} `json:"."`
							F_Logs    struct{} `json:"f:logs"`
							F_Metrics struct{} `json:"f:metrics"`
						} `json:"f:monitoring,omitempty"`
						F_NodeSets       *struct{} `json:"f:nodeSets,omitempty"`
						F_SecureSettings *struct{} `json:"f:secureSettings,omitempty"`
						F_Transport      *struct {
							_         struct{} `json:"."`
							F_Service struct {
								_          struct{} `json:"."`
								F_Metadata struct{} `json:"f:metadata"`
								F_Spec     struct{} `json:"f:spec"`
							} `json:"f:service"`
							F_Tls struct {
								_             struct{} `json:"."`
								F_Certificate struct{} `json:"f:certificate"`
							} `json:"f:tls"`
						} `json:"f:transport,omitempty"`
						F_UpdateStrategy *struct {
							_              struct{} `json:"."`
							F_ChangeBudget struct{} `json:"f:changeBudget"`
						} `json:"f:updateStrategy,omitempty"`
						F_Version *struct{} `json:"f:version,omitempty"`
					} `json:"f:spec"`
					F_Status *struct {
						_                struct{} `json:"."`
						F_AvailableNodes struct{} `json:"f:availableNodes"`
						F_Health         struct{} `json:"f:health"`
						F_Phase          struct{} `json:"f:phase"`
						F_Version        struct{} `json:"f:version"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Auth struct{} `json:"auth"`
			HTTP struct {
				Service struct {
					Metadata struct {
						Annotations struct {
							Service_beta_kubernetes_io_awsLoadBalancerInternal string `json:"service.beta.kubernetes.io/aws-load-balancer-internal"`
							Service_beta_kubernetes_io_awsLoadBalancerType     string `json:"service.beta.kubernetes.io/aws-load-balancer-type"`
						} `json:"annotations"`
					} `json:"metadata"`
					Spec struct {
						Type string `json:"type"`
					} `json:"spec"`
				} `json:"service"`
				Tls struct {
					Certificate struct{} `json:"certificate"`
				} `json:"tls"`
			} `json:"http"`
			Monitoring struct {
				Logs    struct{} `json:"logs"`
				Metrics struct{} `json:"metrics"`
			} `json:"monitoring"`
			NodeSets []struct {
				Config struct {
					Node_Roles                  []string `json:"node.roles"`
					Xpack_Ml_Enabled            bool     `json:"xpack.ml.enabled,omitempty"`
					Xpack_Security_Authc_Realms struct {
						File_File1 struct {
							Order int `json:"order"`
						} `json:"file.file1"`
						Native_Native1 struct {
							Order int `json:"order"`
						} `json:"native.native1"`
						Oidc_Oidc1 struct {
							Claims_Groups            string `json:"claims.groups"`
							Claims_Principal         string `json:"claims.principal"`
							Op_authorizationEndpoint string `json:"op.authorization_endpoint"`
							Op_endsessionEndpoint    string `json:"op.endsession_endpoint"`
							Op_Issuer                string `json:"op.issuer"`
							Op_jwksetPath            string `json:"op.jwkset_path"`
							Op_tokenEndpoint         string `json:"op.token_endpoint"`
							Op_userinfoEndpoint      string `json:"op.userinfo_endpoint"`
							Order                    int    `json:"order"`
							Rp_clientID              string `json:"rp.client_id"`
							Rp_postLogoutRedirectURI string `json:"rp.post_logout_redirect_uri"`
							Rp_redirectURI           string `json:"rp.redirect_uri"`
							Rp_responseType          string `json:"rp.response_type"`
						} `json:"oidc.oidc1"`
					} `json:"xpack.security.authc.realms"`
					Xpack_Security_Authc_Token_Enabled           bool   `json:"xpack.security.authc.token.enabled"`
					Xpack_Security_Enabled                       bool   `json:"xpack.security.enabled"`
					Xpack_security_http_ssl_clientAuthentication string `json:"xpack.security.http.ssl.client_authentication"`
					Xpack_Security_HTTP_Ssl_Enabled              bool   `json:"xpack.security.http.ssl.enabled"`
					Xpack_Security_Transport_Ssl_Enabled         bool   `json:"xpack.security.transport.ssl.enabled"`
				} `json:"config"`
				Count       int    `json:"count"`
				Name        string `json:"name"`
				PodTemplate struct {
					Spec struct {
						Containers []struct {
							Env []struct {
								Name  string `json:"name"`
								Value string `json:"value"`
							} `json:"env"`
							Name      string `json:"name"`
							Resources struct {
								Limits struct {
									Cpu    int    `json:"cpu"`
									Memory string `json:"memory"`
								} `json:"limits"`
								Requests struct {
									Cpu    int    `json:"cpu"`
									Memory string `json:"memory"`
								} `json:"requests"`
							} `json:"resources"`
						} `json:"containers,omitempty"`
						InitContainers []struct {
							Command []string `json:"command"`
							Name    string   `json:"name"`
						} `json:"initContainers"`
					} `json:"spec"`
				} `json:"podTemplate"`
				VolumeClaimTemplates []struct {
					Metadata struct {
						Name string `json:"name"`
					} `json:"metadata"`
					Spec struct {
						AccessModes []string `json:"accessModes"`
						Resources   struct {
							Requests struct {
								Storage string `json:"storage"`
							} `json:"requests"`
						} `json:"resources"`
					} `json:"spec"`
				} `json:"volumeClaimTemplates"`
			} `json:"nodeSets"`
			SecureSettings []struct {
				SecretName string `json:"secretName"`
			} `json:"secureSettings"`
			Transport struct {
				Service struct {
					Metadata struct{} `json:"metadata"`
					Spec     struct{} `json:"spec"`
				} `json:"service"`
				Tls struct {
					Certificate struct{} `json:"certificate"`
				} `json:"tls"`
			} `json:"transport"`
			UpdateStrategy struct {
				ChangeBudget struct{} `json:"changeBudget"`
			} `json:"updateStrategy"`
			Version string `json:"version"`
		} `json:"spec"`
		Status struct {
			AvailableNodes int    `json:"availableNodes"`
			Health         string `json:"health"`
			Phase          string `json:"phase"`
			Version        string `json:"version"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type Events struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion     string      `json:"apiVersion"`
		Count          int         `json:"count"`
		EventTime      interface{} `json:"eventTime"`
		FirstTimestamp time.Time   `json:"firstTimestamp"`
		InvolvedObject struct {
			APIVersion      string `json:"apiVersion"`
			FieldPath       string `json:"fieldPath,omitempty"`
			Kind            string `json:"kind"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"involvedObject"`
		Kind          string    `json:"kind"`
		LastTimestamp time.Time `json:"lastTimestamp"`
		Message       string    `json:"message"`
		Metadata      struct {
			CreationTimestamp time.Time `json:"creationTimestamp"`
			ManagedFields     []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Count          struct{} `json:"f:count"`
					F_FirstTimestamp struct{} `json:"f:firstTimestamp"`
					F_InvolvedObject struct {
						F_APIVersion      struct{}  `json:"f:apiVersion"`
						F_FieldPath       *struct{} `json:"f:fieldPath,omitempty"`
						F_Kind            struct{}  `json:"f:kind"`
						F_Name            struct{}  `json:"f:name"`
						F_Namespace       struct{}  `json:"f:namespace"`
						F_ResourceVersion struct{}  `json:"f:resourceVersion"`
						F_Uid             struct{}  `json:"f:uid"`
					} `json:"f:involvedObject"`
					F_LastTimestamp struct{} `json:"f:lastTimestamp"`
					F_Message       struct{} `json:"f:message"`
					F_Reason        struct{} `json:"f:reason"`
					F_Source        struct {
						F_Component struct{}  `json:"f:component"`
						F_Host      *struct{} `json:"f:host,omitempty"`
					} `json:"f:source"`
					F_Type struct{} `json:"f:type"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Reason             string `json:"reason"`
		ReportingComponent string `json:"reportingComponent"`
		ReportingInstance  string `json:"reportingInstance"`
		Source             struct {
			Component string `json:"component"`
			Host      string `json:"host,omitempty"`
		} `json:"source"`
		Type string `json:"type"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type Kibana struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				Association_k8s_elastic_co_esConf              string `json:"association.k8s.elastic.co/es-conf"`
				Common_k8s_elastic_co_controllerVersion        string `json:"common.k8s.elastic.co/controller-version"`
				Kubectl_kubernetes_io_lastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Generation        int       `json:"generation"`
			ManagedFields     []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata struct {
						F_Annotations struct {
							_                                                *struct{} `json:".,omitempty"`
							F_association_k8s_elastic_co_esConf              *struct{} `json:"f:association.k8s.elastic.co/es-conf,omitempty"`
							F_common_k8s_elastic_co_controllerVersion        *struct{} `json:"f:common.k8s.elastic.co/controller-version,omitempty"`
							F_kubectl_kubernetes_io_lastAppliedConfiguration *struct{} `json:"f:kubectl.kubernetes.io/last-applied-configuration,omitempty"`
						} `json:"f:annotations"`
					} `json:"f:metadata"`
					F_Spec struct {
						_        *struct{} `json:".,omitempty"`
						F_Config *struct {
							_                                       struct{} `json:"."`
							F_Elasticsearch_Password                struct{} `json:"f:elasticsearch.password"`
							F_Elasticsearch_Username                struct{} `json:"f:elasticsearch.username"`
							F_Logging_Verbose                       struct{} `json:"f:logging.verbose"`
							F_Server                                struct{} `json:"f:server"`
							F_Xpack_Security_Authc_Providers        struct{} `json:"f:xpack.security.authc.providers"`
							F_Xpack_Security_Authc_Selector_Enabled struct{} `json:"f:xpack.security.authc.selector.enabled"`
						} `json:"f:config,omitempty"`
						F_Count            *struct{} `json:"f:count,omitempty"`
						F_ElasticsearchRef *struct {
							_      struct{} `json:"."`
							F_Name struct{} `json:"f:name"`
						} `json:"f:elasticsearchRef,omitempty"`
						F_EnterpriseSearchRef *struct {
							_      struct{} `json:"."`
							F_Name struct{} `json:"f:name"`
						} `json:"f:enterpriseSearchRef,omitempty"`
						F_HTTP *struct {
							_         struct{} `json:"."`
							F_Service struct {
								_          struct{} `json:"."`
								F_Metadata struct {
									_             struct{} `json:"."`
									F_Annotations struct {
										_                                                                  struct{} `json:"."`
										F_service_beta_kubernetes_io_awsLoadBalancerAdditionalResourceTags struct{} `json:"f:service.beta.kubernetes.io/aws-load-balancer-additional-resource-tags"`
										F_service_beta_kubernetes_io_awsLoadBalancerInternal               struct{} `json:"f:service.beta.kubernetes.io/aws-load-balancer-internal"`
										F_service_beta_kubernetes_io_awsLoadBalancerType                   struct{} `json:"f:service.beta.kubernetes.io/aws-load-balancer-type"`
									} `json:"f:annotations"`
								} `json:"f:metadata"`
								F_Spec struct {
									_      struct{} `json:"."`
									F_Type struct{} `json:"f:type"`
								} `json:"f:spec"`
							} `json:"f:service"`
							F_Tls struct {
								_                       struct{} `json:"."`
								F_SelfSignedCertificate struct {
									_          struct{} `json:"."`
									F_Disabled struct{} `json:"f:disabled"`
								} `json:"f:selfSignedCertificate"`
							} `json:"f:tls"`
						} `json:"f:http,omitempty"`
						F_Monitoring *struct {
							_         struct{} `json:"."`
							F_Logs    struct{} `json:"f:logs"`
							F_Metrics struct{} `json:"f:metrics"`
						} `json:"f:monitoring,omitempty"`
						F_PodTemplate *struct {
							_      struct{} `json:"."`
							F_Spec struct{} `json:"f:spec"`
						} `json:"f:podTemplate,omitempty"`
						F_Version *struct{} `json:"f:version,omitempty"`
					} `json:"f:spec"`
					F_Status *struct {
						_                                struct{} `json:"."`
						F_AssociationStatus              struct{} `json:"f:associationStatus"`
						F_AvailableNodes                 struct{} `json:"f:availableNodes"`
						F_Count                          struct{} `json:"f:count"`
						F_ElasticsearchAssociationStatus struct{} `json:"f:elasticsearchAssociationStatus"`
						F_Health                         struct{} `json:"f:health"`
						F_Selector                       struct{} `json:"f:selector"`
						F_Version                        struct{} `json:"f:version"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Config struct {
				Elasticsearch_Password string `json:"elasticsearch.password"`
				Elasticsearch_Username string `json:"elasticsearch.username"`
				Logging_Verbose        bool   `json:"logging.verbose"`
				Server                 struct {
					BasePath        string `json:"basePath"`
					PublicBaseURL   string `json:"publicBaseUrl"`
					RewriteBasePath bool   `json:"rewriteBasePath"`
				} `json:"server"`
				Xpack_Security_Authc_Providers struct {
					Basic_Basic1 struct {
						Order int `json:"order"`
					} `json:"basic.basic1"`
					Oidc_Oidc1 struct {
						Order int    `json:"order"`
						Realm string `json:"realm"`
					} `json:"oidc.oidc1"`
				} `json:"xpack.security.authc.providers"`
				Xpack_Security_Authc_Selector_Enabled bool `json:"xpack.security.authc.selector.enabled"`
			} `json:"config"`
			Count            int `json:"count"`
			ElasticsearchRef struct {
				Name string `json:"name"`
			} `json:"elasticsearchRef"`
			EnterpriseSearchRef struct {
				Name string `json:"name"`
			} `json:"enterpriseSearchRef"`
			HTTP struct {
				Service struct {
					Metadata struct {
						Annotations struct {
							Service_beta_kubernetes_io_awsLoadBalancerAdditionalResourceTags string `json:"service.beta.kubernetes.io/aws-load-balancer-additional-resource-tags"`
							Service_beta_kubernetes_io_awsLoadBalancerInternal               string `json:"service.beta.kubernetes.io/aws-load-balancer-internal"`
							Service_beta_kubernetes_io_awsLoadBalancerType                   string `json:"service.beta.kubernetes.io/aws-load-balancer-type"`
						} `json:"annotations"`
					} `json:"metadata"`
					Spec struct {
						Type string `json:"type"`
					} `json:"spec"`
				} `json:"service"`
				Tls struct {
					SelfSignedCertificate struct {
						Disabled bool `json:"disabled"`
					} `json:"selfSignedCertificate"`
				} `json:"tls"`
			} `json:"http"`
			Monitoring struct {
				Logs    struct{} `json:"logs"`
				Metrics struct{} `json:"metrics"`
			} `json:"monitoring"`
			PodTemplate struct {
				Spec struct {
					Containers []struct {
						Name           string `json:"name"`
						ReadinessProbe struct {
							HTTPGet struct {
								Path   string `json:"path"`
								Port   int    `json:"port"`
								Scheme string `json:"scheme"`
							} `json:"httpGet"`
						} `json:"readinessProbe"`
					} `json:"containers"`
				} `json:"spec"`
			} `json:"podTemplate"`
			Version string `json:"version"`
		} `json:"spec"`
		Status struct {
			AssociationStatus              string `json:"associationStatus"`
			AvailableNodes                 int    `json:"availableNodes"`
			Count                          int    `json:"count"`
			ElasticsearchAssociationStatus string `json:"elasticsearchAssociationStatus"`
			Health                         string `json:"health"`
			Selector                       string `json:"selector"`
			Version                        string `json:"version"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type DiagVersion struct {
	DiagnosticsVersion struct {
		BuildDate time.Time `json:"BuildDate"`
		Hash      string    `json:"Hash"`
		Version   string    `json:"Version"`
	} `json:"DiagnosticsVersion"`
	ServerVersion struct {
		BuildDate    time.Time `json:"buildDate"`
		Compiler     string    `json:"compiler"`
		GitCommit    string    `json:"gitCommit"`
		GitTreeState string    `json:"gitTreeState"`
		GitVersion   string    `json:"gitVersion"`
		GoVersion    string    `json:"goVersion"`
		Major        string    `json:"major"`
		Minor        string    `json:"minor"`
		Platform     string    `json:"platform"`
	} `json:"ServerVersion"`
}

type StatefulSet struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Generation        int       `json:"generation"`
			Labels            struct {
				Common_k8s_elastic_co_templateHash           string `json:"common.k8s.elastic.co/template-hash"`
				Common_K8S_Elastic_Co_Type                   string `json:"common.k8s.elastic.co/type"`
				Elasticsearch_k8s_elastic_co_clusterName     string `json:"elasticsearch.k8s.elastic.co/cluster-name"`
				Elasticsearch_k8s_elastic_co_statefulsetName string `json:"elasticsearch.k8s.elastic.co/statefulset-name"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata *struct {
						F_Labels struct {
							_                                              struct{} `json:"."`
							F_common_k8s_elastic_co_templateHash           struct{} `json:"f:common.k8s.elastic.co/template-hash"`
							F_Common_K8S_Elastic_Co_Type                   struct{} `json:"f:common.k8s.elastic.co/type"`
							F_elasticsearch_k8s_elastic_co_clusterName     struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-name"`
							F_elasticsearch_k8s_elastic_co_statefulsetName struct{} `json:"f:elasticsearch.k8s.elastic.co/statefulset-name"`
						} `json:"f:labels"`
						F_OwnerReferences struct {
							_ struct{} `json:"."`
							// "k:{\"uid\":\"5de07191-561b-4b60-b3de-ab58726c31d4\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:ownerReferences"`
					} `json:"f:metadata,omitempty"`
					F_Spec *struct {
						F_PodManagementPolicy  struct{} `json:"f:podManagementPolicy"`
						F_Replicas             struct{} `json:"f:replicas"`
						F_RevisionHistoryLimit struct{} `json:"f:revisionHistoryLimit"`
						F_Selector             struct{} `json:"f:selector"`
						F_ServiceName          struct{} `json:"f:serviceName"`
						F_Template             struct {
							F_Metadata struct {
								F_Annotations struct {
									_                        struct{} `json:"."`
									F_Co_Elastic_Logs_Module struct{} `json:"f:co.elastic.logs/module"`
								} `json:"f:annotations"`
								F_Labels struct {
									_                                                        struct{} `json:"."`
									F_Common_K8S_Elastic_Co_Type                             struct{} `json:"f:common.k8s.elastic.co/type"`
									F_elasticsearch_k8s_elastic_co_clusterName               struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-name"`
									F_elasticsearch_k8s_elastic_co_configHash                struct{} `json:"f:elasticsearch.k8s.elastic.co/config-hash"`
									F_elasticsearch_k8s_elastic_co_httpScheme                struct{} `json:"f:elasticsearch.k8s.elastic.co/http-scheme"`
									F_elasticsearch_k8s_elastic_co_nodeData                  struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data"`
									F_elasticsearch_k8s_elastic_co_nodeData_cold             struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_cold"`
									F_elasticsearch_k8s_elastic_co_nodeData_content          struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_content"`
									F_elasticsearch_k8s_elastic_co_nodeData_hot              struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_hot"`
									F_elasticsearch_k8s_elastic_co_nodeData_warm             struct{} `json:"f:elasticsearch.k8s.elastic.co/node-data_warm"`
									F_elasticsearch_k8s_elastic_co_nodeIngest                struct{} `json:"f:elasticsearch.k8s.elastic.co/node-ingest"`
									F_elasticsearch_k8s_elastic_co_nodeMaster                struct{} `json:"f:elasticsearch.k8s.elastic.co/node-master"`
									F_elasticsearch_k8s_elastic_co_nodeMl                    struct{} `json:"f:elasticsearch.k8s.elastic.co/node-ml"`
									F_elasticsearch_k8s_elastic_co_nodeRemote_cluster_client struct{} `json:"f:elasticsearch.k8s.elastic.co/node-remote_cluster_client"`
									F_elasticsearch_k8s_elastic_co_nodeTransform             struct{} `json:"f:elasticsearch.k8s.elastic.co/node-transform"`
									F_elasticsearch_k8s_elastic_co_nodeVoting_only           struct{} `json:"f:elasticsearch.k8s.elastic.co/node-voting_only"`
									F_elasticsearch_k8s_elastic_co_secureSettingsHash        struct{} `json:"f:elasticsearch.k8s.elastic.co/secure-settings-hash"`
									F_elasticsearch_k8s_elastic_co_statefulsetName           struct{} `json:"f:elasticsearch.k8s.elastic.co/statefulset-name"`
									F_Elasticsearch_K8S_Elastic_Co_Version                   struct{} `json:"f:elasticsearch.k8s.elastic.co/version"`
								} `json:"f:labels"`
							} `json:"f:metadata"`
							F_Spec struct {
								F_Affinity struct {
									_                 struct{} `json:"."`
									F_PodAntiAffinity struct {
										_                                                 struct{} `json:"."`
										F_PreferredDuringSchedulingIgnoredDuringExecution struct{} `json:"f:preferredDuringSchedulingIgnoredDuringExecution"`
									} `json:"f:podAntiAffinity"`
								} `json:"f:affinity"`
								F_AutomountServiceAccountToken struct{} `json:"f:automountServiceAccountToken"`
								F_Containers                   struct {
									// "k:{\"name\":\"elasticsearch\"}" cannot be unmarshalled into a struct field by encoding/json.
								} `json:"f:containers"`
								F_DnsPolicy      struct{} `json:"f:dnsPolicy"`
								F_InitContainers struct {
									_ struct{} `json:"."`
									// "k:{\"name\":\"elastic-internal-init-filesystem\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-init-keystore\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-suspend\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"install-plugins\"}" cannot be unmarshalled into a struct field by encoding/json.
								} `json:"f:initContainers"`
								F_RestartPolicy                 struct{} `json:"f:restartPolicy"`
								F_SchedulerName                 struct{} `json:"f:schedulerName"`
								F_SecurityContext               struct{} `json:"f:securityContext"`
								F_TerminationGracePeriodSeconds struct{} `json:"f:terminationGracePeriodSeconds"`
								F_Volumes                       struct {
									_ struct{} `json:"."`
									// "k:{\"name\":\"downward-api\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-elasticsearch-bin-local\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-elasticsearch-config\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-elasticsearch-config-local\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-elasticsearch-plugins-local\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-http-certificates\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-probe-user\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-remote-certificate-authorities\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-scripts\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-secure-settings\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-transport-certificates\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-unicast-hosts\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-xpack-file-realm\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elasticsearch-data\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elasticsearch-logs\"}" cannot be unmarshalled into a struct field by encoding/json.
								} `json:"f:volumes"`
							} `json:"f:spec"`
						} `json:"f:template"`
						F_UpdateStrategy struct {
							F_Type struct{} `json:"f:type"`
						} `json:"f:updateStrategy"`
						F_VolumeClaimTemplates struct{} `json:"f:volumeClaimTemplates"`
					} `json:"f:spec,omitempty"`
					F_Status *struct {
						F_CollisionCount     struct{} `json:"f:collisionCount"`
						F_CurrentRevision    struct{} `json:"f:currentRevision"`
						F_ObservedGeneration struct{} `json:"f:observedGeneration"`
						F_ReadyReplicas      struct{} `json:"f:readyReplicas"`
						F_Replicas           struct{} `json:"f:replicas"`
						F_UpdateRevision     struct{} `json:"f:updateRevision"`
						F_UpdatedReplicas    struct{} `json:"f:updatedReplicas"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			OwnerReferences []struct {
				APIVersion         string `json:"apiVersion"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
				Controller         bool   `json:"controller"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				Uid                string `json:"uid"`
			} `json:"ownerReferences"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			PodManagementPolicy  string `json:"podManagementPolicy"`
			Replicas             int    `json:"replicas"`
			RevisionHistoryLimit int    `json:"revisionHistoryLimit"`
			Selector             struct {
				MatchLabels struct {
					Common_K8S_Elastic_Co_Type                   string `json:"common.k8s.elastic.co/type"`
					Elasticsearch_k8s_elastic_co_clusterName     string `json:"elasticsearch.k8s.elastic.co/cluster-name"`
					Elasticsearch_k8s_elastic_co_statefulsetName string `json:"elasticsearch.k8s.elastic.co/statefulset-name"`
				} `json:"matchLabels"`
			} `json:"selector"`
			ServiceName string `json:"serviceName"`
			Template    struct {
				Metadata struct {
					Annotations struct {
						Co_Elastic_Logs_Module string `json:"co.elastic.logs/module"`
					} `json:"annotations"`
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						Common_K8S_Elastic_Co_Type                             string `json:"common.k8s.elastic.co/type"`
						Elasticsearch_k8s_elastic_co_clusterName               string `json:"elasticsearch.k8s.elastic.co/cluster-name"`
						Elasticsearch_k8s_elastic_co_configHash                string `json:"elasticsearch.k8s.elastic.co/config-hash"`
						Elasticsearch_k8s_elastic_co_httpScheme                string `json:"elasticsearch.k8s.elastic.co/http-scheme"`
						Elasticsearch_k8s_elastic_co_nodeData                  string `json:"elasticsearch.k8s.elastic.co/node-data"`
						Elasticsearch_k8s_elastic_co_nodeData_cold             string `json:"elasticsearch.k8s.elastic.co/node-data_cold"`
						Elasticsearch_k8s_elastic_co_nodeData_content          string `json:"elasticsearch.k8s.elastic.co/node-data_content"`
						Elasticsearch_k8s_elastic_co_nodeData_hot              string `json:"elasticsearch.k8s.elastic.co/node-data_hot"`
						Elasticsearch_k8s_elastic_co_nodeData_warm             string `json:"elasticsearch.k8s.elastic.co/node-data_warm"`
						Elasticsearch_k8s_elastic_co_nodeIngest                string `json:"elasticsearch.k8s.elastic.co/node-ingest"`
						Elasticsearch_k8s_elastic_co_nodeMaster                string `json:"elasticsearch.k8s.elastic.co/node-master"`
						Elasticsearch_k8s_elastic_co_nodeMl                    string `json:"elasticsearch.k8s.elastic.co/node-ml"`
						Elasticsearch_k8s_elastic_co_nodeRemote_cluster_client string `json:"elasticsearch.k8s.elastic.co/node-remote_cluster_client"`
						Elasticsearch_k8s_elastic_co_nodeTransform             string `json:"elasticsearch.k8s.elastic.co/node-transform"`
						Elasticsearch_k8s_elastic_co_nodeVoting_only           string `json:"elasticsearch.k8s.elastic.co/node-voting_only"`
						Elasticsearch_k8s_elastic_co_secureSettingsHash        string `json:"elasticsearch.k8s.elastic.co/secure-settings-hash"`
						Elasticsearch_k8s_elastic_co_statefulsetName           string `json:"elasticsearch.k8s.elastic.co/statefulset-name"`
						Elasticsearch_K8S_Elastic_Co_Version                   string `json:"elasticsearch.k8s.elastic.co/version"`
					} `json:"labels"`
				} `json:"metadata"`
				Spec struct {
					Affinity struct {
						PodAntiAffinity struct {
							PreferredDuringSchedulingIgnoredDuringExecution []struct {
								PodAffinityTerm struct {
									LabelSelector struct {
										MatchLabels struct {
											Elasticsearch_k8s_elastic_co_clusterName string `json:"elasticsearch.k8s.elastic.co/cluster-name"`
										} `json:"matchLabels"`
									} `json:"labelSelector"`
									TopologyKey string `json:"topologyKey"`
								} `json:"podAffinityTerm"`
								Weight int `json:"weight"`
							} `json:"preferredDuringSchedulingIgnoredDuringExecution"`
						} `json:"podAntiAffinity"`
					} `json:"affinity"`
					AutomountServiceAccountToken bool `json:"automountServiceAccountToken"`
					Containers                   []struct {
						Env []struct {
							Name      string `json:"name"`
							Value     string `json:"value,omitempty"`
							ValueFrom *struct {
								FieldRef struct {
									APIVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
							} `json:"valueFrom,omitempty"`
						} `json:"env"`
						Image           string `json:"image"`
						ImagePullPolicy string `json:"imagePullPolicy"`
						Lifecycle       struct {
							PreStop struct {
								Exec struct {
									Command []string `json:"command"`
								} `json:"exec"`
							} `json:"preStop"`
						} `json:"lifecycle"`
						Name  string `json:"name"`
						Ports []struct {
							ContainerPort int    `json:"containerPort"`
							Name          string `json:"name"`
							Protocol      string `json:"protocol"`
						} `json:"ports"`
						ReadinessProbe struct {
							Exec struct {
								Command []string `json:"command"`
							} `json:"exec"`
							FailureThreshold    int `json:"failureThreshold"`
							InitialDelaySeconds int `json:"initialDelaySeconds"`
							PeriodSeconds       int `json:"periodSeconds"`
							SuccessThreshold    int `json:"successThreshold"`
							TimeoutSeconds      int `json:"timeoutSeconds"`
						} `json:"readinessProbe"`
						Resources struct {
							Limits struct {
								Cpu    string `json:"cpu,omitempty"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								Cpu    string `json:"cpu,omitempty"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						TerminationMessagePath   string `json:"terminationMessagePath"`
						TerminationMessagePolicy string `json:"terminationMessagePolicy"`
						VolumeMounts             []struct {
							MountPath string `json:"mountPath"`
							Name      string `json:"name"`
							ReadOnly  bool   `json:"readOnly,omitempty"`
						} `json:"volumeMounts"`
					} `json:"containers"`
					DnsPolicy      string `json:"dnsPolicy"`
					InitContainers []struct {
						Command []string `json:"command"`
						Env     []struct {
							Name      string `json:"name"`
							Value     string `json:"value,omitempty"`
							ValueFrom *struct {
								FieldRef struct {
									APIVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
							} `json:"valueFrom,omitempty"`
						} `json:"env"`
						Image           string `json:"image"`
						ImagePullPolicy string `json:"imagePullPolicy"`
						Name            string `json:"name"`
						Resources       struct {
							Limits struct {
								Cpu    string `json:"cpu,omitempty"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								Cpu    string `json:"cpu,omitempty"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						SecurityContext *struct {
							Privileged bool `json:"privileged"`
						} `json:"securityContext,omitempty"`
						TerminationMessagePath   string `json:"terminationMessagePath"`
						TerminationMessagePolicy string `json:"terminationMessagePolicy"`
						VolumeMounts             []struct {
							MountPath string `json:"mountPath"`
							Name      string `json:"name"`
							ReadOnly  bool   `json:"readOnly,omitempty"`
						} `json:"volumeMounts"`
					} `json:"initContainers"`
					RestartPolicy                 string   `json:"restartPolicy"`
					SchedulerName                 string   `json:"schedulerName"`
					SecurityContext               struct{} `json:"securityContext"`
					TerminationGracePeriodSeconds int      `json:"terminationGracePeriodSeconds"`
					Volumes                       []struct {
						ConfigMap *struct {
							DefaultMode int    `json:"defaultMode"`
							Name        string `json:"name"`
							Optional    bool   `json:"optional"`
						} `json:"configMap,omitempty"`
						DownwardAPI *struct {
							DefaultMode int `json:"defaultMode"`
							Items       []struct {
								FieldRef struct {
									APIVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
								Path string `json:"path"`
							} `json:"items"`
						} `json:"downwardAPI,omitempty"`
						EmptyDir              *struct{} `json:"emptyDir,omitempty"`
						Name                  string    `json:"name"`
						PersistentVolumeClaim *struct {
							ClaimName string `json:"claimName"`
						} `json:"persistentVolumeClaim,omitempty"`
						Secret *struct {
							DefaultMode int `json:"defaultMode"`
							Items       []struct {
								Key  string `json:"key"`
								Path string `json:"path"`
							} `json:"items,omitempty"`
							Optional   bool   `json:"optional"`
							SecretName string `json:"secretName"`
						} `json:"secret,omitempty"`
					} `json:"volumes"`
				} `json:"spec"`
			} `json:"template"`
			UpdateStrategy struct {
				Type string `json:"type"`
			} `json:"updateStrategy"`
			VolumeClaimTemplates []struct {
				APIVersion string `json:"apiVersion"`
				Kind       string `json:"kind"`
				Metadata   struct {
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Name              string      `json:"name"`
				} `json:"metadata"`
				Spec struct {
					AccessModes []string `json:"accessModes"`
					Resources   struct {
						Requests struct {
							Storage string `json:"storage"`
						} `json:"requests"`
					} `json:"resources"`
					VolumeMode string `json:"volumeMode"`
				} `json:"spec"`
				Status struct {
					Phase string `json:"phase"`
				} `json:"status"`
			} `json:"volumeClaimTemplates"`
		} `json:"spec"`
		Status struct {
			CollisionCount     int    `json:"collisionCount"`
			CurrentRevision    string `json:"currentRevision"`
			ObservedGeneration int    `json:"observedGeneration"`
			ReadyReplicas      int    `json:"readyReplicas"`
			Replicas           int    `json:"replicas"`
			UpdateRevision     string `json:"updateRevision"`
			UpdatedReplicas    int    `json:"updatedReplicas"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type Deployment struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				Deployment_Kubernetes_Io_Revision string `json:"deployment.kubernetes.io/revision"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Generation        int       `json:"generation"`
			Labels            struct {
				Common_k8s_elastic_co_templateHash string `json:"common.k8s.elastic.co/template-hash"`
				Common_K8S_Elastic_Co_Type         string `json:"common.k8s.elastic.co/type"`
				Kibana_K8S_Elastic_Co_Name         string `json:"kibana.k8s.elastic.co/name"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata struct {
						F_Annotations *struct {
							_                                   struct{} `json:"."`
							F_Deployment_Kubernetes_Io_Revision struct{} `json:"f:deployment.kubernetes.io/revision"`
						} `json:"f:annotations,omitempty"`
						F_Labels *struct {
							_                                    struct{} `json:"."`
							F_common_k8s_elastic_co_templateHash struct{} `json:"f:common.k8s.elastic.co/template-hash"`
							F_Common_K8S_Elastic_Co_Type         struct{} `json:"f:common.k8s.elastic.co/type"`
							F_Kibana_K8S_Elastic_Co_Name         struct{} `json:"f:kibana.k8s.elastic.co/name"`
						} `json:"f:labels,omitempty"`
						F_OwnerReferences *struct {
							_ struct{} `json:"."`
							// "k:{\"uid\":\"b0d54df0-cd08-41a9-b58c-de75022f4448\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:ownerReferences,omitempty"`
					} `json:"f:metadata"`
					F_Spec *struct {
						F_ProgressDeadlineSeconds struct{} `json:"f:progressDeadlineSeconds"`
						F_Replicas                struct{} `json:"f:replicas"`
						F_RevisionHistoryLimit    struct{} `json:"f:revisionHistoryLimit"`
						F_Selector                struct{} `json:"f:selector"`
						F_Strategy                struct {
							F_RollingUpdate struct {
								_                struct{} `json:"."`
								F_MaxSurge       struct{} `json:"f:maxSurge"`
								F_MaxUnavailable struct{} `json:"f:maxUnavailable"`
							} `json:"f:rollingUpdate"`
							F_Type struct{} `json:"f:type"`
						} `json:"f:strategy"`
						F_Template struct {
							F_Metadata struct {
								F_Annotations struct {
									_                        struct{} `json:"."`
									F_Co_Elastic_Logs_Module struct{} `json:"f:co.elastic.logs/module"`
								} `json:"f:annotations"`
								F_Labels struct {
									_                                      struct{} `json:"."`
									F_Common_K8S_Elastic_Co_Type           struct{} `json:"f:common.k8s.elastic.co/type"`
									F_kibana_k8s_elastic_co_configChecksum struct{} `json:"f:kibana.k8s.elastic.co/config-checksum"`
									F_Kibana_K8S_Elastic_Co_Name           struct{} `json:"f:kibana.k8s.elastic.co/name"`
									F_Kibana_K8S_Elastic_Co_Version        struct{} `json:"f:kibana.k8s.elastic.co/version"`
								} `json:"f:labels"`
							} `json:"f:metadata"`
							F_Spec struct {
								F_AutomountServiceAccountToken struct{} `json:"f:automountServiceAccountToken"`
								F_Containers                   struct {
									// "k:{\"name\":\"kibana\"}" cannot be unmarshalled into a struct field by encoding/json.
								} `json:"f:containers"`
								F_DnsPolicy      struct{} `json:"f:dnsPolicy"`
								F_InitContainers struct {
									_ struct{} `json:"."`
									// "k:{\"name\":\"elastic-internal-init-config\"}" cannot be unmarshalled into a struct field by encoding/json.
								} `json:"f:initContainers"`
								F_RestartPolicy                 struct{} `json:"f:restartPolicy"`
								F_SchedulerName                 struct{} `json:"f:schedulerName"`
								F_SecurityContext               struct{} `json:"f:securityContext"`
								F_TerminationGracePeriodSeconds struct{} `json:"f:terminationGracePeriodSeconds"`
								F_Volumes                       struct {
									_ struct{} `json:"."`
									// "k:{\"name\":\"elastic-internal-kibana-config\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elastic-internal-kibana-config-local\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"elasticsearch-certs\"}" cannot be unmarshalled into a struct field by encoding/json.
									// "k:{\"name\":\"kibana-data\"}" cannot be unmarshalled into a struct field by encoding/json.
								} `json:"f:volumes"`
							} `json:"f:spec"`
						} `json:"f:template"`
					} `json:"f:spec,omitempty"`
					F_Status *struct {
						F_AvailableReplicas struct{} `json:"f:availableReplicas"`
						F_Conditions        struct {
							_ struct{} `json:"."`
							// "k:{\"type\":\"Available\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"type\":\"Progressing\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:conditions"`
						F_ObservedGeneration struct{} `json:"f:observedGeneration"`
						F_ReadyReplicas      struct{} `json:"f:readyReplicas"`
						F_Replicas           struct{} `json:"f:replicas"`
						F_UpdatedReplicas    struct{} `json:"f:updatedReplicas"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			OwnerReferences []struct {
				APIVersion         string `json:"apiVersion"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
				Controller         bool   `json:"controller"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				Uid                string `json:"uid"`
			} `json:"ownerReferences"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			ProgressDeadlineSeconds int `json:"progressDeadlineSeconds"`
			Replicas                int `json:"replicas"`
			RevisionHistoryLimit    int `json:"revisionHistoryLimit"`
			Selector                struct {
				MatchLabels struct {
					Common_K8S_Elastic_Co_Type string `json:"common.k8s.elastic.co/type"`
					Kibana_K8S_Elastic_Co_Name string `json:"kibana.k8s.elastic.co/name"`
				} `json:"matchLabels"`
			} `json:"selector"`
			Strategy struct {
				RollingUpdate struct {
					MaxSurge       string `json:"maxSurge"`
					MaxUnavailable string `json:"maxUnavailable"`
				} `json:"rollingUpdate"`
				Type string `json:"type"`
			} `json:"strategy"`
			Template struct {
				Metadata struct {
					Annotations struct {
						Co_Elastic_Logs_Module string `json:"co.elastic.logs/module"`
					} `json:"annotations"`
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						Common_K8S_Elastic_Co_Type           string `json:"common.k8s.elastic.co/type"`
						Kibana_k8s_elastic_co_configChecksum string `json:"kibana.k8s.elastic.co/config-checksum"`
						Kibana_K8S_Elastic_Co_Name           string `json:"kibana.k8s.elastic.co/name"`
						Kibana_K8S_Elastic_Co_Version        string `json:"kibana.k8s.elastic.co/version"`
					} `json:"labels"`
				} `json:"metadata"`
				Spec struct {
					AutomountServiceAccountToken bool `json:"automountServiceAccountToken"`
					Containers                   []struct {
						Image           string `json:"image"`
						ImagePullPolicy string `json:"imagePullPolicy"`
						Name            string `json:"name"`
						Ports           []struct {
							ContainerPort int    `json:"containerPort"`
							Name          string `json:"name"`
							Protocol      string `json:"protocol"`
						} `json:"ports"`
						ReadinessProbe struct {
							FailureThreshold int `json:"failureThreshold"`
							HTTPGet          struct {
								Path   string `json:"path"`
								Port   int    `json:"port"`
								Scheme string `json:"scheme"`
							} `json:"httpGet"`
							PeriodSeconds    int `json:"periodSeconds"`
							SuccessThreshold int `json:"successThreshold"`
							TimeoutSeconds   int `json:"timeoutSeconds"`
						} `json:"readinessProbe"`
						Resources struct {
							Limits struct {
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						TerminationMessagePath   string `json:"terminationMessagePath"`
						TerminationMessagePolicy string `json:"terminationMessagePolicy"`
						VolumeMounts             []struct {
							MountPath string `json:"mountPath"`
							Name      string `json:"name"`
							ReadOnly  bool   `json:"readOnly,omitempty"`
						} `json:"volumeMounts"`
					} `json:"containers"`
					DnsPolicy      string `json:"dnsPolicy"`
					InitContainers []struct {
						Command []string `json:"command"`
						Env     []struct {
							Name      string `json:"name"`
							ValueFrom struct {
								FieldRef struct {
									APIVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
							} `json:"valueFrom"`
						} `json:"env"`
						Image           string `json:"image"`
						ImagePullPolicy string `json:"imagePullPolicy"`
						Name            string `json:"name"`
						Resources       struct {
							Limits struct {
								Cpu    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								Cpu    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						SecurityContext struct {
							Privileged bool `json:"privileged"`
						} `json:"securityContext"`
						TerminationMessagePath   string `json:"terminationMessagePath"`
						TerminationMessagePolicy string `json:"terminationMessagePolicy"`
						VolumeMounts             []struct {
							MountPath string `json:"mountPath"`
							Name      string `json:"name"`
							ReadOnly  bool   `json:"readOnly,omitempty"`
						} `json:"volumeMounts"`
					} `json:"initContainers"`
					RestartPolicy                 string   `json:"restartPolicy"`
					SchedulerName                 string   `json:"schedulerName"`
					SecurityContext               struct{} `json:"securityContext"`
					TerminationGracePeriodSeconds int      `json:"terminationGracePeriodSeconds"`
					Volumes                       []struct {
						EmptyDir *struct{} `json:"emptyDir,omitempty"`
						Name     string    `json:"name"`
						Secret   *struct {
							DefaultMode int    `json:"defaultMode"`
							Optional    bool   `json:"optional"`
							SecretName  string `json:"secretName"`
						} `json:"secret,omitempty"`
					} `json:"volumes"`
				} `json:"spec"`
			} `json:"template"`
		} `json:"spec"`
		Status struct {
			AvailableReplicas int `json:"availableReplicas"`
			Conditions        []struct {
				LastTransitionTime time.Time `json:"lastTransitionTime"`
				LastUpdateTime     time.Time `json:"lastUpdateTime"`
				Message            string    `json:"message"`
				Reason             string    `json:"reason"`
				Status             string    `json:"status"`
				Type               string    `json:"type"`
			} `json:"conditions"`
			ObservedGeneration int `json:"observedGeneration"`
			ReadyReplicas      int `json:"readyReplicas"`
			Replicas           int `json:"replicas"`
			UpdatedReplicas    int `json:"updatedReplicas"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type PVC struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				Pv_kubernetes_io_bindCompleted               string `json:"pv.kubernetes.io/bind-completed"`
				Pv_kubernetes_io_boundByController           string `json:"pv.kubernetes.io/bound-by-controller"`
				Volume_beta_kubernetes_io_storageProvisioner string `json:"volume.beta.kubernetes.io/storage-provisioner"`
				Volume_kubernetes_io_selectedNode            string `json:"volume.kubernetes.io/selected-node"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Finalizers        []string  `json:"finalizers"`
			Labels            struct {
				Common_K8S_Elastic_Co_Type                   string `json:"common.k8s.elastic.co/type"`
				Elasticsearch_k8s_elastic_co_clusterName     string `json:"elasticsearch.k8s.elastic.co/cluster-name"`
				Elasticsearch_k8s_elastic_co_statefulsetName string `json:"elasticsearch.k8s.elastic.co/statefulset-name"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata struct {
						F_Annotations *struct {
							_                                              *struct{} `json:".,omitempty"`
							F_pv_kubernetes_io_bindCompleted               *struct{} `json:"f:pv.kubernetes.io/bind-completed,omitempty"`
							F_pv_kubernetes_io_boundByController           *struct{} `json:"f:pv.kubernetes.io/bound-by-controller,omitempty"`
							F_volume_beta_kubernetes_io_storageProvisioner *struct{} `json:"f:volume.beta.kubernetes.io/storage-provisioner,omitempty"`
							F_volume_kubernetes_io_selectedNode            *struct{} `json:"f:volume.kubernetes.io/selected-node,omitempty"`
						} `json:"f:annotations,omitempty"`
						F_Labels *struct {
							_                                              struct{} `json:"."`
							F_Common_K8S_Elastic_Co_Type                   struct{} `json:"f:common.k8s.elastic.co/type"`
							F_elasticsearch_k8s_elastic_co_clusterName     struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-name"`
							F_elasticsearch_k8s_elastic_co_statefulsetName struct{} `json:"f:elasticsearch.k8s.elastic.co/statefulset-name"`
						} `json:"f:labels,omitempty"`
						F_OwnerReferences *struct {
							_ struct{} `json:"."`
							// "k:{\"uid\":\"5de07191-561b-4b60-b3de-ab58726c31d4\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:ownerReferences,omitempty"`
					} `json:"f:metadata"`
					F_Spec *struct {
						F_AccessModes struct{} `json:"f:accessModes"`
						F_Resources   struct {
							F_Requests struct {
								_         struct{} `json:"."`
								F_Storage struct{} `json:"f:storage"`
							} `json:"f:requests"`
						} `json:"f:resources"`
						F_VolumeMode struct{} `json:"f:volumeMode"`
						F_VolumeName struct{} `json:"f:volumeName"`
					} `json:"f:spec,omitempty"`
					F_Status *struct {
						F_AccessModes struct{} `json:"f:accessModes"`
						F_Capacity    struct {
							_         struct{} `json:"."`
							F_Storage struct{} `json:"f:storage"`
						} `json:"f:capacity"`
						F_Phase struct{} `json:"f:phase"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			OwnerReferences []struct {
				APIVersion string `json:"apiVersion"`
				Kind       string `json:"kind"`
				Name       string `json:"name"`
				Uid        string `json:"uid"`
			} `json:"ownerReferences"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			AccessModes []string `json:"accessModes"`
			Resources   struct {
				Requests struct {
					Storage string `json:"storage"`
				} `json:"requests"`
			} `json:"resources"`
			StorageClassName string `json:"storageClassName"`
			VolumeMode       string `json:"volumeMode"`
			VolumeName       string `json:"volumeName"`
		} `json:"spec"`
		Status struct {
			AccessModes []string `json:"accessModes"`
			Capacity    struct {
				Storage string `json:"storage"`
			} `json:"capacity"`
			Phase string `json:"phase"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type Service struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations *struct {
				Service_beta_kubernetes_io_awsLoadBalancerAdditionalResourceTags string `json:"service.beta.kubernetes.io/aws-load-balancer-additional-resource-tags,omitempty"`
				Service_beta_kubernetes_io_awsLoadBalancerInternal               string `json:"service.beta.kubernetes.io/aws-load-balancer-internal"`
				Service_beta_kubernetes_io_awsLoadBalancerProxyProtocol          string `json:"service.beta.kubernetes.io/aws-load-balancer-proxy-protocol,omitempty"`
				Service_beta_kubernetes_io_awsLoadBalancerType                   string `json:"service.beta.kubernetes.io/aws-load-balancer-type"`
			} `json:"annotations,omitempty"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Finalizers        []string  `json:"finalizers,omitempty"`
			Labels            struct {
				Common_K8S_Elastic_Co_Type                   string `json:"common.k8s.elastic.co/type,omitempty"`
				Component                                    string `json:"component,omitempty"`
				Elasticsearch_k8s_elastic_co_clusterName     string `json:"elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
				Elasticsearch_k8s_elastic_co_statefulsetName string `json:"elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
				Kibana_K8S_Elastic_Co_Name                   string `json:"kibana.k8s.elastic.co/name,omitempty"`
				Provider                                     string `json:"provider,omitempty"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata struct {
						F_Annotations *struct {
							_                                                                  struct{}  `json:"."`
							F_service_beta_kubernetes_io_awsLoadBalancerAdditionalResourceTags *struct{} `json:"f:service.beta.kubernetes.io/aws-load-balancer-additional-resource-tags,omitempty"`
							F_service_beta_kubernetes_io_awsLoadBalancerInternal               struct{}  `json:"f:service.beta.kubernetes.io/aws-load-balancer-internal"`
							F_service_beta_kubernetes_io_awsLoadBalancerProxyProtocol          *struct{} `json:"f:service.beta.kubernetes.io/aws-load-balancer-proxy-protocol,omitempty"`
							F_service_beta_kubernetes_io_awsLoadBalancerType                   struct{}  `json:"f:service.beta.kubernetes.io/aws-load-balancer-type"`
						} `json:"f:annotations,omitempty"`
						F_Finalizers *struct {
							_ struct{} `json:"."`
							// "v:\"service.kubernetes.io/load-balancer-cleanup\"" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:finalizers,omitempty"`
						F_Labels *struct {
							_                                              struct{}  `json:"."`
							F_Common_K8S_Elastic_Co_Type                   *struct{} `json:"f:common.k8s.elastic.co/type,omitempty"`
							F_Component                                    *struct{} `json:"f:component,omitempty"`
							F_elasticsearch_k8s_elastic_co_clusterName     *struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
							F_elasticsearch_k8s_elastic_co_statefulsetName *struct{} `json:"f:elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
							F_Kibana_K8S_Elastic_Co_Name                   *struct{} `json:"f:kibana.k8s.elastic.co/name,omitempty"`
							F_Provider                                     *struct{} `json:"f:provider,omitempty"`
						} `json:"f:labels,omitempty"`
						F_OwnerReferences *struct {
							_ struct{} `json:"."`
							// "k:{\"uid\":\"5de07191-561b-4b60-b3de-ab58726c31d4\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"uid\":\"b0d54df0-cd08-41a9-b58c-de75022f4448\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:ownerReferences,omitempty"`
					} `json:"f:metadata"`
					F_Spec *struct {
						F_ClusterIp             *struct{} `json:"f:clusterIP,omitempty"`
						F_ExternalTrafficPolicy *struct{} `json:"f:externalTrafficPolicy,omitempty"`
						F_IpFamilyPolicy        *struct{} `json:"f:ipFamilyPolicy,omitempty"`
						F_Ports                 struct {
							_ struct{} `json:"."`
							// "k:{\"port\":443,\"protocol\":\"TCP\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"port\":5601,\"protocol\":\"TCP\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"port\":9200,\"protocol\":\"TCP\"}" cannot be unmarshalled into a struct field by encoding/json.
							// "k:{\"port\":9300,\"protocol\":\"TCP\"}" cannot be unmarshalled into a struct field by encoding/json.
						} `json:"f:ports"`
						F_PublishNotReadyAddresses *struct{} `json:"f:publishNotReadyAddresses,omitempty"`
						F_Selector                 *struct {
							_                                              struct{}  `json:"."`
							F_Common_K8S_Elastic_Co_Type                   struct{}  `json:"f:common.k8s.elastic.co/type"`
							F_elasticsearch_k8s_elastic_co_clusterName     *struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
							F_elasticsearch_k8s_elastic_co_statefulsetName *struct{} `json:"f:elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
							F_Kibana_K8S_Elastic_Co_Name                   *struct{} `json:"f:kibana.k8s.elastic.co/name,omitempty"`
						} `json:"f:selector,omitempty"`
						F_SessionAffinity struct{} `json:"f:sessionAffinity"`
						F_Type            struct{} `json:"f:type"`
					} `json:"f:spec,omitempty"`
					F_Status *struct {
						F_LoadBalancer struct {
							F_Ingress struct{} `json:"f:ingress"`
						} `json:"f:loadBalancer"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			OwnerReferences []struct {
				APIVersion         string `json:"apiVersion"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
				Controller         bool   `json:"controller"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				Uid                string `json:"uid"`
			} `json:"ownerReferences,omitempty"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			ClusterIp             string   `json:"clusterIP"`
			ClusterIPs            []string `json:"clusterIPs"`
			ExternalTrafficPolicy string   `json:"externalTrafficPolicy,omitempty"`
			IpFamilies            []string `json:"ipFamilies"`
			IpFamilyPolicy        string   `json:"ipFamilyPolicy"`
			/*Ports                 []struct {
				Name       string `json:"name"`
				NodePort   int    `json:"nodePort,omitempty"`
				Port       int    `json:"port"`
				Protocol   string `json:"protocol"`
				TargetPort int    `json:"targetPort"`
			} `json:"ports"`
			*/PublishNotReadyAddresses bool `json:"publishNotReadyAddresses,omitempty"`
			Selector                   *struct {
				Common_K8S_Elastic_Co_Type                   string `json:"common.k8s.elastic.co/type"`
				Elasticsearch_k8s_elastic_co_clusterName     string `json:"elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
				Elasticsearch_k8s_elastic_co_statefulsetName string `json:"elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
				Kibana_K8S_Elastic_Co_Name                   string `json:"kibana.k8s.elastic.co/name,omitempty"`
			} `json:"selector,omitempty"`
			SessionAffinity string `json:"sessionAffinity"`
			Type            string `json:"type"`
		} `json:"spec"`
		Status struct {
			LoadBalancer struct {
				Ingress []struct {
					Hostname string `json:"hostname"`
				} `json:"ingress,omitempty"`
			} `json:"loadBalancer"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type EndPoint struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations *struct {
				Endpoints_kubernetes_io_lastChangeTriggerTime time.Time `json:"endpoints.kubernetes.io/last-change-trigger-time"`
			} `json:"annotations,omitempty"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				Common_K8S_Elastic_Co_Type                   string `json:"common.k8s.elastic.co/type,omitempty"`
				Elasticsearch_k8s_elastic_co_clusterName     string `json:"elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
				Elasticsearch_k8s_elastic_co_statefulsetName string `json:"elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
				Endpointslice_kubernetes_io_skipMirror       string `json:"endpointslice.kubernetes.io/skip-mirror,omitempty"`
				Kibana_K8S_Elastic_Co_Name                   string `json:"kibana.k8s.elastic.co/name,omitempty"`
				Service_Kubernetes_Io_Headless               string `json:"service.kubernetes.io/headless"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_Metadata struct {
						F_Annotations *struct {
							_                                               struct{} `json:"."`
							F_endpoints_kubernetes_io_lastChangeTriggerTime struct{} `json:"f:endpoints.kubernetes.io/last-change-trigger-time"`
						} `json:"f:annotations,omitempty"`
						F_Labels struct {
							_                                              struct{}  `json:"."`
							F_Common_K8S_Elastic_Co_Type                   *struct{} `json:"f:common.k8s.elastic.co/type,omitempty"`
							F_elasticsearch_k8s_elastic_co_clusterName     *struct{} `json:"f:elasticsearch.k8s.elastic.co/cluster-name,omitempty"`
							F_elasticsearch_k8s_elastic_co_statefulsetName *struct{} `json:"f:elasticsearch.k8s.elastic.co/statefulset-name,omitempty"`
							F_endpointslice_kubernetes_io_skipMirror       *struct{} `json:"f:endpointslice.kubernetes.io/skip-mirror,omitempty"`
							F_Kibana_K8S_Elastic_Co_Name                   *struct{} `json:"f:kibana.k8s.elastic.co/name,omitempty"`
							F_Service_Kubernetes_Io_Headless               *struct{} `json:"f:service.kubernetes.io/headless,omitempty"`
						} `json:"f:labels"`
					} `json:"f:metadata"`
					F_Subsets struct{} `json:"f:subsets"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Subsets []struct {
			Addresses []struct {
				Hostname  string `json:"hostname,omitempty"`
				Ip        string `json:"ip"`
				NodeName  string `json:"nodeName,omitempty"`
				TargetRef *struct {
					Kind            string `json:"kind"`
					Name            string `json:"name"`
					Namespace       string `json:"namespace"`
					ResourceVersion string `json:"resourceVersion"`
					Uid             string `json:"uid"`
				} `json:"targetRef,omitempty"`
			} `json:"addresses"`
			Ports []struct {
				Name     string `json:"name"`
				Port     int    `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"ports"`
		} `json:"subsets"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

type Manifest struct {
	CollectionDate      time.Time `json:"collectionDate"`
	DiagType            string    `json:"diagType"`
	DiagVersion         string    `json:"diagVersion"`
	IncludedDiagnostics []struct {
		DiagPath string `json:"diagPath"`
		DiagType string `json:"diagType"`
	} `json:"includedDiagnostics"`
}

type StorageClass struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		AllowVolumeExpansion bool `json:"allowVolumeExpansion"`
		AllowedTopologies    []struct {
			MatchLabelExpressions []struct {
				Key    string   `json:"key"`
				Values []string `json:"values"`
			} `json:"matchLabelExpressions"`
		} `json:"allowedTopologies,omitempty"`
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations *struct {
				Kubectl_kubernetes_io_lastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration,omitempty"`
				Storageclass_kubernetes_io_isDefaultClass      string `json:"storageclass.kubernetes.io/is-default-class,omitempty"`
			} `json:"annotations,omitempty"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				Addonmanager_Kubernetes_Io_Mode string `json:"addonmanager.kubernetes.io/mode,omitempty"`
				Kubernetes_io_clusterService    string `json:"kubernetes.io/cluster-service"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					F_AllowVolumeExpansion struct{}  `json:"f:allowVolumeExpansion"`
					F_AllowedTopologies    *struct{} `json:"f:allowedTopologies,omitempty"`
					F_Metadata             struct {
						F_Annotations *struct {
							_                                                struct{}  `json:"."`
							F_kubectl_kubernetes_io_lastAppliedConfiguration *struct{} `json:"f:kubectl.kubernetes.io/last-applied-configuration,omitempty"`
							F_storageclass_kubernetes_io_isDefaultClass      *struct{} `json:"f:storageclass.kubernetes.io/is-default-class,omitempty"`
						} `json:"f:annotations,omitempty"`
						F_Labels struct {
							_                                 struct{}  `json:"."`
							F_Addonmanager_Kubernetes_Io_Mode *struct{} `json:"f:addonmanager.kubernetes.io/mode,omitempty"`
							F_kubernetes_io_clusterService    struct{}  `json:"f:kubernetes.io/cluster-service"`
						} `json:"f:labels"`
					} `json:"f:metadata"`
					F_MountOptions *struct{} `json:"f:mountOptions,omitempty"`
					F_Parameters   struct {
						_                    struct{}  `json:"."`
						F_Cachingmode        *struct{} `json:"f:cachingmode,omitempty"`
						F_Kind               *struct{} `json:"f:kind,omitempty"`
						F_SkuName            *struct{} `json:"f:skuName,omitempty"`
						F_Skuname            *struct{} `json:"f:skuname,omitempty"`
						F_Storageaccounttype *struct{} `json:"f:storageaccounttype,omitempty"`
					} `json:"f:parameters"`
					F_Provisioner       struct{} `json:"f:provisioner"`
					F_ReclaimPolicy     struct{} `json:"f:reclaimPolicy"`
					F_VolumeBindingMode struct{} `json:"f:volumeBindingMode"`
				} `json:"fieldsV1"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		MountOptions []string `json:"mountOptions,omitempty"`
		Parameters   struct {
			Cachingmode        string `json:"cachingmode,omitempty"`
			Kind               string `json:"kind,omitempty"`
			SkuName            string `json:"skuName,omitempty"`
			Skuname            string `json:"skuname,omitempty"`
			Storageaccounttype string `json:"storageaccounttype,omitempty"`
		} `json:"parameters"`
		Provisioner       string `json:"provisioner"`
		ReclaimPolicy     string `json:"reclaimPolicy"`
		VolumeBindingMode string `json:"volumeBindingMode"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}

func LoadNodes(filename string) (Nodes, error) {
	var config Nodes
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadPods(filename string) (Pods, error) {
	var config Pods
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadES(filename string) (ES, error) {
	var config ES
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadKibana(filename string) (Kibana, error) {
	var config Kibana
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadEvents(filename string) (Events, error) {
	var config Events
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadDiagV(filename string) (DiagVersion, error) {
	var config DiagVersion
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadStatefulSet(filename string) (StatefulSet, error) {
	var config StatefulSet
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadDeploy(filename string) (Deployment, error) {
	var config Deployment
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadPVC(filename string) (PVC, error) {
	var config PVC
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadService(filename string) (Service, error) {
	var config Service
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadEndPoint(filename string) (EndPoint, error) {
	var config EndPoint
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadManifest(filename string) (Manifest, error) {
	var config Manifest
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}

func LoadStorageClass(filename string) (StorageClass, error) {
	var config StorageClass
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)
	return config, err
}
