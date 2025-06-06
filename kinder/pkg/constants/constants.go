/*
Copyright 2019 The Kubernetes Authors.

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

package constants

// constants inherited from kind.
// those values are replicated here with the goal to keep under strict control kind dependencies
const (
	// DefaultNodeImage is the default name:tag for a base image
	DefaultBaseImage = "kindest/base:latest"

	// DefaultNodeImage is the default name:tag for a node image
	DefaultNodeImage = "kindest/node:latest"

	// ControlPlaneNodeRoleValue identifies a node that hosts a Kubernetes
	// control-plane.
	//
	// NOTE: in single node clusters, control-plane nodes act as worker nodes
	ControlPlaneNodeRoleValue string = "control-plane"

	// WorkerNodeRoleValue identifies a node that hosts a Kubernetes worker
	WorkerNodeRoleValue string = "worker"

	// ExternalLoadBalancerNodeRoleValue identifies a node that hosts an
	// external load balancer for the API server in HA configurations.
	//
	// Please note that `kind` nodes (containers) hosting external load balancer are not kubernetes nodes
	ExternalLoadBalancerNodeRoleValue string = "external-load-balancer"

	// ExternalEtcdNodeRoleValue identifies a node that hosts an external-etcd
	// instance.
	//
	// WARNING: this node type is not yet implemented in kind! (in kinder it is implemented)
	//
	// Please note that `kind` nodes (containers) hosting external etcd are not kubernetes nodes
	ExternalEtcdNodeRoleValue string = "external-etcd"

	// DefaultClusterName is the default cluster name
	// TODO: consider if to switch to kinder
	DefaultClusterName = "kind"

	// ClusterLabelKey is applied to each "node" docker container for identification
	// TODO: consider if to switch to a kinder specific label
	ClusterLabelKey = "io.x-k8s.kind.cluster"

	// DeprecatedClusterLabelKey is applied to each "node" docker container for identification
	// This is the deprecated value of ClusterLabelKey, and will be removed in a future release
	DeprecatedClusterLabelKey = "io.k8s.sigs.kind.cluster"

	// NodeRoleLabelKey is applied to each "node" docker container for categorization
	// of nodes by role
	// TODO: consider if to switch to a kinder specific label
	NodeRoleLabelKey = "io.x-k8s.kind.role"

	// DeprecatedNodeRoleLabelKey is applied to each "node" docker container for categorization
	// of nodes by role.
	// This is the deprecated value of NodeRoleKey, and will be removed in a future release
	DeprecatedNodeRoleLabelKey = "io.k8s.sigs.kind.role"

	// PodSubnet defines the default pod subnet used by kind
	// TODO: send a PR to define this value in a kind constant (currently it is not)
	PodSubnet = "10.244.0.0/16"

	// KubeadmConfigPath defines the path to the kubeadm config file in the K8s nodes
	// TODO: send a PR to define this value in a kind constant (currently it is not)
	KubeadmConfigPath = "/kind/kubeadm.conf"

	// KubeadmIgnorePreflightErrors holds the default list of preflight errors to skip
	// on "kubeadm init", "kubeadm join" and "kubeadm upgrade"
	KubeadmIgnorePreflightErrors = "Swap,SystemVerification,FileContent--proc-sys-net-bridge-bridge-nf-call-iptables"

	// APIServerPort is the expected default APIServerPort on the control plane node(s)
	// https://kubernetes.io/docs/reference/access-authn-authz/controlling-access/#api-server-ports-and-ips
	APIServerPort = 6443

	// Token defines a dummy, well known token for automating TLS bootstrap process
	Token = "abcdef.0123456789abcdef"

	// ControlPlanePort defines the port where the control plane is listening on the load balancer node
	ControlPlanePort = 6443

	// LoadBalancerImage defines the loadbalancer image:tag
	LoadBalancerImage = "kindest/haproxy:2.0.0-alpine"

	// ConfigPath defines the path to the config file in the load balancer node
	LoadBalancerConfigPath = "/usr/local/etc/haproxy/haproxy.cfg"
)

// constants used by the ClusterManager / inside actions
const (
	// CertificateKey defines a dummy, well known CertificateKey for automating automatic copy certs process
	// const CertificateKey = "d02db674b27811f4508bf8a5fa19fbe060921340552f13c15c9feb05aaa96824"
	CertificateKey = "0123456789012345678901234567890123456789012345678901234567890123"

	// DiscoveryFile defines the path to a discovery file stored on nodes
	DiscoveryFile = "/kinder/discovery.conf"

	// PatchesDir defines the path to patches stored on node
	PatchesDir = "/kinder/patches"
)

// other constants
const (
	// KinderVersion is the kinder CLI version
	KinderVersion = "1.0.0"
)
