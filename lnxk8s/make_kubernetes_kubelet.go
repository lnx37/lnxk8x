package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var V_make_kubernetes_kubelet string = `
#!/bin/bash
set -e
set -o pipefail
set -u
cd "$(dirname "$0")"
date
source ./env.sh
echo "${KUBE_APISERVER_IP}"
echo "${CONTAINER_RUNTIME}"
mkdir -p artifact/kubernetes/kubelet
cd artifact/kubernetes/kubelet
KUBECONFIG="./kubelet-bootstrap.kubeconfig"
KUBE_APISERVER="https://${KUBE_APISERVER_IP}:6443"
[ -f "$KUBECONFIG" ] && rm -f "$KUBECONFIG"
kubectl config set-cluster kubernetes \
  --certificate-authority=../ca.pem \
  --embed-certs=true \
  --kubeconfig="$KUBECONFIG" \
  --server="${KUBE_APISERVER}"
kubectl config set-credentials system:node:hostname \
  --kubeconfig="$KUBECONFIG" \
  --token=abcdef.0123456789abcdef
kubectl config set-context system:node:hostname@kubernetes \
  --cluster=kubernetes \
  --kubeconfig="$KUBECONFIG" \
  --user=system:node:hostname
kubectl config use-context system:node:hostname@kubernetes --kubeconfig="$KUBECONFIG"
cat <<EOF >kubelet-config.yml
apiVersion: kubelet.config.k8s.io/v1beta1
authentication:
  anonymous:
    enabled: false
  webhook:
    cacheTTL: 0s
    enabled: true
  x509:
    # clientCAFile: /etc/kubernetes/pki/ca.crt
    clientCAFile: /opt/kubernetes/ca.pem
authorization:
  mode: Webhook
  webhook:
    cacheAuthorizedTTL: 0s
    cacheUnauthorizedTTL: 0s
cgroupDriver: systemd
clusterDNS:
- 10.96.0.10
clusterDomain: cluster.local
containerRuntimeEndpoint: ""
cpuManagerReconcilePeriod: 0s
evictionPressureTransitionPeriod: 0s
fileCheckFrequency: 0s
healthzBindAddress: 127.0.0.1
healthzPort: 10248
httpCheckFrequency: 0s
imageMinimumGCAge: 0s
kind: KubeletConfiguration
logging:
  flushFrequency: 0
  options:
    json:
      infoBufferSize: "0"
  verbosity: 0
memorySwap: {}
nodeStatusReportFrequency: 0s
nodeStatusUpdateFrequency: 0s
rotateCertificates: true
runtimeRequestTimeout: 0s
shutdownGracePeriod: 0s
shutdownGracePeriodCriticalPods: 0s
staticPodPath: /opt/kubernetes/manifests
streamingConnectionIdleTimeout: 0s
syncFrequency: 0s
volumeStatsAggPeriod: 0s
EOF
if [ "${CONTAINER_RUNTIME}" = "containerd" ]; then
  CONTAINER_RUNTIME_ENDPOINT="unix:///run/containerd/containerd.sock"
fi
if [ "${CONTAINER_RUNTIME}" = "crio" ]; then
  CONTAINER_RUNTIME_ENDPOINT="unix:///run/crio/crio.sock"
fi
if [ "${CONTAINER_RUNTIME}" = "docker" ]; then
  CONTAINER_RUNTIME_ENDPOINT="unix:///run/cri-dockerd.sock"
fi
cat <<EOF >kubelet.conf
KUBELET_OPTS=" \\
  --bootstrap-kubeconfig=/opt/kubernetes/kubelet/kubelet-bootstrap.kubeconfig \\
  --config=/opt/kubernetes/kubelet/kubelet-config.yml \\
  --container-runtime-endpoint=${CONTAINER_RUNTIME_ENDPOINT} \\
  --kubeconfig=/opt/kubernetes/kubelet/kubelet.kubeconfig \\
  --pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.9 \\
"
EOF
cat <<\EOF >kubelet.service
[Unit]
Description=kubelet: The Kubernetes Node Agent
Documentation=https://kubernetes.io/docs/
Wants=network-online.target
After=network-online.target containerd.service cri-docker.service
[Service]
EnvironmentFile=/opt/kubernetes/kubelet/kubelet.conf
ExecStart=/opt/kubernetes/kubelet/kubelet $KUBELET_OPTS
Restart=always
StartLimitInterval=0
RestartSec=10
[Install]
WantedBy=multi-user.target
EOF
date
`

func X_make_kubernetes_kubelet(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace(V_make_kubernetes_kubelet)
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
